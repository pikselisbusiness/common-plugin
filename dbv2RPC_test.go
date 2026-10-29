package shared

import (
	"strings"
	"testing"
	"time"
)

// recordingHostDB stands in for the host's database: it records statements
// and answers Raw with fixed rows, typed as the MySQL driver returns them.
type recordingHostDB struct {
	statements []string
	args       [][]interface{}
	rows       []map[string]interface{}
}

func (d *recordingHostDB) MigrateModel(string, []MigrateField) error { return nil }
func (d *recordingHostDB) Exec(sql string, values ...interface{}) error {
	d.statements, d.args = append(d.statements, sql), append(d.args, values)
	return nil
}
func (d *recordingHostDB) ExecWithResult(sql string, values ...interface{}) (int64, int64, error) {
	d.statements, d.args = append(d.statements, sql), append(d.args, values)
	return 1, 41, nil
}
func (d *recordingHostDB) Raw(sql string, values ...interface{}) ([]map[string]interface{}, error) {
	d.statements, d.args = append(d.statements, sql), append(d.args, values)
	return d.rows, nil
}

func newDBv2OverRPC(t *testing.T, host DB) DBv2 {
	t.Helper()
	db, closeConn, err := DBv2OverRPC(host)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = closeConn() })
	return db
}

type rpcModel struct {
	ID        int        `gorm:"column:id;primaryKey"`
	OrgID     uint       `gorm:"column:organization_id"`
	OrderID   *int       `gorm:"column:order_id"`
	IsCod     bool       `gorm:"column:is_cod"`
	CodAmount float64    `gorm:"column:cod_amount"`
	CreatedAt *time.Time `gorm:"column:inp_time"`
	Label     string     `gorm:"-"`
}

func TestDBv2OverRPCWritesAndReadsModelValues(t *testing.T) {
	now := time.Date(2026, 10, 29, 12, 0, 0, 0, time.UTC)
	host := &recordingHostDB{rows: []map[string]interface{}{{
		"id": int64(41), "organization_id": int64(3), "order_id": int64(5), "is_cod": int64(1),
		"cod_amount": []byte("10.50"), "inp_time": now,
	}}}
	db := newDBv2OverRPC(t, host)

	// A nil pointer and a time must cross gob as NULL and a time.
	item := rpcModel{OrgID: 3, IsCod: true, CodAmount: 10.5, CreatedAt: &now}
	if res := db.Table("items").Create(&item); res.Error != nil {
		t.Fatalf("create: %v", res.Error)
	}
	if item.ID != 41 {
		t.Fatalf("create must set the id, got %d", item.ID)
	}
	insert := host.statements[len(host.statements)-1]
	if !strings.HasPrefix(insert, "INSERT INTO") || strings.Contains(insert, "Label") {
		t.Fatalf("insert: %s", insert)
	}

	var found []rpcModel
	if res := db.Table("items").Where("organization_id = ?", uint(3)).Where("id IN ?", []int{41, 42}).Find(&found); res.Error != nil {
		t.Fatalf("find: %v", res.Error)
	}
	if len(found) != 1 || found[0].OrderID == nil || *found[0].OrderID != 5 || !found[0].IsCod || found[0].CodAmount != 10.5 ||
		found[0].CreatedAt == nil || !found[0].CreatedAt.Equal(now) {
		t.Fatalf("find decoded: %+v", found)
	}
	sel := host.statements[len(host.statements)-1]
	if !strings.Contains(sel, "organization_id = ?") {
		t.Fatalf("select: %s %v", sel, host.args[len(host.args)-1])
	}
	t.Logf("select: %s %v", sel, host.args[len(host.args)-1])
}
