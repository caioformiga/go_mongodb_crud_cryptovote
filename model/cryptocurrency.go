package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type CryptoCurrency struct {
	Id     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string             `bson:"name"`
	Symbol string             `bson:"symbol"`
	IdHex  string             `bson:"idHex,omitempty"`
}
