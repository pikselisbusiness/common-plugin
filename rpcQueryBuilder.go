package shared

import (
	"encoding/json"
	"fmt"
	"net/rpc"

	"github.com/pikselisbusiness/go-plugin"
)

// RPC Request/Response types for QueryBuilder operations

type QueryRequest struct {
	TableName    string
	SelectCols   []string
	DistinctCols []string
	WhereConds   []WhereCondRPC
	OrConds      []WhereCondRPC
	NotConds     []WhereCondRPC
	OrderClauses []OrderClauseRPC
	GroupCols    []string
	HavingCond   *WhereCondRPC
	Joins        []JoinClauseRPC
	LimitVal     *int
	OffsetVal    *int
	RawSQL       string
	RawArgs      []interface{}
}

type WhereCondRPC struct {
	Query string
	Args  []interface{}
}

type OrderClauseRPC struct {
	Column    string
	Direction string
}

type JoinClauseRPC struct {
	JoinType  string
	Table     string
	Condition string
	Args      []interface{}
}

type QueryResponse struct {
	Rows         []map[string]interface{}
	RowsAffected int64
	LastInsertID int64
	Error        string
}

type ExecRequest struct {
	TableName     string
	Operation     string // "create", "update", "delete", "upsert"
	Data          map[string]interface{}
	WhereConds    []WhereCondRPC
	UpdateColumns []string // for upsert: columns to update on conflict
}

// QueryBuilderRPCClient implements QueryBuilder over RPC (plugin side)
type QueryBuilderRPCClient struct {
	client    *rpc.Client
	broker    *plugin.MuxBroker
	tableName string
	request   QueryRequest
}

func NewQueryBuilderRPCClient(client *rpc.Client, broker *plugin.MuxBroker, tableName string) *QueryBuilderRPCClient {
	return &QueryBuilderRPCClient{
		client:    client,
		broker:    broker,
		tableName: tableName,
		request:   QueryRequest{TableName: tableName},
	}
}

func (q *QueryBuilderRPCClient) Select(columns ...string) QueryBuilder {
	q.request.SelectCols = append(q.request.SelectCols, columns...)
	return q
}

func (q *QueryBuilderRPCClient) Distinct(columns ...string) QueryBuilder {
	q.request.DistinctCols = append(q.request.DistinctCols, columns...)
	return q
}

func (q *QueryBuilderRPCClient) Omit(columns ...string) QueryBuilder {
	// Handled locally, not sent to server
	return q
}

func (q *QueryBuilderRPCClient) Where(query string, args ...interface{}) QueryBuilder {
	q.request.WhereConds = append(q.request.WhereConds, WhereCondRPC{Query: query, Args: args})
	return q
}

func (q *QueryBuilderRPCClient) Or(query string, args ...interface{}) QueryBuilder {
	q.request.OrConds = append(q.request.OrConds, WhereCondRPC{Query: query, Args: args})
	return q
}

func (q *QueryBuilderRPCClient) Not(query string, args ...interface{}) QueryBuilder {
	q.request.NotConds = append(q.request.NotConds, WhereCondRPC{Query: query, Args: args})
	return q
}

func (q *QueryBuilderRPCClient) Order(column string, direction ...OrderDirection) QueryBuilder {
	dir := "ASC"
	if len(direction) > 0 {
		dir = string(direction[0])
	}
	q.request.OrderClauses = append(q.request.OrderClauses, OrderClauseRPC{Column: column, Direction: dir})
	return q
}

func (q *QueryBuilderRPCClient) Limit(limit int) QueryBuilder {
	q.request.LimitVal = &limit
	return q
}

func (q *QueryBuilderRPCClient) Offset(offset int) QueryBuilder {
	q.request.OffsetVal = &offset
	return q
}

func (q *QueryBuilderRPCClient) Group(columns ...string) QueryBuilder {
	q.request.GroupCols = append(q.request.GroupCols, columns...)
	return q
}

func (q *QueryBuilderRPCClient) Having(query string, args ...interface{}) QueryBuilder {
	q.request.HavingCond = &WhereCondRPC{Query: query, Args: args}
	return q
}

func (q *QueryBuilderRPCClient) Join(table string, condition string, args ...interface{}) QueryBuilder {
	q.request.Joins = append(q.request.Joins, JoinClauseRPC{JoinType: "INNER", Table: table, Condition: condition, Args: args})
	return q
}

func (q *QueryBuilderRPCClient) LeftJoin(table string, condition string, args ...interface{}) QueryBuilder {
	q.request.Joins = append(q.request.Joins, JoinClauseRPC{JoinType: "LEFT", Table: table, Condition: condition, Args: args})
	return q
}

