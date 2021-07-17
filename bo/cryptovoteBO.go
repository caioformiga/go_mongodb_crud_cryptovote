package bo

import (
	"errors"
	"log"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/dao"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func validateCryptoVoteArguments(crypto model.CryptoVote) (bool, error) {
	validate := false

	if len(crypto.Name) > 0 {
		validate = true
	} else {
		validate = false
		return validate, errors.New("name não pode ser vazio")
	}

	if len(crypto.Symbol) > 0 {
		validate = true
	} else {
		validate = false
		return validate, errors.New("symbol não pode ser vazio")
	}

	if crypto.Qtd_Upvote >= 0 {
		validate = true
	} else {
		validate = false
		return validate, errors.New("qtd_upvote não pode ser menor do que zero")
	}

	if crypto.Qtd_Downvote >= 0 {
		validate = true
	} else {
		validate = false
		return validate, errors.New("qtd_upvote não pode ser menor do que zero")
	}
	return validate, nil
}

/*
	validateCryptoVoteUniqueData verifica se não existe no banco alguma CryptoVote que
	não atente aos critérios de unique abaixo
	name : unique
	symbol : unique
*/
func validateCryptoVoteUniqueData(crypto model.CryptoVote) (bool, error) {
	var validate bool = true
	var retrievedCryptoVotes []model.CryptoVote

	dao.SetCollectionName("cryptovotes")

	// usa a função criada no pacote dao
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		z := "Problemas no uso de GetMongoClientInstance: " + err.Error()
		log.Print(z)
	}

	// usa a função criada no pacote dao
	retrievedCryptoVotes, err = dao.FindManyCryptoVote(mongodbClient, bson.M{"name": crypto.Name})
	if err != nil {
		z := "Problemas no uso de FindManyCryptoVote: " + err.Error()
		log.Print(z)
	}

	if retrievedCryptoVotes != nil {
		validate = false
		return validate, errors.New("name informado já exite, escolha outro")
	}

	// usa a função criada no pacote dao
	retrievedCryptoVotes, err = dao.FindManyCryptoVote(mongodbClient, bson.M{"symbol": crypto.Symbol})
	if err != nil {
		z := "Problemas no uso de FindManyCryptoVote: " + err.Error()
		log.Print(z)
	}

	if retrievedCryptoVotes != nil {
		validate = false
		return validate, errors.New("symbol informado já exite, escolha outro")
	}
	return validate, err
}

/*
	ValidateCryptoVote recebe os campos e faz a validação
	name : não pode ser vazio, len(name) > 0
	symbol : não pode ser vazio, len(name) > 0
	qtd_upvote : não pode ser menor do que zero, qtd_upvote >= 0
	qtd_downvote : não pode ser menor do que zero, qtd_upvote >= 0
*/
func ValidateCryptoVote(crypto model.CryptoVote) (model.CryptoVote, error) {
	// usa a função criada no pacote bo
	_, err := validateCryptoVoteArguments(crypto)
	if err != nil {
		z := "Problemas na validação de dados da nova CryptoVote: " + err.Error()
		log.Print(z)
		return crypto, err
	} else {
		// usa a função criada no pacote bo
		_, err = validateCryptoVoteUniqueData(crypto)
		if err != nil {
			z := "Problemas na validação unique da nova CryptoVote: " + err.Error()
			log.Print(z)
			return crypto, err
		}
	}
	return crypto, err
}

/*
	CreateCryptoVote não faz a validação das entradas antes de criar uma model.CryptoVote no banco
	entrada
	validatedCryptoVote model.CryptoVote

	retorno
	uma model.CryptoVote armazenada no banco, testes realizados como o mongoDB
*/
func CreateCryptoVote(validatedCryptoVote model.CryptoVote) (model.CryptoVote, error) {

	dao.SetCollectionName("cryptovotes")

	// usa a função criada no pacote dao
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		z := "Problemas no uso de GetMongoClientInstance: " + err.Error()
		log.Print(z)
	}

	// usa a função criada no pacote dao
	insertResult, err := dao.CreateCryptoVote(mongodbClient, validatedCryptoVote)
	if err != nil || insertResult.InsertedID == nil {
		z := "Problemas na execução de dao.CreateCryptoVote: " + err.Error()
		log.Print(z)
	}

	// cria filtro com id para localizar dado
	filter := bson.M{"_id": insertResult.InsertedID}

	// usa a função criada no pacote dao
	savedCryptoVote, err := dao.FindOneCryptoVote(mongodbClient, filter)
	if err != nil {
		z := "Problemas na execução de dao.FindOneCryptoVote: " + err.Error()
		log.Print(z)
	}

	// retorna a nova CryptoCurrency salva o banco
	return savedCryptoVote, err
}

