package bo

import (
	"errors"
	"log"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/dao"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"

	"go.mongodb.org/mongo-driver/bson"
)

func validateCryptoVoteArgumentNotEmpty(arg string) bool {
	validate := false

	if len(arg) > 0 {
		validate = true
	} else {
		validate = false
	}
	return validate
}

func validateCryptoVoteArgumentLenght(arg string, lenght int) bool {
	validate := false

	if len(arg) <= lenght {
		validate = true
	} else {
		validate = false
	}
	return validate
}

func validateCryptoVoteArgumentNotNegative(arg int) bool {
	validate := false

	if arg >= 0 {
		validate = true
	} else {
		validate = false

	}
	return validate
}

/*
	ValidateCryptoVoteArguments verifica se algum campo está fora do valor defalt
	não atente aos critérios de unique abaixo
	Name não pode ser vazio
	Symbol não pode ser vazio
	Qtd_Upvote não pode ser menor do que zero
	Qtd_Downvote não pode ser menor do que zero
*/
func ValidateCryptoVoteArguments(crypto model.CryptoVote) (bool, error) {
	var validate bool = false

	validate = validateCryptoVoteArgumentNotEmpty(crypto.Name)
	if !validate {
		return validate, errors.New("[cryptovote.validationBO] name não pode ser vazio")
	}

	validate = validateCryptoVoteArgumentLenght(crypto.Name, 30)
	if !validate {
		return validate, errors.New("[cryptovote.validationBO] name não pode ter maior do que 30 caracteres")
	}

	validate = validateCryptoVoteArgumentNotEmpty(crypto.Symbol)
	if !validate {
		return validate, errors.New("[cryptovote.validationBO] symbol não pode ser vazio")
	}

	validate = validateCryptoVoteArgumentLenght(crypto.Symbol, 6)
	if !validate {
		return validate, errors.New("[cryptovote.validationBO] name não pode ter maior do que 6 caracteres")
	}

	validate = validateCryptoVoteArgumentNotNegative(crypto.Qtd_Upvote)
	if !validate {
		validate = false
		return validate, errors.New("[cryptovote.validationBO] qtd_upvote não pode ser menor do que zero")
	}

	validate = validateCryptoVoteArgumentNotNegative(crypto.Qtd_Downvote)
	if !validate {
		validate = false
		validate = false
		return validate, errors.New("[cryptovote.validationBO] qtd_downvote não pode ser menor do que zero")
	}
	return validate, nil
}

func validateUnique(key string, value string) (bool, error) {
	var validate bool = true
	var retrivedCryptoVotes []model.CryptoVote

	dao.SetCollectionName("cryptovotes")

	// usa a função criada no pacote dao
	mongodbClient, err := dao.GetMongoClientInstance()
	if err != nil {
		z := "Problemas no uso de GetMongoClientInstance: " + err.Error()
		log.Print(z)
	}

	// usa a função criada no pacote dao
	retrivedCryptoVotes, err = mongodbDAO.FindManyCryptoVote(mongodbClient, bson.M{key: value})
	if err != nil {
		z := "Problemas no uso de FindManyCryptoVote: " + err.Error()
		log.Print(z)
	}

	if retrivedCryptoVotes != nil {
		validate = false
		return validate, errors.New("[cryptovote.validationBO] campo(" + key + ") informado já exite, escolha outro diferente de " + value)
	}
	return validate, err
}

func validateCryptoVoteUniqueSymbol(value string) (bool, error) {
	key := "symbol"
	return validateUnique(key, value)
}

func validateCryptoVoteUniqueName(value string) (bool, error) {
	key := "name"
	return validateUnique(key, value)
}

/*
	ValidateCryptoVoteUniqueData verifica se não existe no banco alguma CryptoVote que
	não atente aos critérios de unique abaixo
	name : unique
	symbol : unique
*/
func ValidateCryptoVoteUniqueData(name string, symbol string) (bool, error) {
	var validate bool = false
	var err error

	// se não for vazio
	if len(name) > 0 {
		validate, err = validateCryptoVoteUniqueName(name)
		if !validate {
			return validate, err
		}
	}

	// se não for vazio
	if len(symbol) > 0 {
		validate, err = validateCryptoVoteUniqueSymbol(symbol)
		if !validate {
			return validate, err
		}
	}
	return validate, nil
}

/*
	ValidateCryptoVote recebe os campos e faz a validação
	name : não pode ser vazio, len(name) > 0
	symbol : não pode ser vazio, len(name) > 0
	qtd_upvote : não pode ser menor do que zero, qtd_upvote >= 0
	qtd_downvote : não pode ser menor do que zero, qtd_upvote >= 0
*/
func ValidateCryptoVote(crypto model.CryptoVote) (bool, error) {
	var validate = false
	var err error

	// usa a função criada no pacote bo
	validate, err = ValidateCryptoVoteArguments(crypto)
	if err != nil {
		z := "Problemas na validação de dados da nova CryptoVote: " + err.Error()
		log.Print(z)
		return validate, err
	} else {
		validate = true
	}

	// usa a função criada no pacote bo
	validate, err = ValidateCryptoVoteUniqueData(crypto.Name, crypto.Symbol)
	if err != nil {
		z := "Problemas na validação unique da nova CryptoVote: " + err.Error()
		log.Print(z)
		return validate, err
	} else {
		validate = true
	}
	return validate, err
}
