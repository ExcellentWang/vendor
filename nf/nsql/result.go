package nsql

import (
	"database/sql"
)

type nfResult struct {
	sql.Result
}

func (r *nfResult) LastInsertId() int64 {
	id, err := r.Result.LastInsertId()
	panicErr(err)
	return id
}

func (r *nfResult) RowsAffected() int64 {
	rows, err := r.Result.RowsAffected()
	panicErr(err)
	return rows
}
