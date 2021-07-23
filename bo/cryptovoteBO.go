package bo

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/dao"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/utils"

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
		z := "[cryptovote.mongodb] Problemas no uso de GetMongoClientInstance: " + err.Error()
		err = errors.New(z)
		return cryptoVote, err
	}

	// usa a função criada no pacote bo
	validate, err := ValidateCryptoVote(cryptoVote)
	if !validate || err != nil {
		return cryptoVote, err
	} else {

		// usa a função criada no pacote dao
		insertResult, err := dao.CreateCryptoVote(mongodbClient, cryptoVote)
		if err != nil || insertResult.InsertedID == nil {
			z := "[cryptovote.mongodb] Problemas na execução de dao.CreateCryptoVote: " + err.Error()
			err = errors.New(z)
			return cryptoVote, err
		}

		// cria filtro com id para localizar dado
		filter := bson.M{"_id": insertResult.InsertedID}

		// usa a função criada no pacote dao
		cryptoVote, err = dao.FindOneCryptoVote(mongodbClient, filter)
		if err != nil {
			z := "[cryptovote.mongodb] Problemas na execução de dao.FindOneCryptoVote: " + err.Error()
			err = errors.New(z)
			return cryptoVote, err
		}
	}
	// retorna a nova CryptoCurrency salva no banco
	return cryptoVote, err
}

/*
	RetrieveAllCryptoVoteByFilter faz uma busca no banco para recuperar uma coleção de model.CryptoVote
	entrada
	entrada
	// filtro vazio para recuperar todos os dados do banco
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   "",
		Symbol: "",
	}

	var filterCryptoVote = model.FilterCryptoVote{
		Name:   "Klever",
		Symbol: "",
	}

	var filterCryptoVote = model.FilterCryptoVote{
		Name:   "",
		Symbol: "KLV",
	}

	retorno
	uma coleção de model.CryptoVote armazenada no banco, testes realizados como o mongoDB
*/
func RetrieveAllCryptoVoteByFilter(filterCryptoVote model.FilterCryptoVote) ([]model.CryptoVote, error) {
	var retrievedCryptoVotes []model.CryptoVote

	dao.SetCollectionName("cryptovotes")

	// usa a função criada no pacote dao
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		z := "[cryptovote.mongodb] Problemas no uso de GetMongoClientInstance: " + err.Error()
		err = errors.New(z)
		return retrievedCryptoVotes, err
	}

	// convertendo de CryptoVote para json
	jsonData, err := json.Marshal(filterCryptoVote)
	if err != nil {
		z := "[cryptovote.json] Problemas para fazer Marshal de CryptoVote para json: " + err.Error()
		err = errors.New(z)
		return retrievedCryptoVotes, err
	}

	// convertendo de json para bson
	var filter = bson.M{}
	err = json.Unmarshal(jsonData, &filter)
	if err != nil {
		z := "[cryptovote.json] Problemas para fazer Unmarshal de json para bson.M: " + err.Error()
		err = errors.New(z)
		return retrievedCryptoVotes, err
	}

	// usa a função criada no pacote dao
	retrievedCryptoVotes, err = dao.FindManyCryptoVote(mongodbClient, filter)
	if err != nil {
		z := "[cryptovote.mongodb] Problemas no uso de dao.FindManyCryptoVote: " + err.Error()
		err = errors.New(z)
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
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   strings.Title(strings.ToLower(strings.TrimSpace(name))),
		Symbol: strings.ToUpper(strings.TrimSpace(symbol)),
	}

	// se pelo menos um dos filtros não for empty pode seguir com a busca
	var validate bool = false
	var err error
	if validateCryptoVoteArgumentNotEmpty(filterCryptoVote.Name) || validateCryptoVoteArgumentNotEmpty(filterCryptoVote.Symbol) {
		validate = true
	} else {
		z := "[cryptovote.validation] um dos filtros deve ser not empty"
		err = errors.New(z)
		return retrievedCryptoVote, err
	}

	// se pelo menos um dos filtros não for empty continua com a busca
	if validate {
		// convertendo de []model.CryptoVote para json apenas os campos que não estão nulos
		jsonData, err := json.Marshal(filterCryptoVote)
		if err != nil {
			z := "[cryptovote.json] Problemas para fazer Marshal de CryptoVote para json: " + err.Error()
			err = errors.New(z)
			return retrievedCryptoVote, err
		}

		filter := bson.M{}
		err = json.Unmarshal(jsonData, &filter)
		if err != nil {
			z := "[cryptovote.json] Problemas para fazer Unmarshal de json para bson.M: " + err.Error()
			err = errors.New(z)
			return retrievedCryptoVote, err
		}

		dao.SetCollectionName("cryptovotes")

		// usa a função criada no pacote dao
		mongodbClient, err := dao.GetMongoClientInstance()
		if err != nil {
			z := "[cryptovote.mongodb] Problemas no uso de GetMongoClientInstance: " + err.Error()
			err = errors.New(z)
			return retrievedCryptoVote, err
		}

		// usa a função criada no pacote dao
		retrievedCryptoVote, err = dao.FindOneCryptoVote(mongodbClient, filter)
		if err != nil {
			z := "[cryptovote.mongodb] Problemas no uso de dao.FindOneCryptoVote: " + err.Error()
			err = errors.New(z)
			return retrievedCryptoVote, err
		}
	}
	return retrievedCryptoVote, err
}

