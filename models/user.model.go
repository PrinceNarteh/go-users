package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id    primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Title string             `bson:"title" json:"title,omitempty" validate:"required"`
	Name  string             `bson:"name" json:"name,omitempty" validate:"required"`
	Email string             `bson:"email" json:"email,omitempty" validate:"required"`
}
