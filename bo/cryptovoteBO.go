package bo

import (
	"encoding/json"
	"errors"
	"log"
	"strings"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/dao"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"

	"go.mongodb.org/mongo-driver/bson"
)

/*
	CreateCryptoVote faz a validação das entradas antes de criar uma model.CryptoVote no banco
	entrada cryptoVote.Name

	cryptoVote.Name se "KLevER" salva "Klever"
	cryptoVote.Symbol se klv salva "KLV"

	retorno
	uma model.CryptoVote armazenada no banco, testes realizados como o mongoDB
*/
func CreateCryptoVote(cryptoVote model.CryptoVote) (model.CryptoVote, error) {

	// popular no padrão
	cryptoVote.Name = strings.Title(strings.ToLower(strings.TrimSpace(cryptoVote.Name)))
	cryptoVote.Symbol = strings.ToUpper(strings.TrimSpace(cryptoVote.Symbol))

	dao.SetCollectionName("cryptovotes")

	// usa a função criada no pacote dao
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		z := "Problemas no uso de GetMongoClientInstance: " + err.Error()
		log.Print(z)
	}

	// usa a função criada no pacote bo
	validate, err := ValidateCryptoVote(cryptoVote)

	if !validate {
		return cryptoVote, err
	} else {

		// usa a função criada no pacote dao
		insertResult, err := dao.CreateCryptoVote(mongodbClient, cryptoVote)
		if err != nil || insertResult.InsertedID == nil {
			z := "Problemas na execução de dao.CreateCryptoVote: " + err.Error()
			log.Print(z)
		}

		// cria filtro com id para localizar dado
		filter := bson.M{"_id": insertResult.InsertedID}

		// usa a função criada no pacote dao
		cryptoVote, err = dao.FindOneCryptoVote(mongodbClient, filter)
		if err != nil {
			z := "Problemas na execução de dao.FindOneCryptoVote: " + err.Error()
			log.Print(z)
		}
	}
	// retorna a nova CryptoCurrency salva no banco
	return cryptoVote, err
}

/*
	RetrieveAllCryptoVoteByFilter faz uma busca no banco para recuperar uma coleção de model.CryptoVote
	entrada
	filter := bson.M{"key": "value"}

	retorno
	uma coleção de model.CryptoVote armazenada no banco, testes realizados como o mongoDB
*/
func RetrieveAllCryptoVoteByFilter(filterCryptoVote model.CryptoVote) ([]model.CryptoVote, error) {
	var retrievedCryptoVotes []model.CryptoVote

	dao.SetCollectionName("cryptovotes")

	// usa a função criada no pacote dao
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		z := "Problemas no uso de GetMongoClientInstance: " + err.Error()
		log.Print(z)
		return retrievedCryptoVotes, err
	}

	// convertendo de CryptoVote para json
	jsonData, err := json.Marshal(filterCryptoVote)
	if err != nil {
		log.Printf("Problemas para fazer Marshal de CryptoVote para bson.M: %v", err)
		return retrievedCryptoVotes, err
	}

	// convertendo de json para bson
	var filter = bson.M{}
	err = json.Unmarshal(jsonData, &filter)
	if err != nil {
		log.Printf("Problemas para fazer Unmarshal de json para bson.M: %v", err)
		return retrievedCryptoVotes, err
	}

	// usa a função criada no pacote dao
	retrievedCryptoVotes, err = dao.FindManyCryptoVote(mongodbClient, filter)
	if err != nil {
		z := "Problemas no uso de FindManyCryptoVote: " + err.Error()
		log.Print(z)
		return retrievedCryptoVotes, err
	}
	return retrievedCryptoVotes, err
}

/*
	RetrieveOneCryptoVoteById faz uma busca para recuperar um único model.CryptoVote
	entrada
	para fazer a buscca pelos menos um arg (name ou symbol) precisa ser diferente de null

	retorno
	uma model.CryptoVote armazenada no banco, testes realizados como o mongoDB
*/
func RetrieveOneCryptoVote(name string, symbol string) (model.CryptoVote, error) {
	var retrievedCryptoVote model.CryptoVote

	// popular os args no padrão para criar o filtro
	var filterCryptoVote model.CryptoVote = model.CryptoVote{
		Name:   strings.Title(strings.ToLower(strings.TrimSpace(name))),
		Symbol: strings.ToUpper(strings.TrimSpace(symbol)),
	}

	// se pelo menos um dos filtros não for empty pode seguir com a busca
	var validate bool = false
	var err error
	if validateCryptoVoteArgumentNotEmpty(filterCryptoVote.Name) || validateCryptoVoteArgumentNotEmpty(filterCryptoVote.Symbol) {
		validate = true
	} else {
		err = errors.New("[cryptovote.validationBO] um dos filtros deve ser not empty")
	}

	// se pelo menos um dos filtros não for empty faz a busca
	if validate {
		// convertendo de []model.CryptoVote para json apenas os campos que não estão nulos
		jsonData, err := json.Marshal(filterCryptoVote)
		if err != nil {
			log.Printf("Problemas para recuperar dados: %v", err)
			return retrievedCryptoVote, err
		}

		filter := bson.M{}
		err = json.Unmarshal(jsonData, &filter)
		if err != nil {
			log.Printf("Problemas para fazer Unmarshal de json para bson.M: %v", err)
			return retrievedCryptoVote, err
		}

		dao.SetCollectionName("cryptovotes")

		// usa a função criada no pacote dao
		mongodbClient, err := dao.GetMongoClientInstance()
		if err != nil {
			z := "Problemas no uso de GetMongoClientInstance: " + err.Error()
			log.Print(z)
			return retrievedCryptoVote, err
		}

		// usa a função criada no pacote dao
		retrievedCryptoVote, err = dao.FindOneCryptoVote(mongodbClient, filter)
		if err != nil {
			z := "Problemas no uso de dao.FindOneCryptoVote: " + err.Error()
			log.Print(z)
			return retrievedCryptoVote, err
		}
	}
	return retrievedCryptoVote, err
}