func (q *QueryBuilderRPCClient) RightJoin(table string, condition string, args ...interface{}) QueryBuilder {
	q.request.Joins = append(q.request.Joins, JoinClauseRPC{JoinType: "RIGHT", Table: table, Condition: condition, Args: args})
	return q
}

func (q *QueryBuilderRPCClient) Preload(association string, args ...interface{}) QueryBuilder {
	// Preload requires special handling - store for later
	return q
}

func (q *QueryBuilderRPCClient) Raw(sql string, values ...interface{}) QueryBuilder {
	q.request.RawSQL = sql
	q.request.RawArgs = values
	return q
}

func (q *QueryBuilderRPCClient) Scopes(funcs ...func(QueryBuilder) QueryBuilder) QueryBuilder {
	var result QueryBuilder = q
	for _, f := range funcs {
		result = f(result)
	}
	return result
}

func (q *QueryBuilderRPCClient) Debug() QueryBuilder {
	return q
}

func (q *QueryBuilderRPCClient) ToSQL() (string, []interface{}) {
	return q.request.RawSQL, q.request.RawArgs
}

// Execution methods - these make RPC calls

func (q *QueryBuilderRPCClient) Find(dest interface{}) QueryResult {
	var resp QueryResponse
	err := q.client.Call("Plugin.QueryFind", q.request, &resp)
	if err != nil {
		return QueryResult{Error: err}
	}
	if resp.Error != "" {
		return QueryResult{Error: fmt.Errorf(resp.Error)}
	}

	// Scan rows into dest using reflection
	if err := scanRowsInto(resp.Rows, dest); err != nil {
		return QueryResult{Error: err}
	}

	return QueryResult{RowsAffected: resp.RowsAffected}
}

func (q *QueryBuilderRPCClient) First(dest interface{}) QueryResult {
	q.Limit(1).Order("id", ASC)
	return q.Find(dest)
}

func (q *QueryBuilderRPCClient) Last(dest interface{}) QueryResult {
	q.Limit(1).Order("id", DESC)
	return q.Find(dest)
}

func (q *QueryBuilderRPCClient) Take(dest interface{}) QueryResult {
	q.Limit(1)
	return q.Find(dest)
}

func (q *QueryBuilderRPCClient) Count(count *int64) QueryResult {
	q.request.SelectCols = []string{"COUNT(*) as count"}
	var resp QueryResponse
	err := q.client.Call("Plugin.QueryFind", q.request, &resp)
	if err != nil {
		return QueryResult{Error: err}
	}
	if len(resp.Rows) > 0 {
		if c, ok := resp.Rows[0]["count"]; ok {
			switch v := c.(type) {
			case float64:
				*count = int64(v)
			case int64:
				*count = v
			case json.Number:
				if n, err := v.Int64(); err == nil {
					*count = n
				}
			}
		}
	}
	return QueryResult{}
}

func (q *QueryBuilderRPCClient) Pluck(column string, dest interface{}) QueryResult {
	q.request.SelectCols = []string{column}
	return q.Find(dest)
}

func (q *QueryBuilderRPCClient) Scan(dest interface{}) QueryResult {
	return q.Find(dest)
}

func (q *QueryBuilderRPCClient) Create(value interface{}) QueryResult {
	data := structToMap(value)
	req := ExecRequest{
		TableName: q.tableName,
		Operation: "create",
		Data:      data,
	}
	var resp QueryResponse
	err := q.client.Call("Plugin.QueryExec", req, &resp)
	if err != nil {
		return QueryResult{Error: err}
	}
	if resp.Error != "" {
		return QueryResult{Error: fmt.Errorf(resp.Error)}
	}

	// Update the ID field on the original struct (like GORM does)
	if resp.LastInsertID > 0 {
		setStructID(value, resp.LastInsertID)
	}

	return QueryResult{RowsAffected: resp.RowsAffected, LastInsertID: resp.LastInsertID}
}

func (q *QueryBuilderRPCClient) Save(value interface{}) QueryResult {
	data := structToMap(value)
	if id, ok := data["id"]; ok && id != nil && id != 0 {
		return q.Where("id = ?", id).Updates(value)
	}
	return q.Create(value)
}

func (q *QueryBuilderRPCClient) Update(column string, value interface{}) QueryResult {
	return q.Updates(map[string]interface{}{column: value})
}

