package nsql

import (
	"github.com/jmoiron/sqlx"
)

func NewDB(driverName, dsn string) DB {
	db := sqlx.MustOpen(driverName, dsn)
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(100)
	return &nfDB{DB: db}
}

type DB interface {
	preparer
	accesser
	Begin() Tx
	Close()
}

type Tx interface {
	preparer
	accesser
	Commit()
	Rollback()
}

type Stmt interface {
	Get(dest interface{}, args ...interface{}) bool

	Select(dest interface{}, args ...interface{})

	Exec(args ...interface{}) Result
}

type NamedStmt interface {
	NamedGet(dest interface{}, arg interface{}) bool

	NamedSelect(dest interface{}, arg interface{})

	NamedExec(arg interface{}) Result
}

type Result interface {
	LastInsertId() int64
	RowsAffected() int64
}

type accesser interface {
	NamedGet(dest interface{}, query string, arg interface{}) bool

	Get(dest interface{}, query string, args ...interface{}) bool

	NamedSelect(dest interface{}, query string, arg interface{})

	Select(dest interface{}, query string, args ...interface{})

	NamedExec(query string, arg interface{}) Result

	Exec(query string, args ...interface{}) Result
}

type preparer interface {
	Prepare(query string) Stmt

	PrepareNamed(query string) NamedStmt
}
