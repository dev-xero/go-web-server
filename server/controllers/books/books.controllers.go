package books

import (
	"database/sql"
	"fmt"
	"go-web-server/server/models/book"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Handle requests to get books
func GetBooks(res http.ResponseWriter, req *http.Request, appDatabase *sql.DB) {
	const getAllBooksQuery string = `SELECT * FROM books`
	var books []bookModel.Book

	rows, err := appDatabase.Query(getAllBooksQuery)
	if err != nil {
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
	}

	for rows.Next() {
		var title, author string
		if err := rows.Scan(&title, &author); err != nil {
			log.Fatal(err)
		}

		books = append(books, bookModel.Book{Title: title, Author: author})
	}

	fmt.Fprint(res, books)
}

// Handle requests to get a specific book
func GetSpecificBook(res http.ResponseWriter, req *http.Request) {
	var params map[string]string = mux.Vars(req)
	var bookName string = params["book"]

	fmt.Fprintf(res, "Requested for the book: %s", bookName)
}
