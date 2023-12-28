package books

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handle requests to get books
func GetBooks(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Requested for all books")
}

// Handle requests to get a specific book
func GetSpecificBook(res http.ResponseWriter, req *http.Request) {
	var params map[string]string = mux.Vars(req)
	var bookName string = params["book"]

	fmt.Fprintf(res, "Requested for the book: %s", bookName)
}
