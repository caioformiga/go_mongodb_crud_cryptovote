package utils

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"

	"go.mongodb.org/mongo-driver/bson"
)

func MarshalCryptoVoteToBsonFilter(cryptoVote model.CryptoVote) (bson.M, error) {
	var err error
	var out bson.M
	var jsonData []byte

	// translate from cryptoVote to json
	jsonData, err = json.Marshal(cryptoVote)
	stringData := string(jsonData)
	if err != nil {
		z := "[utils.translator] Problems to Marshal a json object: " + err.Error()
		err = errors.New(z)
		return out, err
	}
	// translate from json to bson
	var ptr *bson.M
	decoder := json.NewDecoder(strings.NewReader(stringData))
	// define a json.Number type instead of float64 for int
	decoder.UseNumber()
	err = decoder.Decode(&ptr)
	if err != nil {
		z := "[utils.translator] Problems to Unmarshal from json to bson.M: " + err.Error()
		err = errors.New(z)
		return out, err
	}
	out = *ptr

	// perform cast of json.Number type to model defined type
	// details at:  [https://eager.io/blog/go-and-json/]
	jsonQtd_upvote, _ := out["qtd_upvote"].(json.Number).Int64()
	out["qtd_upvote"] = jsonQtd_upvote

	jsonQtd_downvote, _ := out["qtd_downvote"].(json.Number).Int64()
	out["qtd_downvote"] = jsonQtd_downvote

	jsonSum, _ := out["sum"].(json.Number).Int64()
	out["sum"] = jsonSum

	jsonSum_absolute, _ := out["sum_absolute"].(json.Number).Int64()
	out["sum_absolute"] = jsonSum_absolute
	return out, err
}

func MarshalFilterCryptoVoteToBsonFilter(filterCryptoVote model.FilterCryptoVote) (bson.M, error) {
	var err error
	var out bson.M
	var jsonData []byte

	// translate from filterCryptoVote to json
	jsonData, err = json.Marshal(filterCryptoVote)
	stringData := string(jsonData)
	if err != nil {
		z := "[utils.translator] Problems to Marshal a json object: " + err.Error()
		err = errors.New(z)
		return out, err
	}
	// translate from json to bson
	var ptr *bson.M
	decoder := json.NewDecoder(strings.NewReader(stringData))
	// define a json.Number type instead of float64 for int
	decoder.UseNumber()
	err = decoder.Decode(&ptr)
	if err != nil {
		z := "[utils.translator] Problems to Unmarshal from json to bson.M: " + err.Error()
		err = errors.New(z)
		return out, err
	}
	out = *ptr
	return out, err
}
