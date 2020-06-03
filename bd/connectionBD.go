package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN is the Object Connection to a Data Base.
var MongoCN = ConnectionBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://root:twittor2020@twittor-9cnyp.mongodb.net/test")

//Function Connect to a Data Base.
func ConnectionBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Connection Succes!!")
	return client
}

//Function Check Ping Connection to a Data Base 0 Error or 1 Succes!.
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
