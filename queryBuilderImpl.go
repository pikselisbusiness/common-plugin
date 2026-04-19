package shared

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/go-hclog"
)

// QueryBuilderImpl implements QueryBuilder interface
type QueryBuilderImpl struct {
	db        DB
	logger    hclog.Logger
	tableName string
	clauses   []Clause
	debug     bool

	// Query parts
	selectCols   []string
	distinctOn   []string
	omitCols     []string
	whereConds   []whereCondition
	orConds      []whereCondition
	notConds     []whereCondition
	orderClauses []orderClause
	groupCols    []string
	havingCond   *whereCondition
	joins        []joinClause
	preloads     []preloadClause
	limitVal     *int
	offsetVal    *int
	rawSQL       string
	rawArgs      []interface{}
}

type whereCondition struct {
	query string
	args  []interface{}
}

type orderClause struct {
	column    string
	direction OrderDirection
}

type joinClause struct {
	joinType  JoinType
	table     string
	condition string
	args      []interface{}
}

type preloadClause struct {
	association string
	args        []interface{}
}

// NewQueryBuilder creates a new query builder
func NewQueryBuilder(db DB, logger hclog.Logger, tableName string) *QueryBuilderImpl {
	return &QueryBuilderImpl{
		db:        db,
		logger:    logger,
		tableName: tableName,
	}
}

// Select specifies columns to retrieve
func (q *QueryBuilderImpl) Select(columns ...string) QueryBuilder {
	q.selectCols = append(q.selectCols, columns...)
	return q
}

// Distinct specifies distinct columns
func (q *QueryBuilderImpl) Distinct(columns ...string) QueryBuilder {
	q.distinctOn = append(q.distinctOn, columns...)
	return q
}

// Omit specifies columns to omit
func (q *QueryBuilderImpl) Omit(columns ...string) QueryBuilder {
	q.omitCols = append(q.omitCols, columns...)
	return q
}

// Where adds a WHERE condition
func (q *QueryBuilderImpl) Where(query string, args ...interface{}) QueryBuilder {
	q.whereConds = append(q.whereConds, whereCondition{query: query, args: args})
	return q
}

// Or adds an OR condition
func (q *QueryBuilderImpl) Or(query string, args ...interface{}) QueryBuilder {
	q.orConds = append(q.orConds, whereCondition{query: query, args: args})
	return q
}

// Not adds a NOT condition
func (q *QueryBuilderImpl) Not(query string, args ...interface{}) QueryBuilder {
	q.notConds = append(q.notConds, whereCondition{query: query, args: args})
	return q
}

// Order adds ordering.
// Supports both GORM-style .Order("col DESC") and explicit .Order("col", shared.DESC).
func (q *QueryBuilderImpl) Order(column string, direction ...OrderDirection) QueryBuilder {
	// If direction is embedded in the column string (GORM style), parse it out
	if len(direction) == 0 {
		upper := strings.ToUpper(strings.TrimSpace(column))
		if strings.HasSuffix(upper, " ASC") || strings.HasSuffix(upper, " DESC") {
			// Use the raw string as-is, no appended direction
			q.orderClauses = append(q.orderClauses, orderClause{column: column, direction: ""})
			return q
		}
	}

	dir := ASC
	if len(direction) > 0 {
		dir = direction[0]
	}
	q.orderClauses = append(q.orderClauses, orderClause{column: column, direction: dir})
	return q
}

// Limit sets the limit
func (q *QueryBuilderImpl) Limit(limit int) QueryBuilder {
	q.limitVal = &limit
	return q
}

// Offset sets the offset
func (q *QueryBuilderImpl) Offset(offset int) QueryBuilder {
	q.offsetVal = &offset
	return q
}

// Group adds GROUP BY
func (q *QueryBuilderImpl) Group(columns ...string) QueryBuilder {
	q.groupCols = append(q.groupCols, columns...)
	return q
}

// Having adds HAVING clause
func (q *QueryBuilderImpl) Having(query string, args ...interface{}) QueryBuilder {
	q.havingCond = &whereCondition{query: query, args: args}
	return q
}

