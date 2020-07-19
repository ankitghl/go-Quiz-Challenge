package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Initialise Database
func InitialiseDatabase() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb+srv://user2:user2@cluster0-ovyhe.mongodb.net/test?retryWrites=true&w=majority")
	client, _ = mongo.Connect(ctx, clientOptions)
}

// Will return Client Connection or Error
func GetDatabaseCollection(name string) (*mongo.Collection, error) {
	if client == nil {
		InitialiseDatabase()
	}
	collection := client.Database("QuizChallenge").Collection("Users")

	return collection, nil
}
