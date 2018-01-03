package nsql

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type nfStmt struct {
	*sqlx.Stmt
}

func (nf *nfStmt) Get(dest interface{}, args ...interface{}) bool {
	err := nf.Stmt.Get(dest, args...)
	if err == sql.ErrNoRows {
		return false
	}
	panicErr(err)
	return true
}

func (nf *nfStmt) Select(dest interface{}, args ...interface{}) {
	err := nf.Stmt.Select(dest, args...)
	panicErr(err)
}

func (nf *nfStmt) Exec(args ...interface{}) Result {
	result := nf.Stmt.MustExec(args...)
	return &nfResult{result}
}

type nfNamedStmt struct {
	*sqlx.NamedStmt
}

func (nf *nfNamedStmt) NamedGet(dest interface{}, arg interface{}) bool {
	return namedGet(nf.NamedStmt, dest, arg)
}

func (nf *nfNamedStmt) NamedSelect(dest interface{}, arg interface{}) {
	namedSelect(nf.NamedStmt, dest, arg)
}
func (nf *nfNamedStmt) NamedExec(arg interface{}) Result {
	result := nf.NamedStmt.MustExec(arg)
	return &nfResult{result}
}