/*
	UpdateOneCryptoVoteByFilter faz uma atualização de todas as model.CryptoVote que satisfazem o filtro

	filterCryptoVote = model.FilterCryptoVote{
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
func UpdateOneCryptoVoteByFilter(filterCryptoVote model.FilterCryptoVote, cryptoNewData model.CryptoVote) (model.CryptoVote, error) {
	var retrievedCryptoVote model.CryptoVote = model.CryptoVote{}

	dao.SetCollectionName("cryptovotes")

	// usa a função criada no pacote dao
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		z := "[cryptovote.mongodb] Problemas no uso de GetMongoClientInstance: " + err.Error()
		err = errors.New(z)
		return retrievedCryptoVote, err
	}

	cryptoNewData.Name = strings.Title(strings.ToLower(strings.TrimSpace(cryptoNewData.Name)))
	cryptoNewData.Symbol = strings.ToUpper(strings.TrimSpace(cryptoNewData.Symbol))
	cryptoNewData.Sum = cryptoNewData.Qtd_Upvote - cryptoNewData.Qtd_Downvote

	// usa a função criada no pacote bo
	validate, err := ValidateCryptoVote(cryptoNewData)

	if validate {
		filterCryptoVote.Name = strings.Title(strings.ToLower(strings.TrimSpace(filterCryptoVote.Name)))
		filterCryptoVote.Symbol = strings.ToUpper(strings.TrimSpace(filterCryptoVote.Symbol))

		retrievedCryptoVote, err = RetrieveOneCryptoVote(filterCryptoVote.Name, filterCryptoVote.Symbol)
		if err != nil {
			return retrievedCryptoVote, err
		}

		if !retrievedCryptoVote.Id.IsZero() {
			// caso os novos dados respeitem as regras de validateCryptoVoteUniqueData
			// realiza a atualização
			idFilter := bson.M{"_id": retrievedCryptoVote.Id}

			// convertendo de model.CryptoVote para json apenas os campos que não estão nulos
			jsonData, err := json.Marshal(cryptoNewData)
			if err != nil {
				z := "[cryptovote.json] Problemas para fazer Marshal de CryptoVote para json: " + err.Error()
				err = errors.New(z)
				return retrievedCryptoVote, err
			}

			// convertendo de json para bson
			bsonCryptoNewData := bson.M{}
			err = json.Unmarshal(jsonData, &bsonCryptoNewData)
			if err != nil {
				z := "[cryptovote.json] Problemas para fazer Unmarshal de json para bson.M: " + err.Error()
				err = errors.New(z)
				return retrievedCryptoVote, err
			}

			newData := bson.M{
				"$set": bsonCryptoNewData,
			}

			// atualização
			retrievedCryptoVote, err = dao.UpdateOneCryptoVote(mongodbClient, idFilter, newData)
			if err != nil {
				z := "[cryptovote.mongodb] Problemas no uso de dao.UpdateOneCryptoVote: " + err.Error()
				err = errors.New(z)
				return retrievedCryptoVote, err
			}
		}
	}
	return retrievedCryptoVote, err
}

/*
	DeleteAllCryptoVoteByFilter faz uma deleção de todas as model.CryptoVote que satisfazem o filtro

	filterCryptoVote = model.FilterCryptoVote{
		Name:   "Bitcoin",
		Symbol: "",
	}

	retorno
	a quantidade de model.CryptoVote deletadas do banco, testes realizados como o mongoDB
*/
func DeleteAllCryptoVoteByFilter(filterCryptoVote model.FilterCryptoVote) (int64, error) {
	var deletedCount int64 = int64(0)
	var err error

	// popular os args no padrão para criar o filtro
	filterCryptoVote.Name = strings.Title(strings.ToLower(strings.TrimSpace(filterCryptoVote.Name)))
	filterCryptoVote.Symbol = strings.ToUpper(strings.TrimSpace(filterCryptoVote.Symbol))

	// se pelo menos um dos filtros não for empty pode seguir com a busca
	if validateCryptoVoteArgumentNotEmpty(filterCryptoVote.Name) || validateCryptoVoteArgumentNotEmpty(filterCryptoVote.Symbol) {
		dao.SetCollectionName("cryptovotes")

		// usa a função criada no pacote dao
		mongodbClient, err := dao.GetMongoClientInstance()
		if err != nil {
			z := "[cryptovote.mongodb] Problemas no uso de GetMongoClientInstance: " + err.Error()
			err = errors.New(z)
			return deletedCount, err
		}

		// convertendo de CryptoVote para json
		jsonData, err := json.Marshal(filterCryptoVote)
		if err != nil {
			z := "[cryptovote.json] Problemas para fazer Marshal de CryptoVote para jsonData: " + err.Error()
			err = errors.New(z)
			return deletedCount, err
		}

		// convertendo de json para bson
		var filter = bson.M{}
		err = json.Unmarshal(jsonData, &filter)
		if err != nil {
			z := "[cryptovote.json] Problemas para fazer Unmarshal de json para bson.M: " + err.Error()
			err = errors.New(z)
			return deletedCount, err
		}

		// usa a função criada no pacote dao
		deleteResult, err := dao.DeleteManyCryptoVote(mongodbClient, filter)
		deletedCount = deleteResult.DeletedCount
		if err != nil {
			z := "Problemas no uso de dao.DeleteManyCryptoVote: " + err.Error()
			log.Print(z)
			return deletedCount, err
		}
	} else {
		z := "[cryptovote.validation] um dos filtros deve ser not empty " + err.Error()
		err = errors.New(z)
		return deletedCount, err
	}

	return deletedCount, err
}

