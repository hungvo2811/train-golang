package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	db *mongo.Database
)

func ConnectDB() {
	cl, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://SNW:hungpro10a3@snw.wxnes.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Println(err)
		log.Fatal("Cannot connect to database:")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = cl.Connect(ctx)
	if err != nil {
		log.Println(err)
	}
	db = cl.Database("SNW")
	fmt.Println("Connected to MongoDB")
}
