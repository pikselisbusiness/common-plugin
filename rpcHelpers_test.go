package shared

import (
	"database/sql"
	"encoding/json"
	"testing"
	"time"
)

type scanModel struct {
	ID        int          `gorm:"column:id"`
	OrderID   *int         `gorm:"column:order_id"`
	IsCod     bool         `gorm:"column:is_cod"`
	CodAmount float64      `gorm:"column:cod_amount"`
	Weight    *float64     `gorm:"column:weight"`
	Count     uint         `gorm:"column:count"`
	Name      string       `gorm:"column:name"`
	Code      string       `gorm:"column:code"`
	CreatedAt *time.Time   `gorm:"column:inp_time"`
	UpdatedAt time.Time    `gorm:"column:updated_at"`
	DeletedAt sql.NullTime `gorm:"column:deleted_at"`
	Note      *string      `gorm:"column:note"`
	Skipped   string       `gorm:"-"`
}

// Rows as gorm's map scan hands them over from the MySQL driver, with and
// without parseTime, and as JSON decoding produces them.
func TestScanRowsIntoConvertsDriverValues(t *testing.T) {
	now := time.Date(2026, 10, 29, 12, 30, 0, 0, time.Local)
	rows := []map[string]interface{}{
		{
			"id": int64(7), "order_id": int64(5), "is_cod": int64(1), "cod_amount": []byte("10.50"),
			"weight": float64(2.5), "count": int64(3), "name": []byte("LT"), "code": int64(42),
			"inp_time": now, "updated_at": []byte("2026-10-29 12:30:00"), "deleted_at": now, "note": "x",
		},
		{
			"id": json.Number("8"), "order_id": nil, "is_cod": "0", "cod_amount": json.Number("1.25"),
			"count": float64(4), "inp_time": "2026-10-29 12:30:00.123", "updated_at": "2026-10-29T12:30:00Z",
		},
	}
	var out []scanModel
	if err := scanRowsInto(rows, &out); err != nil {
		t.Fatal(err)
	}
	first := out[0]
	if first.ID != 7 || first.OrderID == nil || *first.OrderID != 5 || !first.IsCod || first.CodAmount != 10.5 ||
		first.Weight == nil || *first.Weight != 2.5 || first.Count != 3 || first.Name != "LT" || first.Code != "42" ||
		first.CreatedAt == nil || !first.CreatedAt.Equal(now) || !first.UpdatedAt.Equal(now) ||
		!first.DeletedAt.Valid || first.Note == nil || *first.Note != "x" {
		t.Fatalf("first row: %+v", first)
	}
	second := out[1]
	if second.ID != 8 || second.OrderID != nil || second.IsCod || second.CodAmount != 1.25 || second.Count != 4 ||
		second.CreatedAt == nil || second.UpdatedAt.IsZero() || second.DeletedAt.Valid {
		t.Fatalf("second row: %+v", second)
	}
}

func TestScanRowsIntoLeavesUnreadableValuesEmpty(t *testing.T) {
	var out []scanModel
	err := scanRowsInto([]map[string]interface{}{{"id": int64(3), "cod_amount": []byte("ten"), "inp_time": "not a time"}}, &out)
	if err != nil || len(out) != 1 || out[0].ID != 3 || out[0].CodAmount != 0 {
		t.Fatalf("an unreadable value must leave its field empty and keep the row: %+v %v", out, err)
	}
	var zero scanModel
	if err := scanRowsInto([]map[string]interface{}{{"inp_time": "0000-00-00 00:00:00", "updated_at": []byte("0000-00-00"), "order_id": "", "count": ""}}, &zero); err != nil {
		t.Fatal(err)
	}
	if zero.CreatedAt != nil || !zero.UpdatedAt.IsZero() || zero.OrderID != nil || zero.Count != 0 {
		t.Fatalf("zero dates and empty strings must read as no value: %+v", zero)
	}
	var loose struct {
		Extra interface{} `gorm:"column:extra"`
		Note  *string     `gorm:"column:note"`
	}
	if err := scanRowsInto([]map[string]interface{}{{"extra": int64(5), "note": ""}}, &loose); err != nil {
		t.Fatal(err)
	}
	if loose.Extra != int64(5) || loose.Note == nil || *loose.Note != "" {
		t.Fatalf("interface fields take the value as is, text keeps an empty string: %+v", loose)
	}
}

func TestScanRowsIntoSingleColumnSlices(t *testing.T) {
	var ids []int
	if err := scanRowsInto([]map[string]interface{}{{"id": int64(1)}, {"id": []byte("2")}}, &ids); err != nil {
		t.Fatal(err)
	}
	if len(ids) != 2 || ids[0] != 1 || ids[1] != 2 {
		t.Fatalf("pluck ids: %v", ids)
	}
	var count int64
	if err := scanRowsInto([]map[string]interface{}{{"count": int64(9)}}, &count); err != nil || count != 9 {
		t.Fatalf("scalar: %v %v", count, err)
	}
}
