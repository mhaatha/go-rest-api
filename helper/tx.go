package helper

import (
	"database/sql"
	"log"
)

func CommitOrRollback(tx *sql.Tx) {
	err := recover()

	if err != nil {
		if errorRollback := tx.Rollback(); errorRollback != nil {
			log.Printf("rollback failed: %v", errorRollback)
		}
		panic(err)
	}

	if errorCommit := tx.Commit(); errorCommit != nil {
		panic(errorCommit)
	}
}
