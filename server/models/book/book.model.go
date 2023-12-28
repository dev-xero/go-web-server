package bookModel

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
)

type Book struct {
	Title  string
	Author string
}

func InsertBook(book Book, appDatabase *sql.DB) error {
	const insertBookQuery string = "INSERT INTO books (title, author) VALUES ($1, $2)"
	_, err := appDatabase.Exec(insertBookQuery, book.Title, book.Author)

	if err != nil {
		// Check if the table doesn't exist
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code.Name() == "undefined_table" {
			const createTableQuery string = `
				CREATE TABLE IF NOT EXISTS books (
					title VARCHAR(255),
					author VARCHAR(255)
				)
			`
			_, err := appDatabase.Exec(createTableQuery)
			if err != nil {
				log.Fatal(err)
			}

			// Try inserting again
			_, insertionErr := appDatabase.Exec(insertBookQuery, book.Title, book.Author)
			if insertionErr != nil {
				log.Fatal(insertionErr)
			}
		} else {
			log.Fatal(err)
		}
	}

	fmt.Println("Successfully inserted item into the database")

	return nil
}
