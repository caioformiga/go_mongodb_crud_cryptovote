package bo

import (
	"errors"
	"log"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/dao"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"

	"go.mongodb.org/mongo-driver/bson"
)

func ValidateCryptoVote(crypto model.CryptoCurrency, qtd_upvote int, qtd_downvote int) (bool, error) {
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

	validate, err := ValidateCryptoCurrency(crypto.Name, crypto.Symbol)
	if err != nil {
		log.Fatal(err)
	}
	return validate, nil
}

func CreateCryptoVote(crypto model.CryptoCurrency, qtd_upvote int, qtd_downvote int) model.CryptoVote {
	// usa a função criada no pacote bo
	_, err := ValidateCryptoVote(crypto, qtd_upvote, qtd_downvote)
	if err != nil {
		log.Fatalf("Problemas na validação de dados da nova CryptoCurrency: %v", err)
	}

	dao.SetCollectioName("cryptovotes")

	// usa a função criada no arquivo mongodb.go pacote main (raiz) do projeto
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		log.Fatal(err)
	}

	objCryptoVote := model.CryptoVote{
		Id:           [12]byte{},
		Crypto:       crypto,
		Qtd_Upvote:   qtd_upvote,
		Qtd_Downvote: qtd_downvote,
	}

	// usa a função criada no pacote dao
	insertResult, err := dao.CreateCryptoVote(mongodbClient, objCryptoVote)
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

func DeleteCryptoVoteByFilter(filter bson.M) int64 {
	// usa a função criada no arquivo mongodb.go pacote main (raiz) do projeto
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
