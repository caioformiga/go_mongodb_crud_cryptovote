package utils

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func LoadOneNewEmptyFilterCryptoVote() model.FilterCryptoVote {

	var filterCryptoVote model.FilterCryptoVote = model.FilterCryptoVote{
		Name:   "",
		Symbol: "",
	}

	return filterCryptoVote
}

func LoadOneNewFilterCryptoVoteFromArgs(name string, symbol string) model.FilterCryptoVote {

	var filterCryptoVote model.FilterCryptoVote = model.FilterCryptoVote{
		Name:   name,
		Symbol: symbol,
	}

	return filterCryptoVote
}

func LoadOneNewCryptoVoteDataFromArgs(name string, symbol string) model.CryptoVote {

	var cryptoVote model.CryptoVote = model.CryptoVote{
		Id:           primitive.NewObjectID(),
		Name:         name,
		Symbol:       symbol,
		Qtd_Upvote:   0,
		Qtd_Downvote: 0,
		Sum:          0,
		SumAbsolute:  0,
	}

	return cryptoVote
}

func LoadOneValidCryptoVoteDataFromModel(cryptoVote model.CryptoVote) model.CryptoVote {

	var validCryptoVote model.CryptoVote = model.CryptoVote{
		Id:           cryptoVote.Id,
		Name:         strings.Title(strings.ToLower(strings.TrimSpace(cryptoVote.Name))),
		Symbol:       strings.ToUpper(strings.TrimSpace(cryptoVote.Symbol)),
		Qtd_Upvote:   cryptoVote.Qtd_Upvote,
		Qtd_Downvote: cryptoVote.Qtd_Downvote,
		Sum:          cryptoVote.Sum,
		SumAbsolute:  cryptoVote.SumAbsolute,
	}

	if validCryptoVote.Id.IsZero() {
		validCryptoVote.Id = primitive.NewObjectID()
	}

	return validCryptoVote
}

func LoadManyCryptoVoteDataFromJson(jsonData []byte) ([]model.CryptoVote, error) {
	var ptr *[]model.CryptoVote
	err := json.Unmarshal(jsonData, &ptr)
	if err != nil {
		log.Printf("Something went wrong at json.Unmarshal for []model.CryptoVote: %v\n", err)
	}
	return *ptr, err
}

func LoadOneCryptoVoteDataFromJson(jsonData []byte) (model.CryptoVote, error) {
	var ptr *model.CryptoVote
	err := json.Unmarshal(jsonData, &ptr)
	if err != nil {
		log.Printf("Something went wrong at json.Unmarshal for model.CryptoVote: %v\n", err)
	}
	return *ptr, err
}

var JsonOutDataSorted = []byte(`[{
    "name": "Klever",
    "symbol": "KLV",
	"qtd_upvote": 30000,
	"qtd_downvote": 1,
	"sum": 29999, 
	"sum_absolute": 29999
}, {	
	"name": "Bitcoin",
    "symbol": "BTC",
	"qtd_upvote": 1000,
	"qtd_downvote": 1,
	"sum": 999,
	"sum_absolute": 999
}, {
	"name": "Ethereum",
    "symbol": "ETH",
	"qtd_upvote": 5,
	"qtd_downvote": 15,
	"sum": -10,
	"sum_absolute": 10
}]`)

var JsonOutData = []byte(`[{
    "name": "Bitcoin",
    "symbol": "BTC",
	"qtd_upvote": 1000,
	"qtd_downvote": 1,
	"sum": 999,
	"sum_absolute": 999
}, {
	"name": "Ethereum",
    "symbol": "ETH",
	"qtd_upvote": 5,
	"qtd_downvote": 15,
	"sum": -10,
	"sum_absolute": 10
}, {	
	"name": "Klever",
    "symbol": "KLV",
	"qtd_upvote": 30000,
	"qtd_downvote": 1,
	"sum": 29999, 
	"sum_absolute": 29999
}]`)

var JsonInData = []byte(`[{
	"name": "Bitcoin",
	"symbol": "btc",
	"qtd_upvote": 1000,
	"qtd_downvote": 1,
	"sum": 999,
	"sum_absolute": 999
}, {
	"name": "EthEreum",
	"symbol": "EtH",
	"qtd_upvote": 5,
	"qtd_downvote": 15,
	"sum": -10,
	"sum_absolute": 10
}, {	
	"name": "klever",
	"symbol": "KLV",
	"qtd_upvote": 30000,
	"qtd_downvote": 1,
	"sum": 29999,
	"sum_absolute": 29999
}]`)

var JsonBadData = []byte(`[{
	"name": bool,
	"symbol": "DOGE",
	"qtd_upvote": 0,
	"qtd_downvote": 0,
	"sum": 0
}, {
	"name": "DogeCoin",
	"symbol": bool,
	"qtd_upvote": 0,
	"qtd_downvote": 0,
	"sum": 0
}, {	
	"name": "DogeCoin",
	"symbol": "DOGE",
	"qtd_upvote": bool,
	"qtd_downvote": 0,
	"sum": 0
}, {	
	"name": "DogeCoin",
	"symbol": "DOGE",
	"qtd_upvote": 0,
	"qtd_downvote": bool,
	"sum": 0
}]`)
