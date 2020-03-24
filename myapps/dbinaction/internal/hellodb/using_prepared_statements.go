package hellodb

import "log"

/**
If you don’t want to use a prepared statement, you need to use fmt.Sprint() or similar to assemble the SQL,
and pass this as the only argument to db.Query() or db.QueryRow(). And your driver needs to support plaintext query execution,
which is added in Go 1.1 via the Execer and Queryer interfaces,
*/

func Action_PreparedStatementsInTransactions() {
	db := createDB()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare("INSERT INTO foo VALUES (?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // danger!
	for i := 0; i < 10; i++ {
		_, err = stmt.Exec(i)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	// stmt.Close() runs here!
}