func (q *QueryBuilderRPCClient) Updates(values interface{}) QueryResult {
	data := structToMap(values)
	req := ExecRequest{
		TableName:  q.tableName,
		Operation:  "update",
		Data:       data,
		WhereConds: q.request.WhereConds,
	}
	var resp QueryResponse
	err := q.client.Call("Plugin.QueryExec", req, &resp)
	if err != nil {
		return QueryResult{Error: err}
	}
	if resp.Error != "" {
		return QueryResult{Error: fmt.Errorf(resp.Error)}
	}
	return QueryResult{RowsAffected: resp.RowsAffected}
}

func (q *QueryBuilderRPCClient) Delete(value interface{}, conds ...interface{}) QueryResult {
	if value != nil {
		data := structToMap(value)
		if id, ok := data["id"]; ok && id != nil && id != 0 {
			q.Where("id = ?", id)
		}
	}
	if len(conds) > 0 {
		if query, ok := conds[0].(string); ok {
			q.Where(query, conds[1:]...)
		}
	}

	req := ExecRequest{
		TableName:  q.tableName,
		Operation:  "delete",
		WhereConds: q.request.WhereConds,
	}
	var resp QueryResponse
	err := q.client.Call("Plugin.QueryExec", req, &resp)
	if err != nil {
		return QueryResult{Error: err}
	}
	if resp.Error != "" {
		return QueryResult{Error: fmt.Errorf(resp.Error)}
	}
	return QueryResult{RowsAffected: resp.RowsAffected}
}

// FirstOrCreate finds first record matching conditions, or creates a new one
func (q *QueryBuilderRPCClient) FirstOrCreate(dest interface{}, conds ...interface{}) QueryResult {
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
	createBuilder := NewQueryBuilderRPCClient(q.client, q.broker, q.tableName)
	return createBuilder.Create(dest)
}

// Upsert inserts or updates on duplicate key (MySQL specific)
func (q *QueryBuilderRPCClient) Upsert(value interface{}, updateColumns ...string) QueryResult {
	data := structToMap(value)
	req := ExecRequest{
		TableName:     q.tableName,
		Operation:     "upsert",
		Data:          data,
		UpdateColumns: updateColumns,
	}
	var resp QueryResponse
	err := q.client.Call("Plugin.QueryExec", req, &resp)
	if err != nil {
		return QueryResult{Error: err}
	}
	if resp.Error != "" {
		return QueryResult{Error: fmt.Errorf(resp.Error)}
	}
	return QueryResult{RowsAffected: resp.RowsAffected, LastInsertID: resp.LastInsertID}
}

// QueryBuilderRPCServer handles RPC calls on the host side
type QueryBuilderRPCServer struct {
	db     DB
	broker *plugin.MuxBroker
}

func NewQueryBuilderRPCServer(db DB, broker *plugin.MuxBroker) *QueryBuilderRPCServer {
	return &QueryBuilderRPCServer{db: db, broker: broker}
}

func (s *QueryBuilderRPCServer) QueryFind(req QueryRequest, resp *QueryResponse) error {
	sql, args := s.buildSelectSQL(req)
	rows, err := s.db.Raw(sql, args...)
	if err != nil {
		resp.Error = err.Error()
		return nil
	}
	resp.Rows = rows
	resp.RowsAffected = int64(len(rows))
	return nil
}

func (s *QueryBuilderRPCServer) QueryExec(req ExecRequest, resp *QueryResponse) error {
	var sql string
	var args []interface{}

	switch req.Operation {
	case "create":
		sql, args = s.buildInsertSQL(req)
	case "update":
		sql, args = s.buildUpdateSQL(req)
	case "delete":
		sql, args = s.buildDeleteSQL(req)
	case "upsert":
		sql, args = s.buildUpsertSQL(req)
	default:
		resp.Error = "unknown operation: " + req.Operation
		return nil
	}

	rowsAffected, lastInsertID, err := s.db.ExecWithResult(sql, args...)
	if err != nil {
		resp.Error = err.Error()
	}
	resp.RowsAffected = rowsAffected
	resp.LastInsertID = lastInsertID
	return nil
}

