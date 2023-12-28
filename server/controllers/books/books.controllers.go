package books

import (
	"database/sql"
	"encoding/json"
	"fmt"
	bookModel "go-web-server/server/models/book"
	"go-web-server/server/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Handle requests to get books
func GetBooks(res http.ResponseWriter, req *http.Request, appDatabase *sql.DB) {
	res.Header().Set("Content-Type", "application/json")

	const getAllBooksQuery string = `SELECT * FROM books`
	var books []bookModel.Book
	var response utils.Response

	rows, err := appDatabase.Query(getAllBooksQuery)
	if err != nil {
		response = utils.Response{
			Msg:     err.Error(),
			Success: false,
			Payload: nil,
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			log.Fatal("Failed to serialize response object.")
		}

		res.WriteHeader(http.StatusBadRequest)
		res.Write(jsonResponse)
	}

	for rows.Next() {
		var title, author string
		if err := rows.Scan(&title, &author); err != nil {
			log.Fatal(err)
		}

		books = append(books, bookModel.Book{Title: title, Author: author})
	}

	response = utils.Response{
		Msg:     "Successfully fetched books.",
		Success: true,
		Payload: books,
	}

	jsonResponseData, err := json.Marshal(response)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		log.Fatal("Failed to serialize response object.")
	}

	res.WriteHeader(http.StatusOK)
	res.Write(jsonResponseData)
}

// Handle requests to get a specific book
func GetSpecificBook(res http.ResponseWriter, req *http.Request) {
	var params map[string]string = mux.Vars(req)
	var bookName string = params["book"]

	fmt.Fprintf(res, "Requested for the book: %s", bookName)
}
