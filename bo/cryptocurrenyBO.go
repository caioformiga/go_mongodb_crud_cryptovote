package bo

import (
	"errors"
	"log"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/dao"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"
	"go.mongodb.org/mongo-driver/bson"
)

func validateCryptoCurrency(name string, symbol string) (bool, error) {
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

/*
	CreateCryptoCurrency faz a validação das entradas antes de criar uma model.CryptoCurrency
	entrada
	name deve ser uma string não nula, teste feito usando len(name) > 0
	symbol deve ser uma string não nula, teste feito usando len(symbol) > 0

	retorno
	uma model.CryptoCurrency armazenada no banco, testes realizados como o mongoDB
*/
func CreateCryptoCurrency(name string, symbol string) model.CryptoCurrency {
	// usa a função criada no pacote bo
	_, err := validateCryptoCurrency(name, symbol)
	if err != nil {
		log.Fatalf("Problemas na validação de dados da nova CryptoCurrency: %v", err)
	}

	dao.SetCollectionName("cryptocurrencies")

	// usa a função criada no pacote dao
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

/*
	RetrieveAllCryptoCurrencyByFilter faz uma busca no banco para recuperar uma coleção de model.CryptoCurrency
	entrada
	filter := bson.M{"key": "value"}

	retorno
	uma coleção de model.CryptoCurrency armazenada no banco, testes realizados como o mongoDB
*/
func RetrieveAllCryptoCurrencyByFilter(filter bson.M) []model.CryptoCurrency {
	dao.SetCollectionName("cryptocurrencies")

	// usa a função criada no pacote dao
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		log.Fatal(err)
	}

	retrievedCryptoCurrencies, err := dao.FindManyCryptoCurrency(mongodbClient, filter)
	if err != nil {
		log.Fatal(err)
	}
	return retrievedCryptoCurrencies
}

/*
	UpdateAllCryptoCurrencyByFilter faz uma atualização de todas as model.CryptoCurrency que satisfazem o filtro
	entrada
	filter := bson.M{"key": "value"}

	retorno
	uma coleção de model.CryptoCurrency armazenada no banco, testes realizados como o mongoDB
*/
func UpdateAllCryptoCurrencyByFilter(filter bson.M, newData bson.M) []model.CryptoCurrency {
	dao.SetCollectionName("cryptocurrencies")

	// usa a função criada no pacote dao
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

		savedCryptoCurrency, err := dao.UpdateOneCryptoCurrency(mongodbClient, idFilter, newData)
		if err != nil {
			log.Fatal(err)
		}
		updatedCryptoCurrencies = append(updatedCryptoCurrencies, savedCryptoCurrency)
	}
	return updatedCryptoCurrencies
}

/*
	DeleteCryptoCurrencyByFilter faz uma deleção de todas as model.CryptoCurrency que satisfazem o filtro
	entrada
	filter := bson.M{"key": "value"}

	retorno
	a quantidade de model.CryptoCurrency deletadas do banco, testes realizados como o mongoDB
*/
func DeleteAllCryptoCurrencyByFilter(filter bson.M) int64 {
	dao.SetCollectionName("cryptocurrencies")

	// usa a função criada no pacote dao
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
