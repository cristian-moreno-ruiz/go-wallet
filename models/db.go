package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/go-shadow/moment"
)

// TODO: Maybe I need to add sync.once in order to ensure using a singleton here
var database *mongo.Database
var ctx = context.TODO()

// Connect Connects to DB
func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	database = client.Database("go-experiment")
	logStartup()
}

// This is just a dummy test to be sure I can access db
func logStartup() {
	type Startup struct {
		ID     primitive.ObjectID `bson:"_id"`
		Date   time.Time          `bson:"created_at"`
		Text   string             `bson:"startup successful"`
		Status bool               `bson:"completed"`
	}

	startup := &Startup{
		ID:     primitive.NewObjectID(),
		Date:   moment.New().GetTime(),
		Text:   "test",
		Status: true,
	}

	database.Collection("startups").InsertOne(ctx, startup)
	fmt.Println("Successfully started")
}
