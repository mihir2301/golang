package controller

import (
	"context"
	"fmt"
	"log"
	"mongodb/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://mihir:mihir@cluster0.2aawb3m.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString) //establishing clientoption

	client, err := mongo.Connect(context.TODO(), clientOption) //establishing connection

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongodb connection success")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready")
}

//mongo helper

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id:", inserted.InsertedID)
}

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId) // gets the id from string
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result,err:=collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err) //rresult gives how many are updated
	}
	fmt.Println("modified count",result.ModifiedCount)
}
