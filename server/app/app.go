package app

import (
	"database/sql"
	"fmt"
	"go-web-server/server/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func initializeRoutes(appRouter *mux.Router) {
	bookRouter := appRouter.PathPrefix("/books/").Subrouter()

	routes.InitializeBookRoutes(bookRouter)

	// Handle requests to home
	appRouter.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Welcome to the API.")
	})
}

func connectToDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("An error occurred while loading .env file.")
	}

	var connString string = os.Getenv("DB_CONNECTION_STRING")

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("Failed to connect to remote postgreSQL database.")
	}
	fmt.Println("Successfully connected to remote postgreSQL database.")

	pingError := db.Ping()
	if pingError != nil {
		log.Fatal("An error occurred while pinging the remote postgreSQL database.")
	}
}

func listenForRequests(port int, appRouter *mux.Router) {
	fmt.Printf("Server listening for requests on http://localhost:%d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), appRouter)
	// Handle any errors while starting the server
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}

func Initialize(port int) {
	appRouter := mux.NewRouter()

	connectToDB()
	initializeRoutes(appRouter)
	listenForRequests(port, appRouter)
}