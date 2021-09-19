package flexjson

import (
	"fmt"
	"path/filepath"
	"sync"

	"git.kanosolution.net/kano/dbflex"
	"github.com/eaciit/toolkit"
)

type Connection struct {
	dbflex.ConnectionBase
	sync.RWMutex
	dataPath string
	state    string
}

func (c *Connection) Connect() error {
	panic("not implemented") // TODO: Implement
}

func (c *Connection) State() string {
	return c.state
}

func (c *Connection) Close() {
}

/*
func (c *Connection) Prepare(cmd dbflex.ICommand) (dbflex.IQuery, error) {
	//panic("not implemented") // TODO: Implement
	q := new(Query)
	q.SetCommand(cmd)
	return q, nil
}
*/

func (c *Connection) Execute(_ dbflex.ICommand, _ toolkit.M) (interface{}, error) {
	panic("not implemented") // TODO: Implement
}

func (c *Connection) Cursor(_ dbflex.ICommand, _ toolkit.M) dbflex.ICursor {
	panic("not implemented") // TODO: Implement
}

func (c *Connection) NewQuery() dbflex.IQuery {
	query := new(Query)
	query.SetConnection(c)
	return query
}

func (c *Connection) ObjectNames(_ dbflex.ObjTypeEnum) []string {
	panic("not implemented") // TODO: Implement
}

func (c *Connection) ValidateTable(_ interface{}, _ bool) error {
	panic("not implemented") // TODO: Implement
}

func (c *Connection) DropTable(_ string) error {
	panic("not implemented") // TODO: Implement
}

func (c *Connection) HasTable(_ string) bool {
	panic("not implemented") // TODO: Implement
}

func (c *Connection) EnsureTable(_ string, _ []string, _ interface{}) error {
	panic("not implemented") // TODO: Implement
}

func (c *Connection) BeginTx() error {
	return fmt.Errorf("tx is not implemented")
}

func (c *Connection) Commit() error {
	return fmt.Errorf("tx is not implemented")
}

func (c *Connection) RollBack() error {
	return fmt.Errorf("tx is not implemented")
}

func (c *Connection) SupportTx() bool {
	return false
}

func (c *Connection) IsTx() bool {
	panic("not implemented") // TODO: Implement
}

func (c *Connection) SetThis(_ dbflex.IConnection) dbflex.IConnection {
	panic("not implemented") // TODO: Implement
}

func (c *Connection) This() dbflex.IConnection {
	panic("not implemented") // TODO: Implement
}

func (c *Connection) SetFieldNameTag(_ string) {
	panic("not implemented") // TODO: Implement
}

func (c *Connection) FieldNameTag() string {
	panic("not implemented") // TODO: Implement
}

func (c *Connection) SetKeyNameTag(_ string) {
	panic("not implemented") // TODO: Implement
}

func (c *Connection) KeyNameTag() string {
	panic("not implemented") // TODO: Implement
}

func (c *Connection) jsonPath(tableName string) string {
	return filepath.Join(c.dataPath, fmt.Sprintf("%s_flexjson_table.json", tableName))
}
