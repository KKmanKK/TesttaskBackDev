package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Token struct {
	ID           primitive.ObjectID `bson:"_id, omitempty"`
	User_id      string             `bson:"user_id" json:"user_id"`
	RefreshToken string             `bson:"refreshToken" json:"refreshToken"`
}
