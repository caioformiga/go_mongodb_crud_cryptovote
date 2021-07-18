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
}

func NewCryptoVote(name string, symbol string, qtd_upvote int, qtd_downvote int) CryptoVote {
	var newCryptoVote = CryptoVote{
		Name:         name,
		Symbol:       symbol,
		Qtd_Upvote:   qtd_upvote,
		Qtd_Downvote: qtd_downvote,
	}
	return newCryptoVote
}
