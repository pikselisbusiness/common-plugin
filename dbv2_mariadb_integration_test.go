//go:build integration

package shared

// Plugin DBv2 against a real MariaDB through the same net/rpc + gob transport
// go-plugin uses. Run with a throwaway database:
//
//	docker run -d --rm --name dbv2-test -e MARIADB_ROOT_PASSWORD=test \
//	  -e MARIADB_DATABASE=dbv2 -p 127.0.0.1:18306:3306 mariadb:11.8
//	COMMON_PLUGIN_TEST_DSN='root:test@tcp(127.0.0.1:18306)/dbv2' go test -tags integration -run MariaDB .
//
// The DSN uses parseTime=true like the host: without it gorm's map scan on
// the host already fails on TIMESTAMP columns, before any plugin code.

import (
	"os"
	"sync"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type shipment struct {
	ID             int        `gorm:"column:id;primaryKey;autoIncrement"`
	OrganizationID uint       `gorm:"column:organization_id;not null;index:idx_shipment_org_order,priority:1"`
	OrderID        *int       `gorm:"column:order_id;index:idx_shipment_org_order,priority:2"`
	IsCod          bool       `gorm:"column:is_cod;type:tinyint(1)"`
	CodAmount      float64    `gorm:"column:cod_amount;type:decimal(10,2)"`
	Weight         float64    `gorm:"column:weight;type:double(10,2)"`
	Tracking       string     `gorm:"column:tracking_number;type:varchar(50)"`
	CreatedAt      *time.Time `gorm:"column:inp_time;type:timestamp NULL"`
	Label          string     `gorm:"-"`
}

func (shipment) TableName() string { return "dbv2_shipments" }

func TestDBv2AgainstMariaDB(t *testing.T) {
	dsn := os.Getenv("COMMON_PLUGIN_TEST_DSN")
	if dsn == "" {
		t.Skip("COMMON_PLUGIN_TEST_DSN not set")
	}
	gdb, err := gorm.Open(mysql.Open(dsn+"?parseTime=true&loc=Local"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		t.Fatal(err)
	}
	host := GormDB{DB: gdb}
	db := newDBv2OverRPC(t, host)

	_ = host.Exec("DROP TABLE IF EXISTS dbv2_shipments")
	fields, err := GetMigrateQueries(shipment{})
	if err != nil {
		t.Fatal(err)
	}
	if err := db.MigrateModel("dbv2_shipments", fields); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	index, _ := host.Raw("SELECT COUNT(*) AS count FROM information_schema.STATISTICS WHERE TABLE_SCHEMA = DATABASE() AND TABLE_NAME = 'dbv2_shipments' AND INDEX_NAME = 'idx_shipment_org_order'")
	nullable, _ := host.Raw("SELECT IS_NULLABLE AS nullable FROM information_schema.COLUMNS WHERE TABLE_SCHEMA = DATABASE() AND TABLE_NAME = 'dbv2_shipments' AND COLUMN_NAME = 'organization_id'")
	t.Logf("migrate: index rows %v, organization_id nullable %v", index, nullable)

	created := time.Date(2026, 10, 29, 12, 30, 15, 0, time.Local)
	order := 5
	withCod := shipment{OrganizationID: 1, OrderID: &order, IsCod: true, CodAmount: 10.5, Weight: 2.25, Tracking: "CC1", CreatedAt: &created}
	if res := db.Table("dbv2_shipments").Create(&withCod); res.Error != nil || withCod.ID == 0 {
		t.Fatalf("create: %v id %d", res.Error, withCod.ID)
	}
	noOrder := shipment{OrganizationID: 2}
	if res := db.Table("dbv2_shipments").Create(&noOrder); res.Error != nil {
		t.Fatalf("create with nil pointers: %v", res.Error)
	}

	var got []shipment
	if res := db.Table("dbv2_shipments").Where("organization_id = ?", uint(1)).Find(&got); res.Error != nil {
		t.Fatalf("find: %v", res.Error)
	}
	if len(got) != 1 || got[0].ID != withCod.ID || got[0].OrderID == nil || *got[0].OrderID != 5 || !got[0].IsCod ||
		got[0].CodAmount != 10.5 || got[0].Weight != 2.25 || got[0].Tracking != "CC1" ||
		got[0].CreatedAt == nil || !got[0].CreatedAt.Equal(created) {
		t.Fatalf("round trip: %+v (created %v)", got, got[0].CreatedAt)
	}

	var other shipment
	if res := db.Table("dbv2_shipments").Where("organization_id = ? AND id = ?", uint(2), noOrder.ID).First(&other); res.Error != nil ||
		other.ID != noOrder.ID || other.OrderID != nil || other.CreatedAt != nil {
		t.Fatalf("nil pointers must stay nil: %+v %v", other, res.Error)
	}

	if res := db.Table("dbv2_shipments").Where("organization_id = ? AND id = ?", uint(1), withCod.ID).
		Updates(map[string]interface{}{"tracking_number": "CC2", "order_id": (*int)(nil)}); res.Error != nil || res.RowsAffected != 1 {
		t.Fatalf("update: %v rows %d", res.Error, res.RowsAffected)
	}
	if res := db.Table("dbv2_shipments").Where("organization_id = ? AND id = ?", uint(2), withCod.ID).Updates(map[string]interface{}{"tracking_number": "X"}); res.RowsAffected != 0 {
		t.Fatal("another organization must not update the row")
	}
	var count int64
	db.Table("dbv2_shipments").Where("organization_id = ? AND id IN ?", uint(1), []int{withCod.ID, noOrder.ID}).Count(&count)
	if count != 1 {
		t.Fatalf("IN within the organization: %d", count)
	}

	// Concurrent inserts on a pool must each get their own id.
	var wg sync.WaitGroup
	ids := make([]int, 20)
	for i := range ids {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			row := shipment{OrganizationID: 3, Tracking: string(rune('a' + i))}
			if res := db.Table("dbv2_shipments").Create(&row); res.Error == nil {
				ids[i] = row.ID
			}
		}(i)
	}
	wg.Wait()
	for i, id := range ids {
		var row shipment
		db.Table("dbv2_shipments").Where("id = ?", id).First(&row)
		if row.Tracking != string(rune('a'+i)) {
			t.Fatalf("insert %d got id %d of row %q", i, id, row.Tracking)
		}
	}

	if res := db.Table("dbv2_shipments").Where("organization_id = ?", uint(3)).Delete(nil); res.Error != nil || res.RowsAffected != 20 {
		t.Fatalf("delete: %v rows %d", res.Error, res.RowsAffected)
	}
}
