package shared

// DBv2 is the enhanced database interface for plugins
// Provides GORM-like chainable API over RPC
type DBv2 interface {
	// Core operations (existing)
	MigrateModel(tableName string, fields []MigrateField) error
	Exec(sql string, values ...interface{}) error
	Raw(sql string, values ...interface{}) ([]map[string]interface{}, error)

	// Transaction support
	Begin() (DBv2, error)
	Commit() error
	Rollback() error

	// Query builder - returns chainable interface
	Table(name string) QueryBuilder
	Model(value interface{}) QueryBuilder
}

// QueryBuilder provides chainable query building
type QueryBuilder interface {
	// Selection
	Select(columns ...string) QueryBuilder
	Distinct(columns ...string) QueryBuilder
	Omit(columns ...string) QueryBuilder

	// Conditions
	Where(query string, args ...interface{}) QueryBuilder
	Or(query string, args ...interface{}) QueryBuilder
	Not(query string, args ...interface{}) QueryBuilder

	// Ordering & Pagination
	Order(column string, direction ...OrderDirection) QueryBuilder
	Limit(limit int) QueryBuilder
	Offset(offset int) QueryBuilder

	// Grouping
	Group(columns ...string) QueryBuilder
	Having(query string, args ...interface{}) QueryBuilder

	// Joins
	Join(table string, condition string, args ...interface{}) QueryBuilder
	LeftJoin(table string, condition string, args ...interface{}) QueryBuilder
	RightJoin(table string, condition string, args ...interface{}) QueryBuilder

	// Preloading (for associations)
	Preload(association string, args ...interface{}) QueryBuilder

	// Execution - Read
	Find(dest interface{}) QueryResult
	First(dest interface{}) QueryResult
	Last(dest interface{}) QueryResult
	Take(dest interface{}) QueryResult
	Count(count *int64) QueryResult
	Pluck(column string, dest interface{}) QueryResult
	Scan(dest interface{}) QueryResult

	// Execution - Write
	Create(value interface{}) QueryResult
	Save(value interface{}) QueryResult
	Update(column string, value interface{}) QueryResult
	Updates(values interface{}) QueryResult
	Delete(value interface{}, conds ...interface{}) QueryResult

	// Raw SQL within chain
	Raw(sql string, values ...interface{}) QueryBuilder

	// Scopes for reusable query logic
	Scopes(funcs ...func(QueryBuilder) QueryBuilder) QueryBuilder

	// Debug mode
	Debug() QueryBuilder

	// Get the built SQL (for debugging)
	ToSQL() (string, []interface{})
}
