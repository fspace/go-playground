package hellodb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Main() {
	db, err := sql.Open("mysql",
		"yiqing:yiqing@tcp(127.0.0.1:3306)/playgo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		// do something here
		log.Fatal(err)
	}
	log.Println("now you can access the database ")
}

// -----------------------------------------------------------------
// ## Retrieving Result Sets
func createDB() *sql.DB {
	db, err := sql.Open("mysql",
		"yiqing:yiqing@tcp(127.0.0.1:3306)/playgo")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
func Action_FetchingData() {
	var (
		id   int
		name string
	)
	db := createDB()
	rows, err := db.Query("select id, username from user where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

// -----------------------------------------------------------------

func Action_PreparingQueries() {
	var (
		id   int
		name string
	)

	db := createDB()
	stmt, err := db.Prepare("select id, username from user where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		// ...
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

// -----------------------------------------------------------------
func Action_SingleRowQueries() {
	var name string

	db := createDB()
	err := db.QueryRow("select username from user where id = ?", 1).Scan(&name)
	// Errors from the query are deferred until Scan() is called, and then are returned from that.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
	// You can also call QueryRow() on a prepared statement
	fmt.Println("============= query-row on prepared statement ================")
	stmt, err := db.Prepare("select username from user where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	//var name string
	err = stmt.QueryRow(1).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}

// -----------------------------------------------------------------