func DeleteAllCryptoVote() (int64, error) {
	var deletedCount int64 = int64(0)

	dao.SetCollectionName("cryptovotes")

	// usa a função criada no pacote dao
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		z := "[cryptovote.mongodb] Problemas no uso de GetMongoClientInstance: " + err.Error()
		err = errors.New(z)
		return deletedCount, err
	}

	// convertendo de json para bson
	var filter = bson.M{}

	// usa a função criada no pacote dao
	deleteResult, err := dao.DeleteManyCryptoVote(mongodbClient, filter)
	deletedCount = deleteResult.DeletedCount
	if err != nil {
		z := "Problemas no uso de dao.DeleteManyCryptoVote: " + err.Error()
		log.Print(z)
		return deletedCount, err
	}
	return deletedCount, err
}

/*
	AddUpVote faz uma adição de um voto em UpVote de uma CryptoVote que satisfazem os argumentos
	de entrada usando a função RetrieveOneCryptoVote(name,symbol) para localizar uma CryptoVote.
	Lembrando que cada CryptoVote não pode ter name ou symbol repetidos.

	retorno
	nil se não tiver problema ou erro caso contrário
*/
func AddUpVote(filterCryptoVote model.FilterCryptoVote) (model.CryptoVote, error) {
	// usa a função criada no arquivo cryptovoteBO.go pacote bo
	retrievedCryptoVote, err := RetrieveOneCryptoVote(filterCryptoVote.Name, filterCryptoVote.Symbol)
	if err != nil {
		return retrievedCryptoVote, err
	}

	if !retrievedCryptoVote.Id.IsZero() {
		// soma valor atual de Qtd_Upvote +1
		newQtd_Upvote := retrievedCryptoVote.Qtd_Upvote + 1

		// atualiza sempre o total Up - Down
		newSum := newQtd_Upvote - retrievedCryptoVote.Qtd_Downvote

		typeVote := "qtd_upvote"

		retrievedCryptoVote, err = updateVote(retrievedCryptoVote, typeVote, newQtd_Upvote, newSum)
	}
	return retrievedCryptoVote, err
}

/*
	AddDownVote faz uma adição de um voto em DownVote de uma CryptoVote que satisfazem os argumentos
	de entrada usando a função RetrieveOneCryptoVote(name,symbol) para localizar uma CryptoVote.
	Lembrando que cada CryptoVote não pode ter name ou symbol repetidos.

	retorno
	nil se não tiver problema ou erro caso contrário
*/
func AddDownVote(filterCryptoVote model.FilterCryptoVote) (model.CryptoVote, error) {
	// usa a função criada no arquivo cryptovoteBO.go pacote bo
	retrievedCryptoVote, err := RetrieveOneCryptoVote(filterCryptoVote.Name, filterCryptoVote.Symbol)
	if err != nil {
		return retrievedCryptoVote, err
	}

	if !retrievedCryptoVote.Id.IsZero() {
		// soma valor atual de Qtd_Downvote +1
		newQtd_Downvote := retrievedCryptoVote.Qtd_Downvote + 1

		// atualiza sempre o total Up - Down
		newSum := retrievedCryptoVote.Qtd_Upvote - newQtd_Downvote

		typeVote := "qtd_downvote"

		retrievedCryptoVote, err = updateVote(retrievedCryptoVote, typeVote, newQtd_Downvote, newSum)
		if err != nil {
			z := "[cryptovote.mongodb] Problemas no uso de updateVote: " + err.Error()
			err = errors.New(z)
			return retrievedCryptoVote, err
		}
	}
	return retrievedCryptoVote, err
}