/*
	RetrieveAllCryptoVoteByFilter faz uma busca no banco para recuperar uma coleção de model.CryptoVote
	entrada
	filter := bson.M{"key": "value"}

	retorno
	uma coleção de model.CryptoVote armazenada no banco, testes realizados como o mongoDB
*/
func RetrieveAllCryptoVoteByFilter(filter bson.M) ([]model.CryptoVote, error) {
	dao.SetCollectionName("cryptovotes")

	// usa a função criada no pacote dao
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		z := "Problemas no uso de GetMongoClientInstance: " + err.Error()
		log.Print(z)
	}

	// usa a função criada no pacote dao
	retrievedCryptoVotes, err := dao.FindManyCryptoVote(mongodbClient, filter)
	if err != nil {
		log.Fatal(err)
	}
	return retrievedCryptoVotes, err
}

/*
	RetrieveOneCryptoVoteById faz uma busca no banco usando o ID para recuperar um único model.CryptoVote
	entrada
	id escrito como uma string

	retorno
	uma model.CryptoVote armazenada no banco, testes realizados como o mongoDB
*/
func RetrieveOneCryptoVoteById(id string) (model.CryptoVote, error) {
	// cria os parametros do filtro sem restrições
	primitiveObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		z := "Problemas no uso de primitive.ObjectIDFromHex: " + err.Error()
		log.Print(z)
	}
	filter := bson.M{"_id": primitiveObjectID}

	dao.SetCollectionName("cryptovotes")

	// usa a função criada no pacote dao
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		z := "Problemas no uso de GetMongoClientInstance: " + err.Error()
		log.Print(z)
	}

	// usa a função criada no pacote dao
	retrievedCryptoVote, err := dao.FindOneCryptoVote(mongodbClient, filter)
	if err != nil {
		z := "Problemas no uso de dao.FindOneCryptoVote: " + err.Error()
		log.Print(z)
	}
	return retrievedCryptoVote, err
}

/*
	UpdateAllCryptoVoteByFilter faz uma atualização de todas as model.CryptoVote que satisfazem o filtro
	entrada
	filter := bson.M{"key": "value"}

	retorno
	uma coleção de model.CryptoVote armazenada no banco, testes realizados como o mongoDB
*/
func UpdateAllCryptoVoteByFilter(filter bson.M, newData bson.M) ([]model.CryptoVote, error) {
	dao.SetCollectionName("cryptovotes")

	// usa a função criada no pacote dao
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		z := "Problemas no uso de GetMongoClientInstance: " + err.Error()
		log.Print(z)
	}

	// usa a função criada no pacote dao
	retrievedCryptoVotes, err := dao.FindManyCryptoVote(mongodbClient, filter)
	if err != nil {
		z := "Problemas no uso de dao.FindManyCryptoVote: " + err.Error()
		log.Print(z)
	}

	var updatedCryptoVotes []model.CryptoVote

	// para cada CryptoVote faz um update
	for _, retrievedCryptoVote := range retrievedCryptoVotes {
		// cria filtro com id para localizar dado
		idFilter := bson.M{"_id": retrievedCryptoVote.Id}

		savedCryptoVote, err := dao.UpdateOneCryptoVote(mongodbClient, idFilter, newData)
		if err != nil {
			z := "Problemas no uso de dao.UpdateOneCryptoVote: " + err.Error()
			log.Print(z)
		}
		updatedCryptoVotes = append(updatedCryptoVotes, savedCryptoVote)
	}
	return updatedCryptoVotes, err
}

/*
	DeleteAllCryptoVoteByFilter faz uma deleção de todas as model.CryptoVote que satisfazem o filtro
	entrada
	filter := bson.M{"key": "value"}

	retorno
	a quantidade de model.CryptoVote deletadas do banco, testes realizados como o mongoDB
*/
func DeleteAllCryptoVoteByFilter(filter bson.M) (int64, error) {
	dao.SetCollectionName("cryptovotes")

	// usa a função criada no pacote dao
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		z := "Problemas no uso de GetMongoClientInstance: " + err.Error()
		log.Print(z)
	}

	// usa a função criada no pacote dao
	deleteResult, err := dao.DeleteManyCryptoVote(mongodbClient, filter)
	if err != nil {
		z := "Problemas no uso de dao.DeleteManyCryptoVote: " + err.Error()
		log.Print(z)
	}
	return deleteResult.DeletedCount, err
}

func AddUpVote(id string) {
	// usa a função criada no arquivo cryptovoteBO.go pacote bo
	retrievedCryptoVote, _ := RetrieveOneCryptoVoteById(id)

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
	retrievedCryptoVote, _ := RetrieveOneCryptoVoteById(id)

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
