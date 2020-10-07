package main

import (
	"github.com/gorilla/mux"
	"api/database"
	"api/routes"
	"log"
	"net/http"
)


// Main function
func main() {

	// Database
	database.Connect("mongo")

	// Init router
	r := mux.NewRouter()

	// Add Routes
	routes.SetBooksRoutes(r)
	routes.SetHealthCheckRoutes(r)

	// Create Server
	srv := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("Running on port 8080")

	// Run Server
	log.Println(srv.ListenAndServe())
}