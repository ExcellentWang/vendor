package nsql

import (
	"github.com/jmoiron/sqlx"
)

type nfDB struct {
	*sqlx.DB
}

func (db *nfDB) Get(dest interface{}, query string, args ...interface{}) bool {
	return get(db.DB, dest, query, args...)
}

func (db *nfDB) NamedGet(dest interface{}, query string, arg interface{}) bool {
	stmt, err := db.DB.PrepareNamed(query)
	panicErr(err)
	return namedGet(stmt, dest, arg)
}

func (db *nfDB) NamedSelect(dest interface{}, query string, arg interface{}) {
	stmt, err := db.DB.PrepareNamed(query)
	panicErr(err)
	namedSelect(stmt, dest, arg)
}

func (db *nfDB) Select(dest interface{}, query string, args ...interface{}) {
	selectx(db.DB, dest, query, args...)
}

func (db *nfDB) NamedExec(query string, arg interface{}) Result {
	return namedExec(db.DB, query, arg)
}

func (db *nfDB) Exec(query string, args ...interface{}) Result {
	return exec(db.DB, query, args...)
}

func (db *nfDB) Begin() Tx {
	tx := db.DB.MustBegin()
	return &nfTx{Tx: tx}
}

func (db *nfDB) Prepare(query string) Stmt {
	stmt, err := db.DB.Preparex(query)
	panicErr(err)
	return &nfStmt{Stmt: stmt}
}

func (db *nfDB) PrepareNamed(query string) NamedStmt {
	stmt, err := db.DB.PrepareNamed(query)
	panicErr(err)
	return &nfNamedStmt{NamedStmt: stmt}
}

func (db *nfDB) Close() {
	db.DB.Close()
}
