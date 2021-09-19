package flexjson

import (
	"git.kanosolution.net/kano/dbflex"
)

type Cursor struct {
	e error
}

func (cur *Cursor) Reset() error {
	panic("not implemented") // TODO: Implement
}

func (cur *Cursor) Fetch(_ interface{}) dbflex.ICursor {
	panic("not implemented") // TODO: Implement
}

func (cur *Cursor) Fetchs(_ interface{}, _ int) dbflex.ICursor {
	panic("not implemented") // TODO: Implement
}

func (cur *Cursor) Count() int {
	panic("not implemented") // TODO: Implement
}

func (cur *Cursor) CountAsync() <-chan int {
	panic("not implemented") // TODO: Implement
}

func (cur *Cursor) Close() error {
	panic("not implemented") // TODO: Implement
}

func (cur *Cursor) Error() error {
	return cur.e
}

func (cur *Cursor) SetCountCommand(_ dbflex.ICommand) {
	panic("not implemented") // TODO: Implement
}

func (cur *Cursor) CountCommand() dbflex.ICommand {
	panic("not implemented") // TODO: Implement
}

func (cur *Cursor) Connection() dbflex.IConnection {
	panic("not implemented") // TODO: Implement
}

func (cur *Cursor) SetConnection(_ dbflex.IConnection) {
	panic("not implemented") // TODO: Implement
}

func (cur *Cursor) ConfigRef(key string, def interface{}, out interface{}) {
	panic("not implemented") // TODO: Implement
}

func (cur *Cursor) Set(key string, value interface{}) {
	panic("not implemented") // TODO: Implement
}
