package bd

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var MongoCN = ConnectDB()

//var clientOptions = options.Client().ApplyURI("mongodb://productListUser:productListPassword@localhost:27017/?authMechanism=SCRAM-SHA-1")
var clientOptions = options.Client().ApplyURI("mongodb+srv://dbwalmart:1234567890@cluster0.6leke.mongodb.net/test?retryWrites=true&w=majority")

func ConnectDB() *mongo.Client {
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
	log.Println("Conexión exitosa con la BD")
	return client
}

func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
