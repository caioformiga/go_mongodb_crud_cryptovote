package bo

import (
	"fmt"
	"log"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/dao"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"

	"go.mongodb.org/mongo-driver/bson"
)

const MIN_VOTE int = 0

var objCryptocurrency model.Cryptocurrency
var objCryptoVote model.CryptoVote

var sliceCryptoVotes []model.CryptoVote

var sliceCryptocurrencyName = []string{"Bitcoin", "Ethereum", "Tether", "Binance Coin", "Litecoin", "Klever"}
var sliceCryptocurrencySymbol = []string{"BTC", "ETH", "USDT", "BNB", "LLTC", "KLV"}

func CarregarDados() {
	if len(sliceCryptocurrencyName) != len(sliceCryptocurrencySymbol) {
		fmt.Printf("Tamanhos diferentes entre sliceCryptocurrencyName e sliceCryptocurrencySymbol\n")
		fmt.Printf("Tamanho de sliceCryptocurrencyName = %v\n", len(sliceCryptocurrencyName))
		fmt.Printf("Tamanho de sliceCryptocurrencySymbol = %v\n", len(sliceCryptocurrencySymbol))
		log.Fatal()
	}

	// usa a função criada no arquivo mongodb.go pacote main (raiz) do projeto
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(sliceCryptocurrencyName); i++ {
		// usa a função criada no pacote model
		objCryptocurrency = model.Cryptocurrency{
			Name:   sliceCryptocurrencyName[i],
			Symbol: sliceCryptocurrencySymbol[i],
		}

		objCryptoVote = model.CryptoVote{
			Id:           [12]byte{},
			Crypto:       objCryptocurrency,
			Qtd_Upvote:   0,
			Qtd_Downvote: 0,
		}

		// usa a função criada no pacote dao
		insertResult, err := dao.CreateCryptoVote(mongodbClient, objCryptoVote)
		if err != nil || insertResult.InsertedID == nil {
			log.Fatal(err)
		}
	}
}

func LimparDados() {
	// usa a função criada no arquivo mongodb.go pacote main (raiz) do projeto
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		log.Fatal(err)
	}

	// cria filtro sem parametros para limpar todos os dados
	filter := bson.M{}

	// usa a função criada no pacote dao
	deleteResult, err := dao.DeleteManyCryptoVote(mongodbClient, filter)
	if err != nil {
		log.Fatal(err)
	}

	if deleteResult.DeletedCount >= 0 {
		if deleteResult.DeletedCount == 1 {
			fmt.Printf("Foi removido %d registro\n", deleteResult.DeletedCount)
		} else {
			fmt.Printf("Foram removidos %d registros\n", deleteResult.DeletedCount)
		}
	} else {
		fmt.Printf("NENHUM registro foi removido, pois não foi localizado entradas com o filtro %+v\n", filter)
	}
}

func FindByIdHex(idHex string) model.CryptoVote {
	// usa a função criada no arquivo mongodb.go pacote main (raiz) do projeto
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		log.Fatal(err)
	}

	// usa a função criada no pacote dao
	objCryptoVote, err := dao.FindOneCryptoVoteByIdHex(mongodbClient, idHex)
	if err != nil {
		log.Fatal(err)
	}
	return objCryptoVote
}

func AddUpVote(idHex string) {
	// usa a função criada no arquivo cryptovoteBO.go pacote bo
	objCryptoVote := FindByIdHex(idHex)

	// cria filtro com id para localizar dado
	filter := bson.M{"_id": objCryptoVote.Id}

	// soma valora atual de Qtd_Upvote +1
	qtdNova := objCryptoVote.Qtd_Upvote + 1

	newData := bson.M{
		"$set": bson.M{"qtd_upvote": qtdNova},
	}

	// usa a função criada no arquivo mongodb.go pacote main (raiz) do projeto
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		log.Fatal(err)
	}

	objCryptoVote, err = dao.UpdateOneCryptoVote(mongodbClient, filter, newData)
	if err != nil {
		log.Fatal(err)
	}
}

func AddDownVote(idHex string) {
	// usa a função criada no arquivo cryptovoteBO.go pacote bo
	objCryptoVote := FindByIdHex(idHex)

	// cria filtro com id para localizar dado
	filter := bson.M{"_id": objCryptoVote.Id}

	// soma valora atual de Qtd_Downvote +1
	qtdNova := objCryptoVote.Qtd_Downvote + 1

	newData := bson.M{
		"$set": bson.M{"qtd_downvote": qtdNova},
	}

	// usa a função criada no arquivo mongodb.go pacote main (raiz) do projeto
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		log.Fatal(err)
	}

	objCryptoVote, err = dao.UpdateOneCryptoVote(mongodbClient, filter, newData)
	if err != nil {
		log.Fatal(err)
	}
}
