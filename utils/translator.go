package utils

import (
	"encoding/json"
	"errors"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"

	"go.mongodb.org/mongo-driver/bson"
)

func MarshalCryptoVoteToBsonFilter(cryptoVote model.CryptoVote) (bson.M, error) {
	return marshalInterfaceToBsonFilter(cryptoVote)
}

func MarshalFilterCryptoVoteToBsonFilter(filterCryptoVote model.FilterCryptoVote) (bson.M, error) {
	return marshalInterfaceToBsonFilter(filterCryptoVote)
}

func marshalInterfaceToBsonFilter(i interface{}) (bson.M, error) {
	var err error
	var out bson.M

	// translate from interface i to json
	var jsonData []byte
	jsonData, err = json.Marshal(i)
	if err != nil {
		z := "[cryptovote.translatorBO] Problems to Marshal a json object: " + err.Error()
		err = errors.New(z)
		return out, err
	}

	// translate from json to bson
	var ptr *bson.M
	err = json.Unmarshal(jsonData, &ptr)
	if err != nil {
		z := "[cryptovote.json] Problems to Unmarshal from json to bson.M: " + err.Error()
		err = errors.New(z)
		return out, err
	}

	out = *ptr
	return out, err
}
