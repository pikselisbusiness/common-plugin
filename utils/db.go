package utils

import (
	"fmt"
	"reflect"
	"strings"
)

// Parse required tag value for using
// e.g. Gorm primaryKey
func ParseTagSetting(str string, sep string) map[string]string {
	settings := map[string]string{}
	names := strings.Split(str, sep)

	for i := 0; i < len(names); i++ {
		j := i
		if len(names[j]) > 0 {
			for {
				if names[j][len(names[j])-1] == '\\' {
					i++
					names[j] = names[j][0:len(names[j])-1] + sep + names[i]
					names[i] = ""
				} else {
					break
				}
			}
		}

		values := strings.Split(names[j], ":")
		k := strings.TrimSpace(strings.ToUpper(values[0]))

		if len(values) >= 2 {
			settings[k] = strings.Join(values[1:], ":")
		} else if k != "" {
			settings[k] = k
		}
	}

	return settings
}
func mapsToStructs(data []map[string]interface{}, dest interface{}) error {
	destValue := reflect.ValueOf(dest)

	// Make sure the destination is a pointer to a slice of structs
	if destValue.Kind() != reflect.Ptr || destValue.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("destination must be a pointer to a slice of structs")
	}

	destElem := destValue.Elem()
	destType := destElem.Type().Elem()

	for _, mapData := range data {
		// Create a new instance of the struct
		structValue := reflect.New(destType).Elem()

		for i := 0; i < destType.NumField(); i++ {
			field := structValue.Field(i)
			fieldType := destType.Field(i)

			// Get the key for the field
			key := fieldType.Tag.Get("json")
			if key == "" {
				key = fieldType.Name
			}

			// Check if the key exists in the map
			if value, ok := mapData[key]; ok {
				// Convert the map value to the field type
				fieldValue := reflect.ValueOf(value)
				if fieldValue.Type().ConvertibleTo(field.Type()) {
					field.Set(fieldValue.Convert(field.Type()))
				} else {
					return fmt.Errorf("cannot convert map value to field type for key %s", key)
				}
			}
		}

		// Append the struct to the destination slice
		destElem.Set(reflect.Append(destElem, structValue))
	}

	return nil
}
