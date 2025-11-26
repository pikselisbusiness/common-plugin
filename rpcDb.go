package shared

import (
	"fmt"
	"net/rpc"

	"github.com/pikselisbusiness/go-plugin"
)

type dbRPCClient struct {
	client *rpc.Client
	broker *plugin.MuxBroker
}

type dbRPCServer struct {
	impl   DB
	broker *plugin.MuxBroker
}

// api
type DbMigrateRequest struct {
	TableName string
	Fields    []MigrateField
}

type DbMigrateResponse struct {
	Error error
}
type DbStatementRequest struct {
	Sql    string
	Values []interface{}
}

type DbStatementResponse struct {
	Error error
}

type DbRawResponse struct {
	Items []map[string]interface{}
	Error error
}

func (m *dbRPCServer) MigrateModel(req DbMigrateRequest, resp *DbMigrateResponse) error {
	err := m.impl.MigrateModel(req.TableName, req.Fields)

	if err != nil {
		fmt.Println("Error calling Plugin.Migrate", "error", err)
	}
	resp.Error = err

	return nil
}
func (m *dbRPCClient) MigrateModel(tableName string, fields []MigrateField) error {

	var reply DbMigrateResponse
	err := m.client.Call("Plugin.MigrateModel", DbMigrateRequest{
		TableName: tableName,
		Fields:    fields,
	}, &reply)
	if err != nil {
		fmt.Println("Error calling Plugin.MigrateModel", "error", err)
		return err
	}

	return reply.Error
}

func (m *dbRPCServer) Exec(req DbStatementRequest, resp *DbStatementResponse) error {
	err := m.impl.Exec(req.Sql, req.Values...)

	fmt.Println("CALL query", req.Sql, req.Values)
	if err != nil {
		fmt.Println("Error calling Plugin.Exec", "error", err)
	}
	resp.Error = encodableError(err)

	return nil
}
func (m *dbRPCClient) Exec(sql string, values ...interface{}) error {

	var reply DbStatementResponse
	err := m.client.Call("Plugin.Exec", DbStatementRequest{
		Sql:    sql,
		Values: values,
	}, &reply)

	fmt.Println("CALL query", sql, values)

	if err != nil {
		fmt.Println("Error calling Plugin.Exec", "error", err)
		return err
	}

	return reply.Error
}

func (m *dbRPCServer) Raw(req DbStatementRequest, resp *DbRawResponse) error {
	items, err := m.impl.Raw(req.Sql, req.Values...)

	if err != nil {
		fmt.Println("Error calling Plugin.Raw", "error", err)
	}
	resp.Items = items
	resp.Error = encodableError(err)

	return nil
}
func (m *dbRPCClient) Raw(sql string, values ...interface{}) ([]map[string]interface{}, error) {

	var reply DbRawResponse
	err := m.client.Call("Plugin.Raw", DbStatementRequest{
		Sql:    sql,
		Values: values,
	}, &reply)
	if err != nil {
		fmt.Println("Error calling Plugin.Raw", "error", err)
		return []map[string]interface{}{}, err
	}

	return reply.Items, reply.Error
}
