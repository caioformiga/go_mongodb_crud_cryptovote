package bo

import (
	"errors"
	"log"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/dao"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"
	"go.mongodb.org/mongo-driver/bson"
)

func ValidateCryptoCurrency(name string, symbol string) (bool, error) {
	validate := false

	if len(name) > 0 {
		validate = true
	} else {
		validate = false
		return validate, errors.New("name não pode ser vazio")
	}

	if len(symbol) > 0 {
		validate = true
	} else {
		validate = false
		return validate, errors.New("symbol não pode ser vazio")
	}
	return validate, nil
}

// Cria no banco uma cryptocurrency e retorna
func CreateCryptoCurrency(name string, symbol string) model.CryptoCurrency {
	dao.SetCollectioName("cryptocurrencies")

	// usa a função criada no arquivo mongodb.go pacote main (raiz) do projeto
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		log.Fatal(err)
	}

	objCryptoCurrency := model.CryptoCurrency{
		Id:     [12]byte{},
		Name:   name,
		Symbol: symbol,
	}

	// usa a função criada no pacote dao
	insertResult, err := dao.CreateCryptoCurrency(mongodbClient, objCryptoCurrency)
	if err != nil || insertResult.InsertedID == nil {
		log.Fatal(err)
	}

	// cria filtro com id para localizar dado
	filter := bson.M{"_id": insertResult.InsertedID}

	// usa a função criada no pacote dao
	savedCryptoCurrency, err := dao.FindOneCryptoCurrency(mongodbClient, filter)
	if err != nil {
		log.Fatal(err)
	}

	// retorna a nova CryptoCurrency salva o banco
	return savedCryptoCurrency
}

// Remove todas as cryptocurrencies que atenderem ao filtro
func ReadCryptoCurrencyByFilter(filter bson.M) int64 {
	dao.SetCollectioName("cryptocurrencies")

	// usa a função criada no arquivo mongodb.go pacote main (raiz) do projeto
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		log.Fatal(err)
	}

	deleteResult, err := dao.DeleteManyCryptoCurrency(mongodbClient, filter)
	if err != nil {
		log.Fatal(err)
	}
	return deleteResult.DeletedCount
}

// Remove todas as cryptocurrencies que atenderem ao filtro
func UpdateCryptoCurrencyByFilter(filter bson.M, newData bson.M) []model.CryptoCurrency {
	dao.SetCollectioName("cryptocurrencies")

	// usa a função criada no arquivo mongodb.go pacote main (raiz) do projeto
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		log.Fatal(err)
	}

	// usa a função criada no pacote dao
	sliceCryptoCurrency, err := dao.FindManyCryptoCurrency(mongodbClient, filter)
	if err != nil {
		log.Fatal(err)
	}

	var updatedCryptoCurrencies []model.CryptoCurrency

	// para cada CryptoCurrency faz um update
	for _, objCryptoCurrency := range sliceCryptoCurrency {
		// cria filtro com id para localizar dado
		idFilter := bson.M{"_id": objCryptoCurrency.Id}

		updatedCryptoCurrency, err := dao.UpdateOneCryptoCurrency(mongodbClient, idFilter, newData)
		if err != nil {
			log.Fatal(err)
		}
		updatedCryptoCurrencies = append(updatedCryptoCurrencies, updatedCryptoCurrency)
	}
	return updatedCryptoCurrencies
}

// Remove todas as cryptocurrencies que atenderem ao filtro
func DeleteCryptoCurrencyByFilter(filter bson.M) int64 {
	dao.SetCollectioName("cryptocurrencies")

	// usa a função criada no arquivo mongodb.go pacote main (raiz) do projeto
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		log.Fatal(err)
	}

	deleteResult, err := dao.DeleteManyCryptoCurrency(mongodbClient, filter)
	if err != nil {
		log.Fatal(err)
	}
	return deleteResult.DeletedCount
}
