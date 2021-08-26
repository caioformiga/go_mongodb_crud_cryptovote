package test

import (
	"testing"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/utils"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestMarshalCryptoVoteToBsonFilterFullData(t *testing.T) {
	var out bson.M
	var err error
	var in model.CryptoVote = model.CryptoVote{
		Name:         "Bitcoin",
		Symbol:       "BTC",
		Qtd_Upvote:   20,
		Qtd_Downvote: 10,
		Sum:          10,
		SumAbsolute:  10,
	}
	out, err = utils.MarshalCryptoVoteToBsonFilter(in)
	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, in.Name, out["name"], "name should be equal")
	assert.Equal(t, in.Symbol, out["symbol"], "symbol should be equal")
	assert.Equal(t, in.Qtd_Upvote, out["qtd_upvote"], "qtd_upvote should be equal")
	assert.Equal(t, in.Qtd_Downvote, out["qtd_downvote"], "qtd_downvote should be equal")
	assert.Equal(t, in.Sum, out["sum"], "sum should be equal")
	assert.Equal(t, in.SumAbsolute, out["sum_absolute"], "sum_absolute should be equal")
}

func TestMarshalCryptoVoteToBsonFilterNameOnly(t *testing.T) {
	var out bson.M
	var err error
	var in model.CryptoVote = model.CryptoVote{
		Name: "Bitcoin",
	}
	out, err = utils.MarshalCryptoVoteToBsonFilter(in)
	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, in.Name, out["name"], "name should be equal")
}

func TestMarshalCryptoVoteToBsonFilterSymbolOnly(t *testing.T) {
	var out bson.M
	var err error
	var in model.CryptoVote = model.CryptoVote{
		Symbol: "BTC",
	}
	out, err = utils.MarshalCryptoVoteToBsonFilter(in)
	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, in.Symbol, out["symbol"], "symbol should be equal")
}

func TestMarshalFilterCryptoVoteToBsonFilterFullData(t *testing.T) {
	var out bson.M
	var err error
	var in model.FilterCryptoVote = model.FilterCryptoVote{
		Name:   "Bitcoin",
		Symbol: "BTC",
	}
	out, err = utils.MarshalFilterCryptoVoteToBsonFilter(in)
	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, in.Name, out["name"], "name should be equal")
	assert.Equal(t, in.Symbol, out["symbol"], "symbol should be equal")
}

func TestMarshalFilterCryptoVoteToBsonFilterNameOnly(t *testing.T) {
	var out bson.M
	var err error
	var in model.FilterCryptoVote = model.FilterCryptoVote{
		Name: "Bitcoin",
	}
	out, err = utils.MarshalFilterCryptoVoteToBsonFilter(in)
	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, in.Name, out["name"], "name should be equal")
}

func TestMarshalFilterCryptoVoteToBsonFilterSymbolOnly(t *testing.T) {
	var out bson.M
	var err error
	var in model.FilterCryptoVote = model.FilterCryptoVote{
		Symbol: "BTC",
	}
	out, err = utils.MarshalFilterCryptoVoteToBsonFilter(in)
	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, in.Symbol, out["symbol"], "symbol should be equal")
}
