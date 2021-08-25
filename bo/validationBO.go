package bo

import (
	"errors"
	"log"

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
	Checks each field using below criteria:
	Name cannot be empty
	Symbol cannot be empty
	Qtd_Upvote cannot be less than zero
	Qtd_Downvote cannot be less than zero
*/
func ValidateCryptoVoteArguments(crypto model.CryptoVote) (bool, error) {
	var validate bool = false

	validate = validateCryptoVoteArgumentNotEmpty(crypto.Name)
	if !validate {
		return validate, errors.New("[cryptovote.validationBO] name can't be null")
	}

	validate = validateCryptoVoteArgumentLenght(crypto.Name, 30)
	if !validate {
		return validate, errors.New("[cryptovote.validationBO] name can't have more than 30 characters")
	}

	validate = validateCryptoVoteArgumentNotEmpty(crypto.Symbol)
	if !validate {
		return validate, errors.New("[cryptovote.validationBO] symbol can't be null")
	}

	validate = validateCryptoVoteArgumentLenght(crypto.Symbol, 6)
	if !validate {
		return validate, errors.New("[cryptovote.validationBO] name can't have more than 6 characters")
	}

	validate = validateCryptoVoteArgumentNotNegative(crypto.Qtd_Upvote)
	if !validate {
		validate = false
		return validate, errors.New("[cryptovote.validationBO] qtd_upvote can't have less then zero")
	}

	validate = validateCryptoVoteArgumentNotNegative(crypto.Qtd_Downvote)
	if !validate {
		validate = false
		validate = false
		return validate, errors.New("[cryptovote.validationBO] qtd_downvote can't have less then zero")
	}
	return validate, nil
}

func (c CryptoVoteBO) validateUnique(key string, value string) (bool, error) {
	var retrivedCryptoVotes []model.CryptoVote
	var err error

	// use function from dao package
	retrivedCryptoVotes, err = c.ImplDAO.FindMany(bson.M{key: value})
	if err != nil {
		z := "Problems using FindManyCryptoVote: " + err.Error()
		log.Print(z)
	}

	var validate bool = true

	if len(retrivedCryptoVotes) > 0 {
		validate = false
		return validate, errors.New("[cryptovote.validationBO] field(" + key + ") already exists, choose anoter value different from " + value)
	}
	return validate, err
}

/*
	Check if there is a CryptoVote with similar data, usgin unique criteria below:
	name : unique
	symbol : unique
*/
func (c CryptoVoteBO) ValidateCryptoVoteUniqueData(name string, symbol string) (bool, error) {
	var validate bool = false
	var key string
	var value string
	var err error

	// se não for vazio
	if len(name) > 0 {
		key = "name"
		value = name
		validate, err = c.validateUnique(key, value)
		if !validate {
			return validate, err
		}
	}

	// se não for vazio
	if len(symbol) > 0 {
		key = "symbol"
		value = symbol
		validate, err = c.validateUnique(key, value)
		if !validate {
			return validate, err
		}
	}
	return validate, nil
}

/*
	External function to handle all validation process
*/
func (c CryptoVoteBO) ValidateCryptoVote(crypto model.CryptoVote) (bool, error) {
	var validate = false
	var err error

	// uses function from bo package
	validate, err = ValidateCryptoVoteArguments(crypto)
	if err != nil {
		return validate, err
	} else {
		validate = true
	}

	// uses function from bo package
	validate, err = c.ValidateCryptoVoteUniqueData(crypto.Name, crypto.Symbol)
	if err != nil {
		return validate, err
	} else {
		validate = true
	}
	return validate, err
}