// Join adds INNER JOIN
func (q *QueryBuilderImpl) Join(table string, condition string, args ...interface{}) QueryBuilder {
	q.joins = append(q.joins, joinClause{joinType: InnerJoin, table: table, condition: condition, args: args})
	return q
}

// LeftJoin adds LEFT JOIN
func (q *QueryBuilderImpl) LeftJoin(table string, condition string, args ...interface{}) QueryBuilder {
	q.joins = append(q.joins, joinClause{joinType: LeftJoin, table: table, condition: condition, args: args})
	return q
}

// RightJoin adds RIGHT JOIN
func (q *QueryBuilderImpl) RightJoin(table string, condition string, args ...interface{}) QueryBuilder {
	q.joins = append(q.joins, joinClause{joinType: RightJoin, table: table, condition: condition, args: args})
	return q
}

// Preload adds association preloading
func (q *QueryBuilderImpl) Preload(association string, args ...interface{}) QueryBuilder {
	q.preloads = append(q.preloads, preloadClause{association: association, args: args})
	return q
}

// Raw sets raw SQL
func (q *QueryBuilderImpl) Raw(sql string, values ...interface{}) QueryBuilder {
	q.rawSQL = sql
	q.rawArgs = values
	return q
}

// Scopes applies scope functions
func (q *QueryBuilderImpl) Scopes(funcs ...func(QueryBuilder) QueryBuilder) QueryBuilder {
	var result QueryBuilder = q
	for _, f := range funcs {
		result = f(result)
	}
	return result
}

// Debug enables debug mode
func (q *QueryBuilderImpl) Debug() QueryBuilder {
	q.debug = true
	return q
}

// ToSQL builds and returns the SQL query
func (q *QueryBuilderImpl) ToSQL() (string, []interface{}) {
	if q.rawSQL != "" {
		return q.rawSQL, q.rawArgs
	}
	return q.buildSelectSQL()
}

func (q *QueryBuilderImpl) buildSelectSQL() (string, []interface{}) {
	var sql strings.Builder
	var args []interface{}

	// SELECT
	sql.WriteString("SELECT ")
	if len(q.distinctOn) > 0 {
		sql.WriteString("DISTINCT ")
		sql.WriteString(strings.Join(q.distinctOn, ", "))
	} else if len(q.selectCols) > 0 {
		sql.WriteString(strings.Join(q.selectCols, ", "))
	} else {
		sql.WriteString("*")
	}

	// FROM
	sql.WriteString(" FROM ")
	sql.WriteString(q.tableName)

	// JOINS
	for _, j := range q.joins {
		sql.WriteString(fmt.Sprintf(" %s JOIN %s ON %s", j.joinType, j.table, j.condition))
		args = append(args, j.args...)
	}

	// WHERE
	whereParts, whereArgs := q.buildWhereClauses()
	if len(whereParts) > 0 {
		sql.WriteString(" WHERE ")
		sql.WriteString(strings.Join(whereParts, " AND "))
		args = append(args, whereArgs...)
	}

	// GROUP BY
	if len(q.groupCols) > 0 {
		sql.WriteString(" GROUP BY ")
		sql.WriteString(strings.Join(q.groupCols, ", "))
	}

	// HAVING
	if q.havingCond != nil {
		sql.WriteString(" HAVING ")
		sql.WriteString(q.havingCond.query)
		args = append(args, q.havingCond.args...)
	}

	// ORDER BY
	if len(q.orderClauses) > 0 {
		sql.WriteString(" ORDER BY ")
		orderParts := make([]string, len(q.orderClauses))
		for i, o := range q.orderClauses {
			if o.direction == "" {
				orderParts[i] = o.column
			} else {
				orderParts[i] = fmt.Sprintf("%s %s", o.column, o.direction)
			}
		}
		sql.WriteString(strings.Join(orderParts, ", "))
	}

	// LIMIT
	if q.limitVal != nil {
		sql.WriteString(fmt.Sprintf(" LIMIT %d", *q.limitVal))
	}

	// OFFSET
	if q.offsetVal != nil {
		sql.WriteString(fmt.Sprintf(" OFFSET %d", *q.offsetVal))
	}

	return sql.String(), args
}

