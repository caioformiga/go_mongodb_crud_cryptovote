package utils

import (
	"encoding/json"
	"log"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"
)

func Load_data(jsonData []byte) ([]model.CryptoVote, error) {
	var ptr *[]model.CryptoVote
	err := json.Unmarshal(jsonData, &ptr)
	if err != nil {
		log.Printf("Erro ao fazer json.Unmarshal dos dados de CryptoVotes para model.CryptoVote: %v\n", err)
	}
	return *ptr, err
}

var JsonOutData = []byte(`[{
    "name": "Bitcoin",
    "symbol": "BTC",
	"qtd_upvote": 0,
	"qtd_downvote": 0
}, {
	"name": "Ethereum",
    "symbol": "ETH",
	"qtd_upvote": 0,
	"qtd_downvote": 0
}, {	
	"name": "Klever",
    "symbol": "KLV",
	"qtd_upvote": 0,
	"qtd_downvote": 0
}]`)

var JsonInData = []byte(`[{
	"name": "Bitcoin",
	"symbol": "btc",
	"qtd_upvote": 0,
	"qtd_downvote": 0
}, {
	"name": "EthEreum",
	"symbol": "EtH",
	"qtd_upvote": 0,
	"qtd_downvote": 0
}, {	
	"name": "klever",
	"symbol": "KLV",
	"qtd_upvote": 0,
	"qtd_downvote": 0
}]`)

var JsonBadData = []byte(`[{
	"name": bool,
	"symbol": "DOGE",
	"qtd_upvote": 0,
	"qtd_downvote": 0
}, {
	"name": "DogeCoin",
	"symbol": bool,
	"qtd_upvote": 0,
	"qtd_downvote": 0
}, {	
	"name": "DogeCoin",
	"symbol": "DOGE",
	"qtd_upvote": bool,
	"qtd_downvote": 0
}, {	
	"name": "DogeCoin",
	"symbol": "DOGE",
	"qtd_upvote": 0,
	"qtd_downvote": bool
}]`)
