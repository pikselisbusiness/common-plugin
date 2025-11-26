package shared

type DB interface {
	MigrateModel(tableName string, fields []MigrateField) error
	Exec(sql string, values ...interface{}) error
	Raw(sql string, values ...interface{}) ([]map[string]interface{}, error)
}