func (q *QueryBuilderImpl) buildWhereClauses() ([]string, []interface{}) {
	var parts []string
	var args []interface{}

	for _, w := range q.whereConds {
		parts = append(parts, "("+w.query+")")
		args = append(args, w.args...)
	}

	for _, w := range q.notConds {
		parts = append(parts, "NOT ("+w.query+")")
		args = append(args, w.args...)
	}

	// OR conditions are grouped
	if len(q.orConds) > 0 {
		orParts := make([]string, len(q.orConds))
		for i, w := range q.orConds {
			orParts[i] = "(" + w.query + ")"
			args = append(args, w.args...)
		}
		parts = append(parts, "("+strings.Join(orParts, " OR ")+")")
	}

	return parts, args
}

// Find retrieves multiple records
func (q *QueryBuilderImpl) Find(dest interface{}) QueryResult {
	sql, args := q.ToSQL()
	if q.debug && q.logger != nil {
		q.logger.Info("SQL", "query", sql, "args", args)
	}

	rows, err := q.db.Raw(sql, args...)
	if err != nil {
		return QueryResult{Error: err}
	}

	if err := q.scanRows(rows, dest); err != nil {
		return QueryResult{Error: err}
	}

	return QueryResult{RowsAffected: int64(len(rows))}
}

// First retrieves the first record.
// If no Order has been set, it orders by the struct's primary key (gorm:"primaryKey" tag)
// or falls back to "id".
func (q *QueryBuilderImpl) First(dest interface{}) QueryResult {
	if len(q.orderClauses) == 0 {
		q.Order(q.detectPrimaryKey(dest), ASC)
	}
	q.Limit(1)
	return q.Find(dest)
}

// Last retrieves the last record.
// If no Order has been set, it orders by the struct's primary key (gorm:"primaryKey" tag)
// or falls back to "id".
func (q *QueryBuilderImpl) Last(dest interface{}) QueryResult {
	if len(q.orderClauses) == 0 {
		q.Order(q.detectPrimaryKey(dest), DESC)
	}
	q.Limit(1)
	return q.Find(dest)
}

// Take retrieves one record without ordering
func (q *QueryBuilderImpl) Take(dest interface{}) QueryResult {
	q.Limit(1)
	return q.Find(dest)
}

// Count returns the count of records without mutating the query builder,
// so the same builder can be reused for subsequent queries.
func (q *QueryBuilderImpl) Count(count *int64) QueryResult {
	// Save and restore selectCols so Count doesn't mutate the builder
	origSelect := q.selectCols
	q.selectCols = []string{"COUNT(*) as count"}
	sql, args := q.ToSQL()
	q.selectCols = origSelect

	rows, err := q.db.Raw(sql, args...)
	if err != nil {
		return QueryResult{Error: err}
	}

	if len(rows) > 0 {
		if c, ok := rows[0]["count"]; ok {
			switch v := c.(type) {
			case int64:
				*count = v
			case int:
				*count = int64(v)
			case float64:
				*count = int64(v)
			}
		}
	}

	return QueryResult{}
}

// Pluck retrieves a single column
func (q *QueryBuilderImpl) Pluck(column string, dest interface{}) QueryResult {
	q.selectCols = []string{column}
	sql, args := q.ToSQL()

	rows, err := q.db.Raw(sql, args...)
	if err != nil {
		return QueryResult{Error: err}
	}

	destValue := reflect.ValueOf(dest)
	if destValue.Kind() != reflect.Ptr || destValue.Elem().Kind() != reflect.Slice {
		return QueryResult{Error: fmt.Errorf("dest must be a pointer to a slice")}
	}

	destElem := destValue.Elem()
	for _, row := range rows {
		if val, ok := row[column]; ok && val != nil {
			destElem.Set(reflect.Append(destElem, reflect.ValueOf(val)))
		}
	}

	return QueryResult{RowsAffected: int64(len(rows))}
}

// Scan scans results into dest
func (q *QueryBuilderImpl) Scan(dest interface{}) QueryResult {
	return q.Find(dest)
}

