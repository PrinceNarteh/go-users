package configs

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI"))); err != nil {
		log.Fatal(err)
	} else {
		DB = client
	}
}

func GetCollection(db *mongo.Client, collectionName string) *mongo.Collection {
	return db.Database("go_users").Collection(collectionName)
}
