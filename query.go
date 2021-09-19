package flexjson

import (
	"git.kanosolution.net/kano/dbflex"
	"github.com/eaciit/toolkit"
)

type Query struct {
	this dbflex.IQuery
}

func (q *Query) SetThis(qr dbflex.IQuery) {
	q.this = qr
}

func (q *Query) This() dbflex.IQuery {
	if q.this == nil {
		q.this = q
	}
	return q.this
}

func (q *Query) BuildFilter(_ *dbflex.Filter) (interface{}, error) {
	panic("not implemented") // TODO: Implement
}

func (q *Query) BuildCommand() (interface{}, error) {
	panic("not implemented") // TODO: Implement
}

func (q *Query) Cursor(_ toolkit.M) dbflex.ICursor {
	panic("not implemented") // TODO: Implement
}

func (q *Query) Execute(_ toolkit.M) (interface{}, error) {
	panic("not implemented") // TODO: Implement
}

func (q *Query) SetConfig(_ string, _ interface{}) {
	panic("not implemented") // TODO: Implement
}

func (q *Query) SetConfigM(_ toolkit.M) {
	panic("not implemented") // TODO: Implement
}

func (q *Query) Config(_ string, _ interface{}) interface{} {
	panic("not implemented") // TODO: Implement
}

func (q *Query) ConfigRef(_ string, _ interface{}, _ interface{}) {
	panic("not implemented") // TODO: Implement
}

func (q *Query) DeleteConfig(_ ...string) {
	panic("not implemented") // TODO: Implement
}

func (q *Query) Connection() dbflex.IConnection {
	panic("not implemented") // TODO: Implement
}

func (q *Query) SetConnection(_ dbflex.IConnection) {
	panic("not implemented") // TODO: Implement
}

func (q *Query) SetCommand(_ dbflex.ICommand) dbflex.IQuery {
	panic("not implemented") // TODO: Implement
}

func (q *Query) Command() dbflex.ICommand {
	panic("not implemented") // TODO: Implement
}
