package configs

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func ConnectDB() (*DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI"))); err != nil {
		return nil, err
	} else {
		return &DB{
			client: client,
		}, nil
	}
}

func (db *DB) GetCollection(collectionName string) *mongo.Collection {
	return db.client.Database("go_users").Collection(collectionName)
}
