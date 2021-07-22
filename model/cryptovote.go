package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CryptoVote struct {
	Id           primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Name         string             `json:"name,omitempty" bson:"name"`
	Symbol       string             `json:"symbol,omitempty" bson:"symbol"`
	Qtd_Upvote   int                `json:"qtd_upvote,omitempty" bson:"qtd_upvote"`
	Qtd_Downvote int                `json:"qtd_downvote,omitempty" bson:"qtd_downvote"`
	Sum          int                `json:"sum,omitempty" bson:"sum"`
}

type FilterCryptoVote struct {
	Id     primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Name   string             `json:"name,omitempty" bson:"name"`
	Symbol string             `json:"symbol,omitempty" bson:"symbol"`
}

type PullCryptoVote struct {
	Id      primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Cryptos CryptoVote         `json:"crypto,omitempty" bson:"crypto"`
}
