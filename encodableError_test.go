package shared

import (
	"bytes"
	"encoding/gob"
	"testing"

	"github.com/go-sql-driver/mysql"
)

// A raw *mysql.MySQLError in an RPC response is not gob-registered: encoding it
// breaks the net/rpc connection and kills the plugin. encodableError must flatten it.
func TestEncodableErrorMySQLErrorSurvivesGob(t *testing.T) {
	sqlErr := &mysql.MySQLError{Number: 1071, Message: "Specified key was too long; max key length is 3072 bytes"}

	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(DbMigrateResponse{Error: encodableError(sqlErr)}); err != nil {
		t.Fatalf("encode: %v", err)
	}
	var decoded DbMigrateResponse
	if err := gob.NewDecoder(&buf).Decode(&decoded); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if decoded.Error == nil || decoded.Error.Error() != sqlErr.Error() {
		t.Fatalf("decoded error = %v, want %q", decoded.Error, sqlErr.Error())
	}
}
