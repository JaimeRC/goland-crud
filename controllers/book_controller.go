package controllers

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go-rollercoaster-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// DATABASE INSTANCE
var collection *mongo.Collection

func BookCollection(c *mongo.Database) {
	collection = c.Collection("books")
}

// Get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := []models.Book{}
	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error while getting all books, Reason: %v\n", err)
		http.Error(w, err.Error(), 500)
		return
	}

	// Iterate through the returned cursor.
	for cursor.Next(context.TODO()) {
		var book models.Book
		cursor.Decode(&book)
		books = append(books, book)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get single book
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params

	book := models.Book{}
	err := collection.FindOne(context.TODO(), bson.M{"isdn": params["isdn"]}).Decode(&book)
	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(book)
}

// Add new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.Isbn = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe

	_, err := collection.InsertOne(context.TODO(), book)

	if err != nil {
		log.Printf("Error while inserting new book into db, Reason: %v\n", err)
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(book)
}

// Update book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var updateBook models.Book

	_ = json.NewDecoder(r.Body).Decode(&updateBook)

	newData := bson.M{
		"$set": bson.M{
			"title":       updateBook.Title,
			"author": 		updateBook.Author,
			"updatedAt":	time.Now(),
		},
	}

	_, err := collection.UpdateOne(context.TODO(), bson.M{"isbn": params["isbn"]}, newData)
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(updateBook)

}

// Delete book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	_, err := collection.DeleteOne(context.TODO(), bson.M{"isbn": params["isbn"]})
	if err != nil {
		log.Printf("Error while deleting a single todo, Reason: %v\n", err)
		http.Error(w, err.Error(), 500)
		return
	}

	GetBooks(w,r)
}