// Create inserts a new record
func (q *QueryBuilderImpl) Create(value interface{}) QueryResult {
	columns, values, args := q.extractInsertData(value)
	if len(columns) == 0 {
		return QueryResult{Error: fmt.Errorf("no columns to insert")}
	}

	placeholders := make([]string, len(values))
	for i := range values {
		placeholders[i] = "?"
	}

	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		q.tableName,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "))

	if q.debug && q.logger != nil {
		q.logger.Info("SQL", "query", sql, "args", args)
	}

	rowsAffected, lastInsertID, err := q.db.ExecWithResult(sql, args...)
	if err == nil && lastInsertID > 0 {
		setStructID(value, lastInsertID)
	}
	return QueryResult{Error: err, RowsAffected: rowsAffected, LastInsertID: lastInsertID}
}

// Save updates or creates a record.
// Detects the primary key from the gorm:"primaryKey" tag, falling back to ID/Id fields.
func (q *QueryBuilderImpl) Save(value interface{}) QueryResult {
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return q.Create(value)
	}

	// Find primary key field and its column name
	pkCol, pkField := q.findPrimaryKeyField(v)
	if pkField.IsValid() && !pkField.IsZero() {
		return q.Where(pkCol+" = ?", pkField.Interface()).Updates(value)
	}
	return q.Create(value)
}

// findPrimaryKeyField returns the column name and field value of the primary key.
// Checks gorm:"primaryKey" tag first, then falls back to ID/Id fields.
func (q *QueryBuilderImpl) findPrimaryKeyField(v reflect.Value) (string, reflect.Value) {
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		gormTag := field.Tag.Get("gorm")
		if strings.Contains(gormTag, "primaryKey") {
			colName := q.getColumnName(field)
			return colName, v.Field(i)
		}
	}
	// Fallback to ID/Id
	if f := v.FieldByName("ID"); f.IsValid() {
		return "id", f
	}
	if f := v.FieldByName("Id"); f.IsValid() {
		return "id", f
	}
	return "id", reflect.Value{}
}

// detectPrimaryKey returns the column name of the primary key from the dest struct.
func (q *QueryBuilderImpl) detectPrimaryKey(dest interface{}) string {
	v := reflect.ValueOf(dest)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	// For slices, inspect the element type
	if v.Kind() == reflect.Slice {
		elemType := v.Type().Elem()
		if elemType.Kind() == reflect.Ptr {
			elemType = elemType.Elem()
		}
		if elemType.Kind() == reflect.Struct {
			v = reflect.New(elemType).Elem()
		}
	}
	if v.Kind() == reflect.Struct {
		col, _ := q.findPrimaryKeyField(v)
		return col
	}
	return "id"
}

// Update updates a single column
func (q *QueryBuilderImpl) Update(column string, value interface{}) QueryResult {
	whereParts, whereArgs := q.buildWhereClauses()
	if len(whereParts) == 0 {
		return QueryResult{Error: fmt.Errorf("WHERE clause required for UPDATE")}
	}

	sql := fmt.Sprintf("UPDATE %s SET %s = ? WHERE %s",
		q.tableName, column, strings.Join(whereParts, " AND "))

	args := append([]interface{}{value}, whereArgs...)

	if q.debug && q.logger != nil {
		q.logger.Info("SQL", "query", sql, "args", args)
	}

	err := q.db.Exec(sql, args...)
	return QueryResult{Error: err}
}

// Updates updates multiple columns
func (q *QueryBuilderImpl) Updates(values interface{}) QueryResult {
	columns, _, args := q.extractUpdateData(values)
	if len(columns) == 0 {
		return QueryResult{Error: fmt.Errorf("no columns to update")}
	}

	whereParts, whereArgs := q.buildWhereClauses()
	if len(whereParts) == 0 {
		return QueryResult{Error: fmt.Errorf("WHERE clause required for UPDATE")}
	}

	setParts := make([]string, len(columns))
	for i, col := range columns {
		setParts[i] = col + " = ?"
	}

	sql := fmt.Sprintf("UPDATE %s SET %s WHERE %s",
		q.tableName,
		strings.Join(setParts, ", "),
		strings.Join(whereParts, " AND "))

	args = append(args, whereArgs...)

	if q.debug && q.logger != nil {
		q.logger.Info("SQL", "query", sql, "args", args)
	}

	err := q.db.Exec(sql, args...)
	return QueryResult{Error: err}
}

