package main

import (
	"database/sql"
	"fmt"
	"github.com/samonzeweb/godb/tablenamer"
	"time"

	"github.com/samonzeweb/godb"
	"github.com/samonzeweb/godb/adapters/sqlite"
	"log"
	"os"
)

/*
  To run this example, initialize a SQLite3 DB called 'library.db' and add
  a 'books' table like this :

  create table books (
  	id        integer not null primary key autoincrement,
  	title     text    not null,
  	author    text    not null,
  	published date    not null);
*/

// Struct and its mapping
type Book struct {
	Id        int       `db:"id,key,auto"`
	Title     string    `db:"title"`
	Author    string    `db:"author"`
	Published time.Time `db:"published"`
}

// Optional, default if the struct name (Book)
func (*Book) TableName() string {
	return "books"
}

// See "group by" example
type CountByAuthor struct {
	Author string `db:"author"`
	Count  int    `db:"count"`
}

func main() {
	// Examples fixtures
	var authorTolkien = "J.R.R. tolkien"

	var bookTheHobbit = Book{
		Title:     "The Hobbit",
		Author:    authorTolkien,
		Published: time.Date(1937, 9, 21, 0, 0, 0, 0, time.UTC),
	}

	var bookTheFellowshipOfTheRing = Book{
		Title:     "The Fellowship of the Ring",
		Author:    authorTolkien,
		Published: time.Date(1954, 7, 29, 0, 0, 0, 0, time.UTC),
	}

	var bookTheTwoTowers = Book{
		Title:     "The Two Towers",
		Author:    authorTolkien,
		Published: time.Date(1954, 11, 11, 0, 0, 0, 0, time.UTC),
	}

	var bookTheReturnOfTheKing = Book{
		Title:     "The Return of the King",
		Author:    authorTolkien,
		Published: time.Date(1955, 10, 20, 0, 0, 0, 0, time.UTC),
	}

	var setTheLordOfTheRing = []Book{
		bookTheFellowshipOfTheRing,
		bookTheTwoTowers,
		bookTheReturnOfTheKing,
	}

	// Connect to the DB
	db, err := godb.Open(sqlite.Adapter, "./library.db")
	panicIfErr(err)
	// OPTIONAL: Set logger to show SQL execution logs
	db.SetLogger(log.New(os.Stderr, "", 0))
	// OPTIONAL: Set default table name building style from struct's name(if active struct doesn't have TableName() method)
	db.SetDefaultTableNamer(tablenamer.Plural())
	// Single insert (id will be updated)
	err = db.Insert(&bookTheHobbit).Do()
	panicIfErr(err)

	// Multiple insert
	// Warning : BulkInsert only update ids with PostgreSQL and SQL Server!
	err = db.BulkInsert(&setTheLordOfTheRing).Do()
	panicIfErr(err)

	// Count
	count, err := db.SelectFrom("books").Count()
	panicIfErr(err)
	fmt.Println("Books count : ", count)

	// Custom select
	countByAuthor := make([]CountByAuthor, 0, 0)
	err = db.SelectFrom("books").
		Columns("author", "count(*) as count").
		GroupBy("author").
		Having("count(*) > 3").
		Do(&countByAuthor)
	fmt.Println("Count by authors : ", countByAuthor)

	// Select single object
	singleBook := Book{}
	err = db.Select(&singleBook).
		Where("title = ?", bookTheHobbit.Title).
		Do()
	if err == sql.ErrNoRows {
		// sql.ErrNoRows is only returned when the target is a single instance
		fmt.Println("Book not found !")
	} else {
		panicIfErr(err)
	}

	// Select single record values
	authorName := ""
	title := ""
	err = db.SelectFrom("books").
		Where("title = ?", bookTheHobbit.Title).
		Columns("author", "title").
		Scanx(&authorName, &title)
	if err == sql.ErrNoRows {
		// sql.ErrNoRows is only returned when the target is a single instance
		fmt.Println("Book not found !")
	} else {
		panicIfErr(err)
	}

	// Select multiple objects
	multipleBooks := make([]Book, 0, 0)
	err = db.Select(&multipleBooks).Do()
	panicIfErr(err)
	fmt.Println("Books found : ", len(multipleBooks))

	// Iterator
	iter, err := db.SelectFrom("books").
		Columns("id", "title", "author", "published").
		DoWithIterator()
	panicIfErr(err)
	for iter.Next() {
		book := Book{}
		err := iter.Scan(&book)
		panicIfErr(err)
		fmt.Println(book)
	}
	panicIfErr(iter.Err())
	panicIfErr(iter.Close())

	// Raw query
	subQuery := godb.NewSQLBuffer(0, 0). // sizes are indicative
						Write("select author ").
						Write("from books ").
						WriteCondition(godb.Q("where title = ?", bookTheHobbit.Title))

	queryBuffer := godb.NewSQLBuffer(64, 0).
		Write("select * ").
		Write("from books ").
		Write("where author in (").
		Append(subQuery).
		Write(")")

	panicIfErr(queryBuffer.Err())

	books := make([]Book, 0, 0)
	err = db.RawSQL(queryBuffer.SQL(), queryBuffer.Arguments()...).Do(&books)
	panicIfErr(err)
	fmt.Printf("Raw query found %d books\n", len(books))

	// Update and transactions
	err = db.Begin()
	panicIfErr(err)

	updated, err := db.UpdateTable("books").Set("author", "Tolkien").Do()
	panicIfErr(err)
	fmt.Println("Books updated : ", updated)

	bookTheHobbit.Author = "Tolkien"
	err = db.Update(&bookTheHobbit).Do()
	panicIfErr(err)
	fmt.Println("Books updated : ", updated)

	err = db.Rollback()
	panicIfErr(err)

	// Delete
	deleted, err := db.Delete(&bookTheHobbit).Do()
	panicIfErr(err)
	fmt.Println("Books deleted : ", deleted)

	deleted, err = db.DeleteFrom("books").
		WhereQ(godb.Or(
			godb.Q("author = ?", authorTolkien),
			godb.Q("author = ?", "Georged Orwell"),
		)).
		Do()
	panicIfErr(err)
	fmt.Println("Books deleted : ", deleted)

	// Bye
	err = db.Close()
	panicIfErr(err)
}

// It's just an example, what did you expect ? (never do that in real code)
func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
