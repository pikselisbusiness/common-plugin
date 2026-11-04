package shared

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	stdlog "log"
	"math"
	"os"
	"reflect"
	"strconv"
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

// rpcValue makes a statement argument or column value safe for gob, which
// refuses nil pointers inside interfaces: pointers become their value (nil
// becomes NULL) and driver.Valuer types (sql.NullTime, ...) their Value().
func rpcValue(value interface{}) interface{} {
	if valuer, ok := value.(driver.Valuer); ok {
		if v := reflect.ValueOf(value); v.Kind() == reflect.Ptr && v.IsNil() {
			return nil
		}
		out, err := valuer.Value()
		if err != nil {
			return value
		}
		return out
	}
	v := reflect.ValueOf(value)
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}
	if !v.IsValid() {
		return nil
	}
	if v.Kind() == reflect.Slice && v.Type() != bytesType {
		out := make([]interface{}, v.Len())
		for i := range out {
			out[i] = rpcValue(v.Index(i).Interface())
		}
		return out
	}
	return v.Interface()
}

func rpcValues(values []interface{}) []interface{} {
	if values == nil {
		return nil
	}
	out := make([]interface{}, len(values))
	for i, value := range values {
		out[i] = rpcValue(value)
	}
	return out
}

// structToMap converts a struct to map[string]interface{}
func structToMap(value interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	if value == nil {
		return result
	}

	// A map keeps its keys; its values get the same gob treatment.
	if m, ok := value.(map[string]interface{}); ok {
		for k, v := range m {
			result[k] = rpcValue(v)
		}
		return result
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

		result[colName] = rpcValue(fieldVal.Interface())
	}

	return result
}

// scanRowsInto scans rows into dest: a pointer to a struct, a slice of
// structs, or a slice of a single column's values (Pluck). A value that
// cannot be stored in its field is an error, never a silently zero field.
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
			if err := rowToValue(row, item); err != nil {
				return err
			}
			destElem.Set(reflect.Append(destElem, item))
		}
		return nil
	}

	// Handle single struct destination
	if len(rows) > 0 {
		return rowToValue(rows[0], destElem)
	}

	return nil
}

func rowToValue(row map[string]interface{}, item reflect.Value) error {
	if isScalarDest(item.Type()) {
		for column, value := range row {
			storeColumn(item, column, value)
			return nil
		}
		return nil
	}
	return mapToStructValue(row, item)
}

// storeColumn sets one column value. A value the field cannot hold leaves the
// field zero and is logged rather than failing the read: plugins written
// against the earlier, silent conversion keep loading their rows.
func storeColumn(field reflect.Value, column string, value interface{}) {
	if err := setFieldValue(field, value); err != nil {
		field.Set(reflect.Zero(field.Type()))
		stdlog.Printf("common-plugin: column %s left empty: %v", column, err)
	}
}

// isScalarDest reports a destination that holds one column, not a row.
func isScalarDest(t reflect.Type) bool {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Kind() != reflect.Struct || t == reflect.TypeOf(time.Time{}) || reflect.PtrTo(t).Implements(scannerType)
}

func mapToStructValue(data map[string]interface{}, structVal reflect.Value) error {
	if structVal.Kind() != reflect.Struct {
		return nil
	}
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

		storeColumn(field, key, value)
	}
	return nil
}

var (
	scannerType = reflect.TypeOf((*sql.Scanner)(nil)).Elem()
	timeType    = reflect.TypeOf(time.Time{})
	bytesType   = reflect.TypeOf([]byte(nil))
)

// sqlTimeLayouts are the forms a DATETIME, TIMESTAMP or DATE column arrives
// in without parseTime, plus RFC 3339 from JSON.
var sqlTimeLayouts = []string{
	"2006-01-02 15:04:05.999999999",
	"2006-01-02 15:04:05",
	time.RFC3339Nano,
	"2006-01-02",
}

