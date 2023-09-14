package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id, omitempty"`
	Email    string             `bson:"email" json:"email" unique:"true" index:"unique"`
	Password string             `bson:"password" json:"password"`
}
