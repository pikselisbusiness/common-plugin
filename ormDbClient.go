package shared

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"
)

type MigrateField struct {
	FieldName string
	GormTag   string
	FieldType string
}
type TableNameFunc func() string

// GetMigrateQueries extracts migration fields from a struct model
func GetMigrateQueries(model interface{}) ([]MigrateField, error) {

	tagKey := "gorm"

	migrateFields := make([]MigrateField, 0)

	// Get the type of the input struct
	t := reflect.TypeOf(model)

	// If the input is a pointer, get the underlying type
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// Iterate through the fields of the struct
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Skip unexported fields
		if !field.IsExported() {
			continue
		}

		// Get the value of the specified tag key for the field
		tagValue, _ := field.Tag.Lookup(tagKey)

		// Skip fields marked with gorm:"-"
		if tagValue == "-" {
			continue
		}

		fieldType := getFieldTypeName(field.Type)

		migrateFields = append(migrateFields, MigrateField{
			FieldName: field.Name,
			FieldType: fieldType,
			GormTag:   tagValue,
		})

	}

	return migrateFields, nil
}

// getFieldTypeName returns a string representation of the field type
func getFieldTypeName(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Ptr:
		return "*" + getFieldTypeName(t.Elem())
	case reflect.Slice:
		return "[]" + getFieldTypeName(t.Elem())
	default:
		if t.PkgPath() != "" {
			return t.PkgPath() + "." + t.Name()
		}
		return t.Name()
	}
}

// AutoMigrate is a convenience function that extracts fields and calls MigrateModel
// Usage: shared.AutoMigrate(db, &MyModel{})
func AutoMigrate(db DB, models ...interface{}) error {
	for _, model := range models {
		tableName := GetTableNameFromModel(model)
		fields, err := GetMigrateQueries(model)
		if err != nil {
			return fmt.Errorf("failed to get migrate queries for %s: %w", tableName, err)
		}
		if err := db.MigrateModel(tableName, fields); err != nil {
			return fmt.Errorf("failed to migrate %s: %w", tableName, err)
		}
	}
	return nil
}

// GetTableNameFromModel extracts table name from model
// Checks for TableName() method first, then converts struct name to snake_case
func GetTableNameFromModel(model interface{}) string {
	// Check if model implements TableName() method
	if tabler, ok := model.(interface{ TableName() string }); ok {
		return tabler.TableName()
	}

	// Get type name and convert to snake_case + pluralize
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return toSnakeCasePlural(t.Name())
}

// toSnakeCasePlural converts CamelCase to snake_case and adds 's'
func toSnakeCasePlural(s string) string {
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
	return strings.ToLower(result.String()) + "s"
}

// CreateStructFromFields dynamically creates a new struct based on the provided MigrateField properties.
func CreateStructFromFields(fields []MigrateField, tableName string) interface{} {
	// Define a new struct type
	structType := reflect.StructOf(fieldsToStructFields(fields))

	// Create a new instance of the struct
	instance := reflect.New(structType).Elem()

	// Add TableName method to the struct
	// addTableNameMethod(instance, tableName)

	return instance.Interface()
}

// fieldsToStructFields converts MigrateField to reflect.StructField.
func fieldsToStructFields(fields []MigrateField) []reflect.StructField {
	structFields := make([]reflect.StructField, len(fields))

	for i, field := range fields {
		structFields[i] = reflect.StructField{
			Name: field.FieldName,
			Type: mapFieldType(field.FieldType),
			Tag:  reflect.StructTag(fmt.Sprintf(`%s:"%s"`, "gorm", field.GormTag)),
		}
	}

	return structFields
}

// addTableNameMethod adds a TableName method to the struct.
func addTableNameMethod(structValue reflect.Value, tableName string) {
	// Define the TableName function
	tableNameFunc := func() string {
		return tableName
	}

	// Create a method value with the proper type
	methodValue := reflect.MakeFunc(reflect.TypeOf(tableNameFunc), func(args []reflect.Value) (results []reflect.Value) {
		return []reflect.Value{reflect.ValueOf(tableName)}
	})

	// Set the method on the struct value
	structValue.MethodByName("TableName").Set(methodValue)
}

// mapFieldType maps a string type to its corresponding reflect.Type.
func mapFieldType(fieldType string) reflect.Type {
	switch fieldType {
	case "string":
		return reflect.TypeOf("")
	case "Time":
		return reflect.TypeOf(time.Time{})
	case "time.Time":
		return reflect.TypeOf(time.Time{})
	case "int":
		return reflect.TypeOf(0)
	case "int8":
		return reflect.TypeOf(int8(0))
	case "int16":
		return reflect.TypeOf(int16(0))
	case "int32":
		return reflect.TypeOf(int32(0))
	case "int64":
		return reflect.TypeOf(int64(0))
	case "uint":
		return reflect.TypeOf(uint(0))
	case "uint8":
		return reflect.TypeOf(uint8(0))
	case "uint16":
		return reflect.TypeOf(uint16(0))
	case "uint32":
		return reflect.TypeOf(uint32(0))
	case "uint64":
		return reflect.TypeOf(uint64(0))
	case "float32":
		return reflect.TypeOf(float32(0))
	case "float64":
		return reflect.TypeOf(0.0)
	case "bool":
		return reflect.TypeOf(false)
	case "NullTime":
		return reflect.TypeOf(sql.NullTime{})
	case "database/sql.NullTime":
		return reflect.TypeOf(sql.NullTime{})
	case "NullString":
		return reflect.TypeOf(sql.NullString{})
	case "NullInt64":
		return reflect.TypeOf(sql.NullInt64{})
	case "NullFloat64":
		return reflect.TypeOf(sql.NullFloat64{})
	case "NullBool":
		return reflect.TypeOf(sql.NullBool{})
	// Add more types as needed
	default:
		// Default to string type if the field type is not recognized
		return reflect.TypeOf("")
	}
}