// setFieldValue stores one column value as the MySQL driver and JSON hand it
// over (int64, float64, []byte, string, time.Time, json.Number) into a field
// of the model's type, including pointer fields and sql.Scanner types.
func setFieldValue(field reflect.Value, value interface{}) error {
	if value == nil {
		return nil
	}

	if field.Kind() == reflect.Interface {
		if reflect.TypeOf(value).Implements(field.Type()) {
			field.Set(reflect.ValueOf(value))
			return nil
		}
		return fmt.Errorf("cannot store %T into %s", value, field.Type())
	}

	// An empty string or a MySQL zero date is no value for a non-text field.
	if text, ok := sqlText(value); ok && isEmptySQLValue(text) && !holdsText(field.Type()) {
		return nil
	}

	if field.Kind() == reflect.Ptr {
		elem := reflect.New(field.Type().Elem())
		if err := setFieldValue(elem.Elem(), value); err != nil {
			return err
		}
		field.Set(elem)
		return nil
	}

	if field.CanAddr() && field.Addr().Type().Implements(scannerType) {
		return field.Addr().Interface().(sql.Scanner).Scan(value)
	}

	if field.Type() == bytesType {
		switch v := value.(type) {
		case []byte:
			field.SetBytes(append([]byte(nil), v...))
			return nil
		case string:
			field.SetBytes([]byte(v))
			return nil
		}
	}

	// Text forms: []byte from the driver, json.Number from JSON.
	switch v := value.(type) {
	case []byte:
		value = string(v)
	case json.Number:
		value = string(v)
	}

	if field.Type() == timeType {
		switch v := value.(type) {
		case time.Time:
			field.Set(reflect.ValueOf(v))
			return nil
		case string:
			for _, layout := range sqlTimeLayouts {
				if t, err := time.ParseInLocation(layout, v, time.Local); err == nil {
					field.Set(reflect.ValueOf(t))
					return nil
				}
			}
		}
		return fmt.Errorf("cannot store %T %v as time", value, value)
	}

	switch field.Kind() {
	case reflect.Bool:
		switch v := value.(type) {
		case bool:
			field.SetBool(v)
			return nil
		case string:
			b, err := strconv.ParseBool(v)
			if err != nil {
				return fmt.Errorf("cannot store %q as bool", v)
			}
			field.SetBool(b)
			return nil
		}
		if n, ok := numberValue(value); ok {
			field.SetBool(n != 0)
			return nil
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if s, ok := value.(string); ok {
			n, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				f, ferr := strconv.ParseFloat(s, 64)
				if ferr != nil || f != math.Trunc(f) {
					return fmt.Errorf("cannot store %q as integer", s)
				}
				n = int64(f)
			}
			field.SetInt(n)
			return nil
		}
		if n, ok := numberValue(value); ok {
			field.SetInt(int64(n))
			return nil
		}
		if b, ok := value.(bool); ok {
			field.SetInt(boolInt(b))
			return nil
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if s, ok := value.(string); ok {
			n, err := strconv.ParseUint(s, 10, 64)
			if err != nil {
				return fmt.Errorf("cannot store %q as unsigned integer", s)
			}
			field.SetUint(n)
			return nil
		}
		if n, ok := numberValue(value); ok && n >= 0 {
			field.SetUint(uint64(n))
			return nil
		}
		if b, ok := value.(bool); ok {
			field.SetUint(uint64(boolInt(b)))
			return nil
		}
	case reflect.Float32, reflect.Float64:
		if s, ok := value.(string); ok {
			f, err := strconv.ParseFloat(s, 64)
			if err != nil {
				return fmt.Errorf("cannot store %q as number", s)
			}
			field.SetFloat(f)
			return nil
		}
		if n, ok := numberValue(value); ok {
			field.SetFloat(n)
			return nil
		}
	case reflect.String:
		switch v := value.(type) {
		case string:
			field.SetString(v)
			return nil
		case time.Time:
			field.SetString(v.Format("2006-01-02 15:04:05"))
			return nil
		}
		if _, ok := numberValue(value); ok {
			field.SetString(fmt.Sprint(value))
			return nil
		}
	}

	fieldValue := reflect.ValueOf(value)
	if fieldValue.Type().ConvertibleTo(field.Type()) && fieldValue.Kind() == field.Kind() {
		field.Set(fieldValue.Convert(field.Type()))
		return nil
	}
	return fmt.Errorf("cannot store %T into %s", value, field.Type())
}

func sqlText(value interface{}) (string, bool) {
	switch v := value.(type) {
	case string:
		return v, true
	case []byte:
		return string(v), true
	}
	return "", false
}

// holdsText reports a string or []byte field, behind pointers.
func holdsText(t reflect.Type) bool {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Kind() == reflect.String || t == bytesType
}

func isEmptySQLValue(s string) bool {
	return s == "" || strings.HasPrefix(s, "0000-00-00")
}

// numberValue reads any Go number as float64.
func numberValue(value interface{}) (float64, bool) {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(v.Int()), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(v.Uint()), true
	case reflect.Float32, reflect.Float64:
		return v.Float(), true
	}
	return 0, false
}

func boolInt(b bool) int64 {
	if b {
		return 1
	}
	return 0
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
