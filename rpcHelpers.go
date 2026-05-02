package shared

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"
	"time"
)

// setStructID sets the ID field on a struct pointer after insert
// Supports ID, Id field names and uint, uint64, int, int64 types
func setStructID(value interface{}, id int64) {
	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Ptr {
		return
	}
	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return
	}

	// Try common ID field names
	idField := v.FieldByName("ID")
	if !idField.IsValid() {
		idField = v.FieldByName("Id")
	}
	if !idField.IsValid() {
		// Try to find field with gorm:"primaryKey" tag
		t := v.Type()
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			gormTag := field.Tag.Get("gorm")
			if strings.Contains(gormTag, "primaryKey") || strings.Contains(gormTag, "primary_key") {
				idField = v.Field(i)
				break
			}
		}
	}

	if !idField.IsValid() || !idField.CanSet() {
		return
	}

	// Set the value based on field type
	switch idField.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		idField.SetUint(uint64(id))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		idField.SetInt(id)
	}
}

// structToMap converts a struct to map[string]interface{}
func structToMap(value interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	if value == nil {
		return result
	}

	// If already a map, return it
	if m, ok := value.(map[string]interface{}); ok {
		return m
	}

	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return result
	}

	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldVal := v.Field(i)

		if !fieldVal.CanInterface() {
			continue
		}

		// Get column name
		colName := getColumnNameFromField(field)
		if colName == "-" || colName == "" {
			continue
		}

		result[colName] = fieldVal.Interface()
	}

	return result
}

// scanRowsInto scans rows into destination
func scanRowsInto(rows []map[string]interface{}, dest interface{}) error {
	destValue := reflect.ValueOf(dest)
	if destValue.Kind() != reflect.Ptr {
		return nil
	}

	destElem := destValue.Elem()

	// Handle slice destination
	if destElem.Kind() == reflect.Slice {
		destType := destElem.Type().Elem()
		for _, row := range rows {
			item := reflect.New(destType).Elem()
			mapToStructValue(row, item)
			destElem.Set(reflect.Append(destElem, item))
		}
		return nil
	}

	// Handle single struct destination
	if destElem.Kind() == reflect.Struct && len(rows) > 0 {
		mapToStructValue(rows[0], destElem)
	}

	return nil
}

func mapToStructValue(data map[string]interface{}, structVal reflect.Value) {
	structType := structVal.Type()

	for i := 0; i < structType.NumField(); i++ {
		field := structVal.Field(i)
		fieldType := structType.Field(i)

		key := getColumnNameFromField(fieldType)
		value, ok := data[key]
		if !ok || value == nil {
			continue
		}

		if !field.CanSet() {
			continue
		}

		setFieldValue(field, value)
	}
}

func setFieldValue(field reflect.Value, value interface{}) {
	fieldValue := reflect.ValueOf(value)

	// Handle json.Number
	if num, ok := value.(json.Number); ok {
		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if n, err := num.Int64(); err == nil {
				field.SetInt(n)
				return
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if n, err := num.Int64(); err == nil {
				field.SetUint(uint64(n))
				return
			}
		case reflect.Float32, reflect.Float64:
			if f, err := num.Float64(); err == nil {
				field.SetFloat(f)
				return
			}
		}
	}

	// Handle float64 to int conversion (common in JSON)
	if f, ok := value.(float64); ok {
		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			field.SetInt(int64(f))
			return
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			field.SetUint(uint64(f))
			return
		}
	}

	// Handle time.Time
	if field.Type() == reflect.TypeOf(time.Time{}) {
		switch v := value.(type) {
		case string:
			if t, err := time.Parse("2006-01-02 15:04:05", v); err == nil {
				field.Set(reflect.ValueOf(t))
				return
			}
			if t, err := time.Parse(time.RFC3339, v); err == nil {
				field.Set(reflect.ValueOf(t))
				return
			}
		case time.Time:
			field.Set(reflect.ValueOf(v))
			return
		}
	}

	// Direct conversion
	if fieldValue.Type().ConvertibleTo(field.Type()) {
		field.Set(fieldValue.Convert(field.Type()))
	}
}

