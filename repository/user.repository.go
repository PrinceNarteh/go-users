package repository

import (
	"context"
	"log"
	"time"

	"github.com/PrinceNarteh/go-users/configs"
	"github.com/PrinceNarteh/go-users/models"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func GetUserById(userId string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Fatal(err)
	}

	if err := userCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&user); err != nil {
		return models.User{}, err
	}

	return user, nil
}
