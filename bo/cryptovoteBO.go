package bo

import (
	"errors"
	"log"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/dao"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func validateCryptoVote(crypto model.CryptoCurrency, qtd_upvote int, qtd_downvote int) (bool, error) {
	validate := false

	if qtd_upvote >= 0 {
		validate = true
	} else {
		validate = false
		return validate, errors.New("qtd_upvote não pode ser menor do que zero")
	}

	if qtd_downvote >= 0 {
		validate = true
	} else {
		validate = false
		return validate, errors.New("qtd_upvote não pode ser menor do que zero")
	}

	validate, err := validateCryptoCurrency(crypto.Name, crypto.Symbol)
	if err != nil {
		log.Fatal(err)
	}
	return validate, nil
}

/*
	CreateCryptoVote faz a validação das entradas antes de criar uma model.CryptoVote
	entrada
	qtd_upvote deve ser >= 0
	qtd_downvote deve ser >= 0

	retorno
	uma model.CryptoVote armazenada no banco, testes realizados como o mongoDB
*/
func CreateCryptoVote(crypto model.CryptoCurrency, qtd_upvote int, qtd_downvote int) model.CryptoVote {
	// usa a função criada no pacote bo
	_, err := validateCryptoVote(crypto, qtd_upvote, qtd_downvote)
	if err != nil {
		log.Fatalf("Problemas na validação de dados da nova CryptoCurrency: %v", err)
	}

	dao.SetCollectionName("cryptovotes")

	// usa a função criada no pacote dao
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		log.Fatal(err)
	}

	retrievedCryptoVote := model.CryptoVote{
		Id:           [12]byte{},
		Crypto:       crypto,
		Qtd_Upvote:   qtd_upvote,
		Qtd_Downvote: qtd_downvote,
	}

	// usa a função criada no pacote dao
	insertResult, err := dao.CreateCryptoVote(mongodbClient, retrievedCryptoVote)
	if err != nil || insertResult.InsertedID == nil {
		log.Fatal(err)
	}

	// cria filtro com id para localizar dado
	filter := bson.M{"_id": insertResult.InsertedID}

	// usa a função criada no pacote dao
	savedCryptoVote, err := dao.FindOneCryptoVote(mongodbClient, filter)
	if err != nil {
		log.Fatal(err)
	}

	// retorna a nova CryptoCurrency salva o banco
	return savedCryptoVote
}

/*
	RetrieveAllCryptoVoteByFilter faz uma busca no banco para recuperar uma coleção de model.CryptoVote
	entrada
	filter := bson.M{"key": "value"}

	retorno
	uma coleção de model.CryptoVote armazenada no banco, testes realizados como o mongoDB
*/
func RetrieveAllCryptoVoteByFilter(filter bson.M) []model.CryptoVote {
	dao.SetCollectionName("cryptovotes")

	// usa a função criada no pacote dao
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		log.Fatal(err)
	}

	// usa a função criada no pacote dao
	retrievedCryptoVotes, err := dao.FindManyCryptoVote(mongodbClient, filter)
	if err != nil {
		log.Fatal(err)
	}
	return retrievedCryptoVotes
}

/*
	RetrieveOneCryptoVoteById faz uma busca no banco usando o ID para recuperar um único model.CryptoVote
	entrada
	id escrito como uma string

	retorno
	uma model.CryptoVote armazenada no banco, testes realizados como o mongoDB
*/
func RetrieveOneCryptoVoteById(id string) model.CryptoVote {
	// cria os parametros do filtro sem restrições
	primitiveObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": primitiveObjectID}

	dao.SetCollectionName("cryptovotes")

	// usa a função criada no pacote dao
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		log.Fatal(err)
	}

	// usa a função criada no pacote dao
	retrievedCryptoVote, err := dao.FindOneCryptoVote(mongodbClient, filter)
	if err != nil {
		log.Fatal(err)
	}
	return retrievedCryptoVote
}

/*
	UpdateAllCryptoVoteByFilter faz uma atualização de todas as model.CryptoVote que satisfazem o filtro
	entrada
	filter := bson.M{"key": "value"}

	retorno
	uma coleção de model.CryptoVote armazenada no banco, testes realizados como o mongoDB
*/
func UpdateAllCryptoVoteByFilter(filter bson.M, newData bson.M) []model.CryptoVote {
	dao.SetCollectionName("cryptocurrencies")

	// usa a função criada no pacote dao
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		log.Fatal(err)
	}

	// usa a função criada no pacote dao
	retrievedCryptoVotes, err := dao.FindManyCryptoVote(mongodbClient, filter)
	if err != nil {
		log.Fatal(err)
	}

	var updatedCryptoVotes []model.CryptoVote

	// para cada CryptoVote faz um update
	for _, retrievedCryptoVote := range retrievedCryptoVotes {
		// cria filtro com id para localizar dado
		idFilter := bson.M{"_id": retrievedCryptoVote.Id}

		savedCryptoVote, err := dao.UpdateOneCryptoVote(mongodbClient, idFilter, newData)
		if err != nil {
			log.Fatal(err)
		}
		updatedCryptoVotes = append(updatedCryptoVotes, savedCryptoVote)
	}
	return updatedCryptoVotes
}

/*
	DeleteAllCryptoVoteByFilter faz uma deleção de todas as model.CryptoVote que satisfazem o filtro
	entrada
	filter := bson.M{"key": "value"}

	retorno
	a quantidade de model.CryptoVote deletadas do banco, testes realizados como o mongoDB
*/
func DeleteAllCryptoVoteByFilter(filter bson.M) int64 {
	// usa a função criada no pacote dao
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		log.Fatal(err)
	}

	// usa a função criada no pacote dao
	deleteResult, err := dao.DeleteManyCryptoVote(mongodbClient, filter)
	if err != nil {
		log.Fatal(err)
	}
	return deleteResult.DeletedCount
}

func AddUpVote(id string) {
	// usa a função criada no arquivo cryptovoteBO.go pacote bo
	retrievedCryptoVote := RetrieveOneCryptoVoteById(id)

	// cria filtro com id para localizar dado
	filter := bson.M{"_id": retrievedCryptoVote.Id}

	// soma valora atual de Qtd_Upvote +1
	qtdNova := retrievedCryptoVote.Qtd_Upvote + 1

	newData := bson.M{
		"$set": bson.M{"qtd_upvote": qtdNova},
	}

	// usa a função criada no pacote dao
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		log.Fatal(err)
	}

	retrievedCryptoVote, err = dao.UpdateOneCryptoVote(mongodbClient, filter, newData)
	if err != nil {
		log.Fatal(err)
	}
}

func AddDownVote(id string) {
	// usa a função criada no arquivo cryptovoteBO.go pacote bo
	retrievedCryptoVote := RetrieveOneCryptoVoteById(id)

	// cria filtro com id para localizar dado
	filter := bson.M{"_id": retrievedCryptoVote.Id}

	// soma valora atual de Qtd_Downvote +1
	qtdNova := retrievedCryptoVote.Qtd_Downvote + 1

	newData := bson.M{
		"$set": bson.M{"qtd_downvote": qtdNova},
	}

	// usa a função criada no pacote dao
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		log.Fatal(err)
	}

	retrievedCryptoVote, err = dao.UpdateOneCryptoVote(mongodbClient, filter, newData)
	if err != nil {
		log.Fatal(err)
	}
}
