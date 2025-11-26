package shared

import (
	"fmt"
	"reflect"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/pikselisbusiness/common-plugin/utils"
)

var (
	tagKey string = "gorm"
)

type DbBuilder interface {
	Where(query string, params interface{}) DbBuilder
	Select(query string) DbBuilder
	Find(target interface{}) DbBuilder
}
type Result struct {
	Error        error
	RowsAffected int
}

type DbTransaction struct {
	db              DB
	logger          hclog.Logger
	selectQuery     string
	whereQuery      string
	tableName       string
	whereParameters []interface{}
}

type DbBuilderImplementation struct {
	db     DB
	logger hclog.Logger
}

func NewDbBuilder(db DB, logger hclog.Logger) *DbBuilderImplementation {
	return &DbBuilderImplementation{db: db, logger: logger}
}

func (d *DbBuilderImplementation) Table(table string) *DbTransaction {

	tx := &DbTransaction{
		tableName: table,
		db:        d.db,
		logger:    d.logger,
	}

	return tx
}
func (d *DbTransaction) Where(query string, params ...interface{}) *DbTransaction {

	d.whereQuery = query
	d.whereParameters = params
	return d
}
func (d *DbTransaction) Select(query string) *DbTransaction {

	d.selectQuery = query
	return d
}
func (d *DbTransaction) Find(dest interface{}) Result {

	query := "SELECT " + d.selectQuery + " FROM " + d.tableName

	params := make([]interface{}, 0)
	if d.whereQuery != "" {
		query += " WHERE " + d.whereQuery
		params = d.whereParameters
	}
	items, err := d.db.Raw(query, params...)
	if err != nil {
		return Result{
			Error: err,
		}
	}
	destValue := reflect.ValueOf(dest)

	// Make sure the destination is a pointer to a slice of structs
	if destValue.Kind() != reflect.Ptr || destValue.Elem().Kind() != reflect.Slice {
		return Result{
			Error: fmt.Errorf("destination must be a pointer to a struct"),
		}
	}

	destElem := destValue.Elem()
	destType := destElem.Type().Elem()

	for _, mapData := range items {
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
				if value != nil {
					if fieldValue.Type().ConvertibleTo(field.Type()) {
						field.Set(fieldValue.Convert(field.Type()))
					} else {

						// for time.Time field inequalities - convert to time.Time
						if field.Type().Name() == "Time" && fieldValue.Type().Name() == "string" {

							timeValue, err := time.Parse("2006-01-02 15:04:05", value.(string))
							if err != nil {
								return Result{
									Error: err,
								}
							}
							fieldValue := reflect.ValueOf(timeValue)
							field.Set(fieldValue.Convert(field.Type()))
						} else {

							return Result{
								Error: fmt.Errorf("cannot convert map value to field type for key %s, needed type %v, received type %v", key, field.Type().Name(), fieldValue.Type()),
							}
						}
					}
				}
			}
		}

		// Append the struct to the destination slice
		destElem.Set(reflect.Append(destElem, structValue))
	}

	return Result{
		Error:        nil,
		RowsAffected: 0,
	}
}
func (d *DbTransaction) Delete(model interface{}) Result {

	// Get the type and value of the input struct
	t := reflect.TypeOf(model)
	val := reflect.ValueOf(model)

	// If the input is a pointer, get the underlying type
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	primaryField := "id"
	var primaryValue interface{}

	// Iterate through the fields of the struct
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		valueField := val.Field(i)

		d.logger.Info("value", "value", valueField)

		val := reflect.ValueOf(valueField)

		// get all tag settings for tagKey separated by ;
		tagSetting := utils.ParseTagSetting(field.Tag.Get(tagKey), ";")
		isPrimaryKey := utils.CheckTruth(tagSetting["PRIMARYKEY"], tagSetting["PRIMARY_KEY"])

		// Get the value of the specified tag key for the field
		if isPrimaryKey {
			primaryField = tagSetting["COLUMN"]
			primaryValue = val.Interface()
		}
	}

	deleteQuery := "DELETE FROM " + d.tableName + " WHERE " + primaryField + " = ? "

	d.logger.Info("Delete query", "query", deleteQuery, "primary value", primaryValue)
	err := d.db.Exec(deleteQuery, primaryValue)

	return Result{
		Error: err,
	}
}

func testBuilder(db DB) {

	builder := NewDbBuilder(db, nil)

	type ItemTest struct {
		Balance string
	}

	var items []ItemTest

	builder.Table("test").Select("balance, name_field").Where("id = @id", map[string]interface{}{"id": 1}).Find(&items)

}
