package nsql

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

func namedGet(stmt *sqlx.NamedStmt, dest interface{}, arg interface{}) bool {
	err := stmt.Get(dest, arg)
	if err == sql.ErrNoRows {
		return false
	}
	panicErr(err)
	return true
}

func get(queryer sqlx.Queryer, dest interface{}, query string, args ...interface{}) bool {
	err := sqlx.Get(queryer, dest, query, args...)
	if err == sql.ErrNoRows {
		return false
	}
	panicErr(err)
	return true
}

func namedSelect(stmt *sqlx.NamedStmt, dest interface{}, arg interface{}) {
	err := stmt.Select(dest, arg)
	panicErr(err)
}

func selectx(queryer sqlx.Queryer, dest interface{}, query string, args ...interface{}) {
	err := sqlx.Select(queryer, dest, query, args...)
	panicErr(err)
}

func exec(e sqlx.Execer, query string, args ...interface{}) Result {
	result := sqlx.MustExec(e, query, args...)
	return &nfResult{result}
}

func namedExec(ext sqlx.Ext, query string, arg interface{}) Result {
	result, err := sqlx.NamedExec(ext, query, arg)
	panicErr(err)
	return &nfResult{result}
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}
