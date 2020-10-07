package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Book struct {
	ID     		primitive.ObjectID  `json:"id" bson:"_id"`
	Isbn   		string  			`json:"isbn" bson:"isbn"`
	Title  		string  			`json:"title" bson:"title"`
	Author 		*Author 			`json:"author" bson:"author"`
	CreatedAt 	time.Time			`json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time			`json:"updatedAt" bson:"updatedAt"`
}
