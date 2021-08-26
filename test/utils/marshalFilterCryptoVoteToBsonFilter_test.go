package test

import (
	"testing"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/utils"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

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
