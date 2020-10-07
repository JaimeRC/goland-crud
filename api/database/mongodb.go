package database

import (
	"context"
	"fmt"
	"api/controllers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)


func Connect(host string) {

	// Database Config
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", "mongodb", 27018))






	
	//Set up a context required by mongo.Connect
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	//To close the connection at the end
	//defer cancel()

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		fmt.Println("Couldn't connect to the database", err)
	} else {
		fmt.Println("Connected to MongoDB!")
	}

	db := client.Database("apiBook")

	controllers.BookCollection(db)
}
