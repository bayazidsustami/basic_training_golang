package utils

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		PanicErr(errorRollback)
	} else {
		errorCommit := tx.Commit()
		PanicErr(errorCommit)
	}
}
