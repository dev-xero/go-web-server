package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-web-server/server/routes"
	"go-web-server/server/utils"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func initializeRoutes(appRouter *mux.Router, appDatabase *sql.DB) {
	bookRouter := appRouter.PathPrefix("/books/").Subrouter()

	routes.InitializeBookRoutes(bookRouter, appDatabase)

	// Handle requests to home
	appRouter.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		response := utils.Response{
			Msg:     "Welcome to the API.",
			Success: true,
			Payload: nil,
		}

		jsonResponseData, err := json.Marshal(response)
		if err != nil {
			http.Error(res, "Internal server error.", http.StatusInternalServerError)
			log.Fatal(err)
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		res.Write(jsonResponseData)
	})
}

func connectToDB() *sql.DB {
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

	return db
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
	appDatabase := connectToDB()

	initializeRoutes(appRouter, appDatabase)
	listenForRequests(port, appRouter)
}
