package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CryptoVote struct {
	Id           primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Name         string             `json:"name,omitempty" bson:"name"`
	Symbol       string             `json:"symbol,omitempty" bson:"symbol"`
	Qtd_Upvote   int64              `json:"qtd_upvote,omitempty" bson:"qtd_upvote"`
	Qtd_Downvote int64              `json:"qtd_downvote,omitempty" bson:"qtd_downvote"`
	Sum          int64              `json:"sum,omitempty" bson:"sum"`
	SumAbsolute  int64              `json:"sum_absolute,omitempty" bson:"sum_absolute"`
}

type FilterCryptoVote struct {
	Id     primitive.ObjectID `json:"-" bson:"-"`
	Name   string             `json:"name,omitempty" bson:"name"`
	Symbol string             `json:"symbol,omitempty" bson:"symbol"`
}

type SumaryCryptoVote struct {
	Token       string `json:"token" bson:"token"`
	SumType     string `json:"sum_type" bson:"sum_type"`
	SumAbsolute int64  `json:"sum_absolute" bson:"sum_absolute"`
}
