package shared

import (
	"strings"
	"testing"
)

type upsertTestModel struct {
	ID  uint   `gorm:"primaryKey;column:id"`
	SKU string `gorm:"column:sku;uniqueIndex"`
}

func TestPostgresUpsertUsesUniqueIndexWhenIDIsZero(t *testing.T) {
	t.Setenv("DATABASE_DRIVER", "postgres")

	model := upsertTestModel{SKU: "ABC"}
	conflictColumns := detectConflictColumns(model, []string{"sku"})
	if len(conflictColumns) != 1 || conflictColumns[0] != "sku" {
		t.Fatalf("expected sku conflict column, got %#v", conflictColumns)
	}

	sql, err := buildUpsertSQL("products", []string{"sku"}, []string{"?"}, []string{`"sku" = EXCLUDED."sku"`}, conflictColumns)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(sql, `ON CONFLICT ("sku") DO UPDATE`) {
		t.Fatalf("expected postgres ON CONFLICT SQL, got %s", sql)
	}
}

func TestMysqlUpsertKeepsDuplicateKeySyntax(t *testing.T) {
	t.Setenv("DATABASE_DRIVER", "mysql")

	sql, err := buildUpsertSQL("products", []string{"sku"}, []string{"?"}, []string{"sku = VALUES(sku)"}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(sql, "ON DUPLICATE KEY UPDATE") {
		t.Fatalf("expected mysql duplicate key SQL, got %s", sql)
	}
}