// Delete removes records
func (q *QueryBuilderImpl) Delete(value interface{}, conds ...interface{}) QueryResult {
	// Add conditions from value if it has an ID
	if value != nil {
		v := reflect.ValueOf(value)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		if v.Kind() == reflect.Struct {
			idField := v.FieldByName("ID")
			if !idField.IsValid() {
				idField = v.FieldByName("Id")
			}
			if idField.IsValid() && idField.Uint() > 0 {
				q.Where("id = ?", idField.Uint())
			}
		}
	}

	// Add additional conditions
	if len(conds) > 0 {
		if query, ok := conds[0].(string); ok {
			q.Where(query, conds[1:]...)
		}
	}

	whereParts, whereArgs := q.buildWhereClauses()
	if len(whereParts) == 0 {
		return QueryResult{Error: fmt.Errorf("WHERE clause required for DELETE")}
	}

	sql := fmt.Sprintf("DELETE FROM %s WHERE %s",
		q.tableName, strings.Join(whereParts, " AND "))

	if q.debug && q.logger != nil {
		q.logger.Info("SQL", "query", sql, "args", whereArgs)
	}

	err := q.db.Exec(sql, whereArgs...)
	return QueryResult{Error: err}
}

// FirstOrCreate finds first record matching conditions, or creates a new one
func (q *QueryBuilderImpl) FirstOrCreate(dest interface{}, conds ...interface{}) QueryResult {
	// Add conditions if provided
	if len(conds) > 0 {
		if query, ok := conds[0].(string); ok {
			q.Where(query, conds[1:]...)
		}
	}

	// Try to find first
	result := q.First(dest)
	if result.RowsAffected > 0 {
		return result
	}

	// Not found, create new
	// Reset query builder for insert
	createBuilder := &QueryBuilderImpl{
		db:        q.db,
		logger:    q.logger,
		tableName: q.tableName,
		debug:     q.debug,
	}
	return createBuilder.Create(dest)
}

// Upsert inserts or updates on duplicate key (MySQL specific)
// updateColumns specifies which columns to update on conflict; empty = update all non-key columns
func (q *QueryBuilderImpl) Upsert(value interface{}, updateColumns ...string) QueryResult {
	columns, _, args := q.extractInsertData(value)
	if len(columns) == 0 {
		return QueryResult{Error: fmt.Errorf("no columns to insert")}
	}

	placeholders := make([]string, len(columns))
	for i := range columns {
		placeholders[i] = "?"
	}

	// Build ON DUPLICATE KEY UPDATE clause
	var updateParts []string
	if len(updateColumns) > 0 {
		// Update only specified columns
		for _, col := range updateColumns {
			updateParts = append(updateParts, fmt.Sprintf("%s = VALUES(%s)", col, col))
		}
	} else {
		// Update all columns except id
		for _, col := range columns {
			if col != "id" {
				updateParts = append(updateParts, fmt.Sprintf("%s = VALUES(%s)", col, col))
			}
		}
	}

	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) ON DUPLICATE KEY UPDATE %s",
		q.tableName,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
		strings.Join(updateParts, ", "))

	if q.debug && q.logger != nil {
		q.logger.Info("SQL", "query", sql, "args", args)
	}

	rowsAffected, lastInsertID, err := q.db.ExecWithResult(sql, args...)
	return QueryResult{Error: err, RowsAffected: rowsAffected, LastInsertID: lastInsertID}
}

// Helper: scan rows into destination
func (q *QueryBuilderImpl) scanRows(rows []map[string]interface{}, dest interface{}) error {
	destValue := reflect.ValueOf(dest)
	if destValue.Kind() != reflect.Ptr {
		return fmt.Errorf("dest must be a pointer")
	}

	destElem := destValue.Elem()

	// Handle slice destination
	if destElem.Kind() == reflect.Slice {
		destType := destElem.Type().Elem()
		for _, row := range rows {
			item := reflect.New(destType).Elem()
			if err := q.mapToStruct(row, item); err != nil {
				return err
			}
			destElem.Set(reflect.Append(destElem, item))
		}
		return nil
	}

	// Handle single struct destination
	if destElem.Kind() == reflect.Struct && len(rows) > 0 {
		return q.mapToStruct(rows[0], destElem)
	}

	return nil
}

