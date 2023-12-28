package routes

import (
	"database/sql"
	"go-web-server/server/controllers/books"
	"net/http"

	"github.com/gorilla/mux"
)

// Handle book routes
func InitializeBookRoutes(router *mux.Router, appDatabase *sql.DB) {
	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		books.GetBooks(res, req, appDatabase)
	}).Methods("GET")
	router.HandleFunc("/{book}", books.GetSpecificBook).Methods("GET")
}
