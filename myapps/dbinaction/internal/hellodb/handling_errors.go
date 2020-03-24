package hellodb

import (
	"database/sql"
	"fmt"
	"log"
)

func Action_ErrorsFromQueryRow() {
	db := createDB()

	var name string
	id := 1000000
	err := db.QueryRow("select username from user where id = ?", id).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			// there were no rows, but otherwise no error occurred
			log.Println("there were no rows for ID: ", id)
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println(name)
}
