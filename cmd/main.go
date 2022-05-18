package main

import (
	"cchallenge/internal/bd"
	"cchallenge/internal/handlers"
	"log"
)

func main() {
	//err := godotenv.Load(".env")
	//
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}

	// Initialize the bd
	if bd.CheckConnection() == 0 {
		log.Fatal("Error connecting to the database")
		return
	}

	handlers.Handlers()
}
