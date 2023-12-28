package routes

import (
	"go-web-server/server/controllers/books"

	"github.com/gorilla/mux"
)

// Handle book routes
func InitializeBookRoutes(router *mux.Router) {
	router.HandleFunc("/", books.GetBooks).Methods("GET")
	router.HandleFunc("/{book}", books.GetSpecificBook).Methods("GET")
}
