package routes

import (
	"github.com/gorilla/mux"
	"api/controllers"
)

// SetContactsRoutes agrega las rutas de contactos
func SetBooksRoutes(r *mux.Router) {

	subRouter := r.PathPrefix("/api").Subrouter()

	// Route handles & endpoints
	subRouter.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	subRouter.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
	subRouter.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	subRouter.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	subRouter.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
}