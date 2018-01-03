package nsql

import (
	"github.com/jmoiron/sqlx"
)

type nfTx struct {
	*sqlx.Tx
}

func (tx *nfTx) NamedGet(dest interface{}, query string, arg interface{}) bool {
	stmt, err := tx.Tx.PrepareNamed(query)
	panicErr(err)
	return namedGet(stmt, dest, arg)
}

func (tx *nfTx) Get(dest interface{}, query string, args ...interface{}) bool {
	return get(tx.Tx, dest, query, args...)
}

func (tx *nfTx) NamedSelect(dest interface{}, query string, arg interface{}) {
	stmt, err := tx.Tx.PrepareNamed(query)
	panicErr(err)
	namedSelect(stmt, dest, arg)
}

func (tx *nfTx) Select(dest interface{}, query string, args ...interface{}) {
	selectx(tx.Tx, dest, query, args...)
}

func (tx *nfTx) NamedExec(query string, arg interface{}) Result {
	return namedExec(tx.Tx, query, arg)
}

func (tx *nfTx) Exec(query string, args ...interface{}) Result {
	return exec(tx.Tx, query, args...)
}

func (tx *nfTx) Commit() {
	err := tx.Tx.Commit()
	panicErr(err)
}

func (tx *nfTx) Rollback() {
	err := tx.Tx.Rollback()
	panicErr(err)
}

func (tx *nfTx) Prepare(query string) Stmt {
	stmt, err := tx.Tx.Preparex(query)
	panicErr(err)
	return &nfStmt{Stmt: stmt}
}

func (db *nfTx) PrepareNamed(query string) NamedStmt {
	stmt, err := db.Tx.PrepareNamed(query)
	panicErr(err)
	return &nfNamedStmt{NamedStmt: stmt}
}