func getColumnNameFromField(field reflect.StructField) string {
	// Check gorm tag first
	gormTag := field.Tag.Get("gorm")
	if gormTag != "" {
		if gormTag == "-" {
			return "-"
		}
		if idx := strings.Index(gormTag, "column:"); idx >= 0 {
			name := gormTag[idx+7:]
			if end := strings.Index(name, ";"); end >= 0 {
				name = name[:end]
			}
			return name
		}
	}

	// Check json tag
	jsonTag := field.Tag.Get("json")
	if jsonTag != "" && jsonTag != "-" {
		if idx := strings.Index(jsonTag, ","); idx >= 0 {
			return jsonTag[:idx]
		}
		return jsonTag
	}

	// Default to snake_case
	return toSnakeCaseHelper(field.Name)
}

func currentSQLDialect() string {
	dialect := strings.ToLower(os.Getenv("DATABASE_DRIVER"))
	if dialect == "" {
		dialect = strings.ToLower(os.Getenv("DB_DRIVER"))
	}
	if dialect == "" {
		dialect = strings.ToLower(os.Getenv("DATABASE_DIALECT"))
	}
	switch dialect {
	case "postgres", "postgresql", "pg":
		return "postgres"
	default:
		return "mysql"
	}
}

func isPostgresDialect() bool {
	return currentSQLDialect() == "postgres"
}

func quoteSQLIdentifier(identifier string) string {
	quote := "`"
	if isPostgresDialect() {
		quote = `"`
	}
	parts := strings.Split(identifier, ".")
	for i, part := range parts {
		if part == "*" || part == "" {
			continue
		}
		if strings.HasPrefix(part, quote) && strings.HasSuffix(part, quote) {
			continue
		}
		parts[i] = quote + strings.ReplaceAll(part, quote, quote+quote) + quote
	}
	return strings.Join(parts, ".")
}

func detectConflictColumns(value interface{}, columns []string) []string {
	columnSet := make(map[string]bool, len(columns))
	for _, col := range columns {
		columnSet[col] = true
	}

	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		if columnSet["id"] {
			return []string{"id"}
		}
		return nil
	}

	t := v.Type()
	var uniqueColumn string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		col := getColumnNameFromField(field)
		if col == "-" || !columnSet[col] {
			continue
		}
		gormTag := field.Tag.Get("gorm")
		if strings.Contains(gormTag, "primaryKey") || strings.Contains(gormTag, "primary_key") {
			fieldValue := v.Field(i)
			if fieldValue.CanInterface() && !fieldValue.IsZero() {
				return []string{col}
			}
		}
		if uniqueColumn == "" && (strings.Contains(gormTag, "uniqueIndex") || strings.Contains(gormTag, "unique")) {
			uniqueColumn = col
		}
	}
	if uniqueColumn != "" {
		return []string{uniqueColumn}
	}
	return nil
}

func buildUpsertSQL(tableName string, columns, placeholders, updateColumns, conflictColumns []string) (string, error) {
	if !isPostgresDialect() {
		return "INSERT INTO " + tableName + " (" + strings.Join(columns, ", ") + ") VALUES (" + strings.Join(placeholders, ", ") + ") ON DUPLICATE KEY UPDATE " + strings.Join(updateColumns, ", "), nil
	}
	if len(conflictColumns) == 0 {
		return "", &UpsertConflictError{}
	}

	quotedConflictColumns := make([]string, len(conflictColumns))
	for i, col := range conflictColumns {
		quotedConflictColumns[i] = quoteSQLIdentifier(col)
	}

	conflictSQL := " ON CONFLICT (" + strings.Join(quotedConflictColumns, ", ") + ")"
	if len(updateColumns) == 0 {
		conflictSQL += " DO NOTHING"
	} else {
		conflictSQL += " DO UPDATE SET " + strings.Join(updateColumns, ", ")
	}
	return "INSERT INTO " + tableName + " (" + strings.Join(columns, ", ") + ") VALUES (" + strings.Join(placeholders, ", ") + ")" + conflictSQL, nil
}

type UpsertConflictError struct{}

func (e *UpsertConflictError) Error() string {
	return "postgres upsert requires an id value or a struct field tagged with gorm uniqueIndex/unique"
}

func toSnakeCaseHelper(s string) string {
	runes := []rune(s)
	var result strings.Builder
	for i, r := range runes {
		if i > 0 && r >= 'A' && r <= 'Z' {
			prevUpper := runes[i-1] >= 'A' && runes[i-1] <= 'Z'
			nextLower := i+1 < len(runes) && runes[i+1] >= 'a' && runes[i+1] <= 'z'
			if !prevUpper || nextLower {
				result.WriteByte('_')
			}
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}
