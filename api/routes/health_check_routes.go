package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"api/models"
	"net/http"
)



func SetHealthCheckRoutes(r *mux.Router) {

	// Route handles & endpoints
	r.HandleFunc("/health_check", getHealthCheck).Methods("GET")
}

// Get all books
func getHealthCheck(w http.ResponseWriter, r *http.Request) {
	var response models.Status
	response.Status = true
	response.Message = "Server is running!!!"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}