func (q *QueryBuilderImpl) mapToStruct(data map[string]interface{}, structVal reflect.Value) error {
	structType := structVal.Type()

	for i := 0; i < structType.NumField(); i++ {
		field := structVal.Field(i)
		fieldType := structType.Field(i)

		// Get column name from json or gorm tag
		key := fieldType.Tag.Get("json")
		if key == "" || key == "-" {
			key = fieldType.Tag.Get("gorm")
			if idx := strings.Index(key, "column:"); idx >= 0 {
				key = key[idx+7:]
				if end := strings.Index(key, ";"); end >= 0 {
					key = key[:end]
				}
			} else {
				key = fieldType.Name
			}
		}
		if idx := strings.Index(key, ","); idx >= 0 {
			key = key[:idx]
		}

		value, ok := data[key]
		if !ok || value == nil {
			continue
		}

		if !field.CanSet() {
			continue
		}

		fieldValue := reflect.ValueOf(value)
		if fieldValue.Type().ConvertibleTo(field.Type()) {
			field.Set(fieldValue.Convert(field.Type()))
		} else if field.Type() == reflect.TypeOf(time.Time{}) {
			if str, ok := value.(string); ok {
				if t, err := time.Parse("2006-01-02 15:04:05", str); err == nil {
					field.Set(reflect.ValueOf(t))
				}
			}
		}
	}

	return nil
}

func (q *QueryBuilderImpl) extractInsertData(value interface{}) ([]string, []string, []interface{}) {
	var columns []string
	var placeholders []string
	var args []interface{}

	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() == reflect.Map {
		for _, key := range v.MapKeys() {
			columns = append(columns, key.String())
			placeholders = append(placeholders, "?")
			args = append(args, v.MapIndex(key).Interface())
		}
		return columns, placeholders, args
	}

	if v.Kind() != reflect.Struct {
		return columns, placeholders, args
	}

	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldVal := v.Field(i)

		// Skip unexported fields
		if !fieldVal.CanInterface() {
			continue
		}

		// Get column name
		colName := q.getColumnName(field)
		if colName == "-" || colName == "" {
			continue
		}

		// Skip zero values for auto-increment fields
		if colName == "id" && fieldVal.IsZero() {
			continue
		}

		columns = append(columns, colName)
		placeholders = append(placeholders, "?")
		args = append(args, fieldVal.Interface())
	}

	return columns, placeholders, args
}

func (q *QueryBuilderImpl) extractUpdateData(value interface{}) ([]string, []string, []interface{}) {
	var columns []string
	var placeholders []string
	var args []interface{}

	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() == reflect.Map {
		for _, key := range v.MapKeys() {
			columns = append(columns, key.String())
			args = append(args, v.MapIndex(key).Interface())
		}
		return columns, placeholders, args
	}

	if v.Kind() != reflect.Struct {
		return columns, placeholders, args
	}

	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldVal := v.Field(i)

		if !fieldVal.CanInterface() {
			continue
		}

		colName := q.getColumnName(field)
		if colName == "-" || colName == "" || colName == "id" {
			continue
		}

		// Skip zero values unless explicitly set
		if fieldVal.IsZero() {
			continue
		}

		columns = append(columns, colName)
		args = append(args, fieldVal.Interface())
	}

	return columns, placeholders, args
}

func (q *QueryBuilderImpl) getColumnName(field reflect.StructField) string {
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

	// Default to snake_case of field name
	return toSnakeCase(field.Name)
}

func toSnakeCase(s string) string {
	runes := []rune(s)
	var result strings.Builder
	for i, r := range runes {
		if i > 0 && r >= 'A' && r <= 'Z' {
			// Don't insert underscore if previous char was also uppercase and next is lowercase
			// e.g. "ID" -> "id", "XMLParser" -> "xml_parser", "UserID" -> "user_id"
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
