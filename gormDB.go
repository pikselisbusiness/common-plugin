package shared

import (
	"net"
	"net/rpc"

	"gorm.io/gorm"
)

// GormDB is the host side of DB over gorm, the way business-go serves it to
// every plugin. Plugin tests pair it with DBv2OverRPC and a test database.
type GormDB struct {
	DB *gorm.DB
}

func (d GormDB) MigrateModel(tableName string, fields []MigrateField) error {
	return d.DB.Table(tableName).AutoMigrate(CreateStructFromFields(fields, tableName))
}

func (d GormDB) Exec(sql string, values ...interface{}) error {
	return d.DB.Exec(sql, values...).Error
}

// ExecWithResult reads LAST_INSERT_ID() on the connection that ran the
// statement; another connection of the pool would answer for its own insert.
func (d GormDB) ExecWithResult(sql string, values ...interface{}) (int64, int64, error) {
	var rowsAffected, lastInsertID int64
	err := d.DB.Connection(func(conn *gorm.DB) error {
		result := conn.Exec(sql, values...)
		if result.Error != nil {
			return result.Error
		}
		rowsAffected = result.RowsAffected
		if rowsAffected == 0 {
			return nil
		}
		var idResult struct{ ID int64 }
		if err := conn.Raw("SELECT LAST_INSERT_ID() as id").Scan(&idResult).Error; err != nil {
			return err
		}
		lastInsertID = idResult.ID
		return nil
	})
	if err != nil {
		return 0, 0, err
	}
	return rowsAffected, lastInsertID, nil
}

func (d GormDB) Raw(sql string, values ...interface{}) ([]map[string]interface{}, error) {
	var items []map[string]interface{}
	return items, d.DB.Raw(sql, values...).Scan(&items).Error
}

// DBv2OverRPC connects a plugin-side DBv2 to host through net/rpc with gob,
// the transport go-plugin runs between plugin and host, so a plugin test
// meets the same encoding as production. close ends the connection.
func DBv2OverRPC(host DB) (db DBv2, close func() error, err error) {
	server := rpc.NewServer()
	if err := server.RegisterName("Plugin", &dbRPCServer{impl: host}); err != nil {
		return nil, nil, err
	}
	hostConn, pluginConn := net.Pipe()
	go server.ServeConn(hostConn)
	client := rpc.NewClient(pluginConn)
	return NewDBv2RPCClient(client, nil, &dbRPCClient{client: client}), client.Close, nil
}