func updateVote(retrievedCryptoVote model.CryptoVote, typeVote string, newQtd int, newSum int) (model.CryptoVote, error) {
	dao.SetCollectionName("cryptovotes")

	// usa a função criada no pacote dao
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		z := "[cryptovote.mongodb] Problemas no uso de GetMongoClientInstance: " + err.Error()
		err = errors.New(z)
		return retrievedCryptoVote, err
	}

	// cria filtro com id para localizar dado
	filter := bson.M{"_id": retrievedCryptoVote.Id}

	newData := bson.M{
		"$set": bson.M{
			typeVote: newQtd,
			"sum":    newSum,
		},
	}

	retrievedCryptoVote, err = dao.UpdateOneCryptoVote(mongodbClient, filter, newData)
	if err != nil {
		z := "[cryptovote.mongodb] Problemas no uso de dao.UpdateOneCryptoVote: " + err.Error()
		err = errors.New(z)
		return retrievedCryptoVote, err
	}

	return retrievedCryptoVote, err
}

/*
	SumaryAllCryptoVote faz uma busca ordenando as CryptoVote pelo campo Sum
	entrada pageSize representa o total de CryptoVote retornadas,
	campo Sum é o resultado de valor absoluto (Upvote - DownVote)

	caso pageSize == zero, usa o padrão 10
	flag_zero = 0
	flag_default_page_size = 10

	retorno
	slice de []model.SumaryVote
	nil se não tiver problema ou erro caso contrário
*/
func SumaryAllCryptoVote(pageSize int64) ([]model.SumaryVote, error) {
	var sumaryCryptoVotes []model.SumaryVote
	var retrievedCryptoVotes []model.CryptoVote
	var filterCryptoVote model.FilterCryptoVote
	var err error

	var flag_zero = 0
	var flag_default_page_size = 10
	// caso seja vazio ou zero usa o valor padrao 10
	if pageSize == int64(flag_zero) {
		pageSize = int64(flag_default_page_size)
	}

	dao.SetCollectionName("cryptovotes")

	// usa a função criada no pacote dao
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		z := "[cryptovote.mongodb] Problemas no uso de GetMongoClientInstance: " + err.Error()
		err = errors.New(z)
		return sumaryCryptoVotes, err
	}

	// se for nil precisamos criar um filtro vazio
	filterCryptoVote = model.FilterCryptoVote{
		Name:   "",
		Symbol: "",
	}

	// convertendo de CryptoVote para json
	jsonData, err := json.Marshal(filterCryptoVote)
	if err != nil {
		z := "[cryptovote.json] Problemas para fazer Marshal de CryptoVote para json: " + err.Error()
		err = errors.New(z)
		return sumaryCryptoVotes, err
	}

	// convertendo de json para bson
	var filter = bson.M{}
	err = json.Unmarshal(jsonData, &filter)
	if err != nil {
		z := "[cryptovote.json] Problemas para fazer Unmarshal de json para bson.M: " + err.Error()
		err = errors.New(z)
		return sumaryCryptoVotes, err
	}

	retrievedCryptoVotes, err = dao.FindManyCryptoVoteLimitedSortedByField(mongodbClient, filter, pageSize, "sum", -1)
	if err != nil {
		z := "[cryptovote.mongodb] Problemas em dao.FindManyCryptoVoteLimitedSortedByField: " + err.Error()
		err = errors.New(z)
		return sumaryCryptoVotes, err
	}

	if int64(len(retrievedCryptoVotes)) > pageSize {
		z := fmt.Sprintf("[cryptovote.mongodb] Problemas para garantir o tamanho limite de pageSize(%d) CryptoVotes", pageSize)
		err = errors.New(z)
		return nil, err
	}

	sumaryCryptoVotes = nil
	for _, cryptoVotes := range retrievedCryptoVotes {

		var sumary model.SumaryVote
		sumary.CryptoVote = cryptoVotes
		sumary.SumAbsolute = utils.Abs(cryptoVotes.Sum)

		var sumType string
		if cryptoVotes.Qtd_Upvote == cryptoVotes.Qtd_Downvote {
			sumType = "Equal"
		} else {
			if cryptoVotes.Qtd_Upvote < cryptoVotes.Qtd_Downvote {
				sumType = "Up vote"
			}
			sumType = "Down vote"
		}
		sumary.SumType = sumType

		sumaryCryptoVotes = append(sumaryCryptoVotes, sumary)
	}
	return sumaryCryptoVotes, err
}
