package hellodb

import "log"

func Action_StatementsThatModifyData() {
	// Use Exec(), preferably with a prepared statement, to accomplish an INSERT, UPDATE, DELETE,
	// or another statement that doesn’t return rows.
	db := createDB()
	stmt, err := db.Prepare("INSERT INTO user(username) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("Dolly")
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

	// Executing the statement produces a sql.Result that gives access to statement metadata: the last inserted ID and the number of
	// rows affected.
	//
	//What if you don’t care about the result? What if you just want to execute a statement and check if there were any errors,
	// but ignore the result? Wouldn’t the following two statements do the same thing?
	//
	//_, err := db.Exec("DELETE FROM users")  // OK
	//_, err := db.Query("DELETE FROM users") // BAD
}

/**
Working with Transactions

In Go, a transaction is essentially an object that reserves a connection to the datastore. It lets you do all of the operations
we’ve seen thus far, but guarantees that they’ll be executed on the same connection.

You begin a transaction with a call to db.Begin(), and close it with a Commit() or Rollback() method on the resulting Tx variable.
Under the covers, the Tx gets a connection from the pool, and reserves it for use only with that transaction.
The methods on the Tx map one-for-one to methods you can call on the database itself, such as Query() and so forth.
*/
