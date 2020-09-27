package main

import (
	"github.com/gorilla/mux"
	"go-rollercoaster-api/database"
	"go-rollercoaster-api/routes"
	"log"
	"net/http"
)


// Main function
func main() {

	// Database
	database.Connect("localhost")

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