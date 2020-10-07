package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Author struct
type Author struct {
	ID     		primitive.ObjectID  `json:"id" bson:"_id"`
	Firstname 	string 				`json:"firstname"`
	Lastname 	string 				`json:"lastname"`
	CreatedAt 	time.Time			`json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time			`json:"updatedAt" bson:"updatedAt"`
}