package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/lgustavopalmieri/labs-go-expert-auctiont/configuration/database/mongodb"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatal("Error trying to load env variables")
		return
	}

	_, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