/*
	UpdateOneCryptoVoteByFilter faz uma atualização de todas as model.CryptoVote que satisfazem o filtro

	filterCryptoVote = model.CryptoVote{
		Name:   "Bitcoin",
		Symbol: "",
	}

	newCryptoData := model.CryptoVote{
		Name:         "FormiCOIN",
		Symbol:       "FORMFORMFORMFORM",
		Qtd_Upvote:   0,
		Qtd_Downvote: 0,
	}

	retorno
	uma coleção de model.CryptoVote armazenada no banco, testes realizados como o mongoDB
*/
func UpdateOneCryptoVoteByFilter(filterCryptoVote model.CryptoVote, cryptoNewData model.CryptoVote) (model.CryptoVote, error) {
	var retrievedCryptoVote model.CryptoVote = model.CryptoVote{}

	dao.SetCollectionName("cryptovotes")

	// usa a função criada no pacote dao
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		z := "Problemas no uso de GetMongoClientInstance: " + err.Error()
		log.Print(z)
		return retrievedCryptoVote, err
	}

	cryptoNewData.Name = strings.Title(strings.ToLower(strings.TrimSpace(cryptoNewData.Name)))
	cryptoNewData.Symbol = strings.ToUpper(strings.TrimSpace(cryptoNewData.Symbol))

	// usa a função criada no pacote bo
	validate, err := ValidateCryptoVote(cryptoNewData)

	if validate {
		filterCryptoVote.Name = strings.Title(strings.ToLower(strings.TrimSpace(filterCryptoVote.Name)))
		filterCryptoVote.Symbol = strings.ToUpper(strings.TrimSpace(filterCryptoVote.Symbol))

		retrievedCryptoVote, err = RetrieveOneCryptoVote(filterCryptoVote.Name, filterCryptoVote.Symbol)
		if err != nil {
			z := "Problemas no uso de dao.FindOneCryptoVote: " + err.Error()
			log.Print(z)
			return retrievedCryptoVote, err
		}

		if !retrievedCryptoVote.Id.IsZero() {
			// caso os novos dados respeitem as regras de validateCryptoVoteUniqueData
			// realiza a atualização
			idFilter := bson.M{"_id": retrievedCryptoVote.Id}

			// convertendo de model.CryptoVote para json apenas os campos que não estão nulos
			jsonData, err := json.Marshal(cryptoNewData)
			if err != nil {
				z := "Problemas no Marshal: " + err.Error()
				log.Print(z)
				return retrievedCryptoVote, err
			}

			// convertendo de json para bson
			bsonCryptoNewData := bson.M{}
			err = json.Unmarshal(jsonData, &bsonCryptoNewData)
			if err != nil {
				z := "Problemas no Unmarshal: " + err.Error()
				log.Print(z)
				return retrievedCryptoVote, err
			}

			newData := bson.M{
				"$set": bsonCryptoNewData,
			}

			// atualização
			retrievedCryptoVote, err = dao.UpdateOneCryptoVote(mongodbClient, idFilter, newData)
			if err != nil {
				z := "Problemas no uso de dao.UpdateOneCryptoVote: " + err.Error()
				log.Print(z)
				return retrievedCryptoVote, err
			}
		}
	}
	return retrievedCryptoVote, err
}

