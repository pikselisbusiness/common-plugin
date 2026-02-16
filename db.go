package shared

type DB interface {
	MigrateModel(tableName string, fields []MigrateField) error
	Exec(sql string, values ...interface{}) error
	// ExecWithResult executes SQL and returns affected rows and last insert ID
	ExecWithResult(sql string, values ...interface{}) (rowsAffected int64, lastInsertID int64, err error)
	Raw(sql string, values ...interface{}) ([]map[string]interface{}, error)
}
