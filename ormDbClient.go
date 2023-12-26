package shared

import (
	"fmt"
	"reflect"
	"time"
)

type MigrateField struct {
	FieldName string
	GormTag   string
	FieldType string
}
type TableNameFunc func() string

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

		// Get the value of the specified tag key for the field
		tagValue, _ := field.Tag.Lookup(tagKey)

		migrateFields = append(migrateFields, MigrateField{
			FieldName: field.Name,
			FieldType: field.Type.Name(),
			GormTag:   tagValue,
		})

	}

	return migrateFields, nil
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
	case "time.Time":
		return reflect.TypeOf(time.Time{})
	case "int":
		return reflect.TypeOf(0)
	case "float64":
		return reflect.TypeOf(0.0)
	// Add more types as needed
	default:
		// Default to string type if the field type is not recognized
		return reflect.TypeOf("")
	}
}