/*
	DeleteAllCryptoVoteByFilter faz uma deleção de todas as model.CryptoVote que satisfazem o filtro

	filterCryptoVote = model.CryptoVote{
		Name:   "Bitcoin",
		Symbol: "",
	}

	retorno
	a quantidade de model.CryptoVote deletadas do banco, testes realizados como o mongoDB
*/
func DeleteAllCryptoVoteByFilter(filterCryptoVote model.CryptoVote) (int64, error) {
	var validate bool = false
	var err error

	// popular os args no padrão para criar o filtro
	filterCryptoVote.Name = strings.Title(strings.ToLower(strings.TrimSpace(filterCryptoVote.Name)))
	filterCryptoVote.Symbol = strings.ToUpper(strings.TrimSpace(filterCryptoVote.Symbol))

	// se pelo menos um dos filtros não for empty pode seguir com a busca
	if validateCryptoVoteArgumentNotEmpty(filterCryptoVote.Name) || validateCryptoVoteArgumentNotEmpty(filterCryptoVote.Symbol) {
		validate = true
	}

	if !validate {
		z := "Problemas na validateCryptoVoteArgumentNotEmpty: " + err.Error()
		log.Print(z)
		err = errors.New("[cryptovote.validationBO] um dos filtros deve ser not empty")
		return 0, err
	}

	dao.SetCollectionName("cryptovotes")

	// usa a função criada no pacote dao
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		z := "Problemas no uso de GetMongoClientInstance: " + err.Error()
		log.Print(z)
		return 0, err
	}

	// convertendo de CryptoVote para json
	jsonData, err := json.Marshal(filterCryptoVote)
	if err != nil {
		log.Printf("Problemas para fazer Marshal de CryptoVote para bson.M: %v", err)
		return 0, err
	}

	// convertendo de json para bson
	var filter = bson.M{}
	err = json.Unmarshal(jsonData, &filter)
	if err != nil {
		log.Printf("Problemas para fazer Unmarshal de json para bson.M: %v", err)
		return 0, err
	}

	// usa a função criada no pacote dao
	deleteResult, err := dao.DeleteManyCryptoVote(mongodbClient, filter)
	if err != nil {
		z := "Problemas no uso de dao.DeleteManyCryptoVote: " + err.Error()
		log.Print(z)
		return 0, err
	}
	return deleteResult.DeletedCount, err
}

func DeleteAllCryptoVote() (int64, error) {
	dao.SetCollectionName("cryptovotes")

	// usa a função criada no pacote dao
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		z := "Problemas no uso de GetMongoClientInstance: " + err.Error()
		log.Print(z)
		return 0, err
	}

	// convertendo de json para bson
	var filter = bson.M{}

	// usa a função criada no pacote dao
	deleteResult, err := dao.DeleteManyCryptoVote(mongodbClient, filter)
	if err != nil {
		z := "Problemas no uso de dao.DeleteManyCryptoVote: " + err.Error()
		log.Print(z)
		return 0, err
	}
	return deleteResult.DeletedCount, err
}

/*
	AddUpVote faz uma adição de um voto em UpVote de uma CryptoVote que satisfazem os argumentos
	de entrada usando a função RetrieveOneCryptoVote(name,symbol) para localizar uma CryptoVote.
	Lembrando que cada CryptoVote não pode ter name ou symbol repetidos.

	retorno
	nil se não tiver problema ou erro caso contrário
*/
func AddUpVote(name string, symbol string) (model.CryptoVote, error) {
	// usa a função criada no arquivo cryptovoteBO.go pacote bo
	cryptoVote, err := RetrieveOneCryptoVote(name, symbol)
	if err != nil {
		z := "Problemas no uso de RetrieveOneCryptoVote: " + err.Error()
		log.Print(z)
	}

	if !cryptoVote.Id.IsZero() {
		// soma valor atual de Qtd_Upvote +1
		newQtd := cryptoVote.Qtd_Upvote + 1

		typeVote := "qtd_upvote"

		cryptoVote, err = updateVote(cryptoVote, typeVote, newQtd)
	}
	return cryptoVote, err
}

/*
	AddDownVote faz uma adição de um voto em DownVote de uma CryptoVote que satisfazem os argumentos
	de entrada usando a função RetrieveOneCryptoVote(name,symbol) para localizar uma CryptoVote.
	Lembrando que cada CryptoVote não pode ter name ou symbol repetidos.

	retorno
	nil se não tiver problema ou erro caso contrário
*/
func AddDownVote(name string, symbol string) (model.CryptoVote, error) {
	// usa a função criada no arquivo cryptovoteBO.go pacote bo
	cryptoVote, err := RetrieveOneCryptoVote(name, symbol)
	if err != nil {
		z := "Problemas no uso de UpdateAllCryptoVoteByFilter: " + err.Error()
		log.Print(z)
	}

	if !cryptoVote.Id.IsZero() {
		// soma valor atual de Qtd_Downvote +1
		newQtd := cryptoVote.Qtd_Downvote + 1

		typeVote := "qtd_downvote"

		cryptoVote, err = updateVote(cryptoVote, typeVote, newQtd)
	}
	return cryptoVote, err
}

func updateVote(retrievedCryptoVote model.CryptoVote, typeVote string, newQtd int) (model.CryptoVote, error) {
	dao.SetCollectionName("cryptovotes")

	// usa a função criada no pacote dao
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		z := "Problemas no uso de GetMongoClientInstance: " + err.Error()
		log.Print(z)
	}

	// cria filtro com id para localizar dado
	filter := bson.M{"_id": retrievedCryptoVote.Id}

	newData := bson.M{
		"$set": bson.M{typeVote: newQtd},
	}

	retrievedCryptoVote, err = dao.UpdateOneCryptoVote(mongodbClient, filter, newData)
	if err != nil {
		z := "Problemas no uso de dao.UpdateOneCryptoVote: " + err.Error()
		log.Print(z)
	}

	return retrievedCryptoVote, err
}
