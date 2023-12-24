package shared

type DB interface {
	Migrate(model interface{}) error
	Exec(sql string, values ...interface{}) error
	Raw(sql string, values ...interface{}) ([]map[string]interface{}, error)
}