func (s *QueryBuilderRPCServer) buildSelectSQL(req QueryRequest) (string, []interface{}) {
	if req.RawSQL != "" {
		return req.RawSQL, req.RawArgs
	}

	var sql string
	var args []interface{}

	// SELECT
	sql = "SELECT "
	if len(req.DistinctCols) > 0 {
		sql += "DISTINCT " + join(req.DistinctCols, ", ")
	} else if len(req.SelectCols) > 0 {
		sql += join(req.SelectCols, ", ")
	} else {
		sql += "*"
	}

	// FROM
	sql += " FROM " + req.TableName

	// JOINS
	for _, j := range req.Joins {
		sql += fmt.Sprintf(" %s JOIN %s ON %s", j.JoinType, j.Table, j.Condition)
		args = append(args, j.Args...)
	}

	// WHERE
	whereParts, whereArgs := s.buildWhereClauses(req.WhereConds, req.OrConds, req.NotConds)
	if len(whereParts) > 0 {
		sql += " WHERE " + join(whereParts, " AND ")
		args = append(args, whereArgs...)
	}

	// GROUP BY
	if len(req.GroupCols) > 0 {
		sql += " GROUP BY " + join(req.GroupCols, ", ")
	}

	// HAVING
	if req.HavingCond != nil {
		sql += " HAVING " + req.HavingCond.Query
		args = append(args, req.HavingCond.Args...)
	}

	// ORDER BY
	if len(req.OrderClauses) > 0 {
		orderParts := make([]string, len(req.OrderClauses))
		for i, o := range req.OrderClauses {
			orderParts[i] = o.Column + " " + o.Direction
		}
		sql += " ORDER BY " + join(orderParts, ", ")
	}

	// LIMIT
	if req.LimitVal != nil {
		sql += fmt.Sprintf(" LIMIT %d", *req.LimitVal)
	}

	// OFFSET
	if req.OffsetVal != nil {
		sql += fmt.Sprintf(" OFFSET %d", *req.OffsetVal)
	}

	return sql, args
}

func (s *QueryBuilderRPCServer) buildWhereClauses(where, or, not []WhereCondRPC) ([]string, []interface{}) {
	var parts []string
	var args []interface{}

	for _, w := range where {
		parts = append(parts, "("+w.Query+")")
		args = append(args, w.Args...)
	}

	for _, w := range not {
		parts = append(parts, "NOT ("+w.Query+")")
		args = append(args, w.Args...)
	}

	if len(or) > 0 {
		orParts := make([]string, len(or))
		for i, w := range or {
			orParts[i] = "(" + w.Query + ")"
			args = append(args, w.Args...)
		}
		parts = append(parts, "("+join(orParts, " OR ")+")")
	}

	return parts, args
}

func (s *QueryBuilderRPCServer) buildInsertSQL(req ExecRequest) (string, []interface{}) {
	var columns []string
	var placeholders []string
	var args []interface{}

	for col, val := range req.Data {
		if col == "id" && (val == nil || val == 0 || val == float64(0)) {
			continue
		}
		columns = append(columns, col)
		placeholders = append(placeholders, "?")
		args = append(args, val)
	}

	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		req.TableName,
		join(columns, ", "),
		join(placeholders, ", "))

	return sql, args
}

func (s *QueryBuilderRPCServer) buildUpdateSQL(req ExecRequest) (string, []interface{}) {
	var setParts []string
	var args []interface{}

	for col, val := range req.Data {
		if col == "id" {
			continue
		}
		setParts = append(setParts, col+" = ?")
		args = append(args, val)
	}

	whereParts, whereArgs := s.buildWhereClauses(req.WhereConds, nil, nil)
	args = append(args, whereArgs...)

	sql := fmt.Sprintf("UPDATE %s SET %s WHERE %s",
		req.TableName,
		join(setParts, ", "),
		join(whereParts, " AND "))

	return sql, args
}

func (s *QueryBuilderRPCServer) buildDeleteSQL(req ExecRequest) (string, []interface{}) {
	whereParts, whereArgs := s.buildWhereClauses(req.WhereConds, nil, nil)

	sql := fmt.Sprintf("DELETE FROM %s WHERE %s",
		req.TableName,
		join(whereParts, " AND "))

	return sql, whereArgs
}

func (s *QueryBuilderRPCServer) buildUpsertSQL(req ExecRequest) (string, []interface{}) {
	var columns []string
	var placeholders []string
	var args []interface{}

	for col, val := range req.Data {
		if col == "id" && (val == nil || val == 0 || val == float64(0)) {
			continue
		}
		columns = append(columns, col)
		placeholders = append(placeholders, "?")
		args = append(args, val)
	}

	// Build ON DUPLICATE KEY UPDATE clause
	var updateParts []string
	if len(req.UpdateColumns) > 0 {
		// Update only specified columns
		for _, col := range req.UpdateColumns {
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
		req.TableName,
		join(columns, ", "),
		join(placeholders, ", "),
		join(updateParts, ", "))

	return sql, args
}

// Helper functions

func join(parts []string, sep string) string {
	if len(parts) == 0 {
		return ""
	}
	result := parts[0]
	for i := 1; i < len(parts); i++ {
		result += sep + parts[i]
	}
	return result
}
