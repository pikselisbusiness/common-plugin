package shared

import (
	"net/rpc"
	"reflect"

	"github.com/hashicorp/go-hclog"
	"github.com/pikselisbusiness/go-plugin"
)

// DBv2Impl implements DBv2 interface on the host side
type DBv2Impl struct {
	db     DB
	logger hclog.Logger
}

func NewDBv2(db DB, logger hclog.Logger) *DBv2Impl {
	return &DBv2Impl{db: db, logger: logger}
}

func (d *DBv2Impl) MigrateModel(tableName string, fields []MigrateField) error {
	return d.db.MigrateModel(tableName, fields)
}

func (d *DBv2Impl) Exec(sql string, values ...interface{}) error {
	return d.db.Exec(sql, values...)
}

func (d *DBv2Impl) Raw(sql string, values ...interface{}) ([]map[string]interface{}, error) {
	return d.db.Raw(sql, values...)
}

func (d *DBv2Impl) Begin() (DBv2, error) {
	// For now, transactions are not supported over RPC
	// Return self - operations will be non-transactional
	return d, nil
}

func (d *DBv2Impl) Commit() error {
	return nil
}

func (d *DBv2Impl) Rollback() error {
	return nil
}

func (d *DBv2Impl) Table(name string) QueryBuilder {
	return NewQueryBuilder(d.db, d.logger, name)
}

func (d *DBv2Impl) Model(value interface{}) QueryBuilder {
	tableName := getTableName(value)
	return NewQueryBuilder(d.db, d.logger, tableName)
}

// getTableName extracts table name from model
func getTableName(value interface{}) string {
	// Check if model implements TableName() method
	if tabler, ok := value.(interface{ TableName() string }); ok {
		return tabler.TableName()
	}

	// Get type name and convert to snake_case
	t := reflect.TypeOf(value)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return toSnakeCaseHelper(t.Name()) + "s"
}

// DBv2RPCClient implements DBv2 on the plugin side (RPC client)
type DBv2RPCClient struct {
	client *rpc.Client
	broker *plugin.MuxBroker
	db     DB // fallback to basic DB interface
}

func NewDBv2RPCClient(client *rpc.Client, broker *plugin.MuxBroker, db DB) *DBv2RPCClient {
	return &DBv2RPCClient{client: client, broker: broker, db: db}
}

func (d *DBv2RPCClient) MigrateModel(tableName string, fields []MigrateField) error {
	return d.db.MigrateModel(tableName, fields)
}

func (d *DBv2RPCClient) Exec(sql string, values ...interface{}) error {
	return d.db.Exec(sql, values...)
}

func (d *DBv2RPCClient) Raw(sql string, values ...interface{}) ([]map[string]interface{}, error) {
	return d.db.Raw(sql, values...)
}

func (d *DBv2RPCClient) Begin() (DBv2, error) {
	return d, nil
}

func (d *DBv2RPCClient) Commit() error {
	return nil
}

func (d *DBv2RPCClient) Rollback() error {
	return nil
}

func (d *DBv2RPCClient) Table(name string) QueryBuilder {
	return NewQueryBuilderRPCClient(d.client, d.broker, name)
}

func (d *DBv2RPCClient) Model(value interface{}) QueryBuilder {
	tableName := getTableName(value)
	return NewQueryBuilderRPCClient(d.client, d.broker, tableName)
}
