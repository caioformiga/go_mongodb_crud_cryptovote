package test

import (
	"testing"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/bo"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/mock"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestValidateCryptoVoteInvalidDataMissingName(t *testing.T) {
	var in = utils.LoadOneNewCryptoVoteDataFromArgs("", "FORM")
	// try with empty name CryptoVote, expects it to return validation error
	doTestValidateCryptoVoteInvalidData(t, in)
}

func TestValidateCryptoVoteInvalidDataMissingSymbol(t *testing.T) {
	var in = utils.LoadOneNewCryptoVoteDataFromArgs("FormiCOIN", "")
	// try with empty symbol CryptoVote, expects it to return validation error
	doTestValidateCryptoVoteInvalidData(t, in)
}

func TestValidateCryptoVoteInvalidDataNameToLarge(t *testing.T) {
	var in = utils.LoadOneNewCryptoVoteDataFromArgs("Formitcoinhjauheauhuehuahueuauehuahuheuahuehua", "FORM")
	// try with field name > 30 CryptoVote, expects it to return validation error
	doTestValidateCryptoVoteInvalidData(t, in)
}

func TestValidateCryptoVoteInvalidDataSymbolToLarge(t *testing.T) {
	var in = utils.LoadOneNewCryptoVoteDataFromArgs("Formitcoin", "FORMFORMFORMFORMFORMFORM")
	// try with field symbol > 6 CryptoVote, expects it to return validation error
	doTestValidateCryptoVoteInvalidData(t, in)
}

func TestValidateCryptoVoteInvalidDataUpvoteNegative(t *testing.T) {
	var in = utils.LoadOneNewCryptoVoteDataFromArgs("Cacau Coin", "CC")
	in.Qtd_Upvote = -1
	// try with Qtd_Upvote lesser then zero, expects it to return validation error
	doTestValidateCryptoVoteInvalidData(t, in)
}

func TestValidateCryptoVoteInvalidDataDownVoteNegative(t *testing.T) {
	var in = utils.LoadOneNewCryptoVoteDataFromArgs("Laranja Coin", "LC")
	in.Qtd_Downvote = -1
	// try with Qtd_Downvote lesser then zero, expects it to return validation error
	doTestValidateCryptoVoteInvalidData(t, in)
}

/*
	doTestValidateCryptoVoteInvalidData is a helper function to perform test using invalid data,
	expects it to return validation error.
*/
func doTestValidateCryptoVoteInvalidData(t *testing.T, in model.CryptoVote) {
	var boInstance bo.CryptoVoteBO = bo.NewCryptoVoteBO(nil)
	boInstance.ImplDAO.GetDbService().SetDatabaseName("test_cryptovotes")
	validate, err := boInstance.ValidateCryptoVote(in)
	assert.NotNil(t, err, "err should not be nil")
	assert.False(t, validate, "validate should be false")
}

/*
	configBOmockedDAO is a helper function to create a boInstance within a mockDAO to simulate
	database-related actions.
*/
func configBOmockedDAO(t *testing.T) (bo.CryptoVoteBO, *mock.MockInterfaceCryptoVoteDAO) {
	// Create the controller
	var mockCtrl = gomock.NewController(t)
	defer mockCtrl.Finish()

	// Create the mock object using controller
	var mockDAO = mock.NewMockInterfaceCryptoVoteDAO(mockCtrl)

	// Pass a mock at constructor define a mock implementation of a interfaces.InterfaceCryptoVoteDAO,
	// to simulate DAO
	var boInstance = bo.NewCryptoVoteBO(mockDAO)
	return boInstance, mockDAO
}

func TestValidateCryptoVoteValidFullData(t *testing.T) {
	var boInstance bo.CryptoVoteBO
	var mockDAO *mock.MockInterfaceCryptoVoteDAO
	var in model.CryptoVote = utils.LoadOneNewCryptoVoteDataFromArgs("FormiCOIN", "FORM")

	boInstance, mockDAO = configBOmockedDAO(t)
	if mockDAO != nil {
		// prepare mock simulation at below order
		gomock.InOrder(
			mockDAO.EXPECT().FindMany(gomock.Any()).Times(2).Return([]model.CryptoVote{}, nil),
		)
	}

	// try with full data CryptoVote,  expects it to return validation true
	validate, err := boInstance.ValidateCryptoVote(in)
	assert.Nil(t, err, "err should be nil")
	assert.True(t, validate, "validate should be true")
}

func TestValidateCryptoVoteInvalidNotUniqueData(t *testing.T) {
	var boInstance bo.CryptoVoteBO
	var mockDAO *mock.MockInterfaceCryptoVoteDAO
	var in model.CryptoVote = utils.LoadOneNewCryptoVoteDataFromArgs("FormiCOIN", "FORM")

	boInstance, mockDAO = configBOmockedDAO(t)
	if mockDAO != nil {
		// add CryptoVote name = "FormiCOIN" symbol = "FORM"
		list := []model.CryptoVote{in}

		// prepare mock simulation at below order
		gomock.InOrder(
			mockDAO.EXPECT().FindMany(gomock.Any()).Times(1).Return(list, nil),
		)
	}

	// try with full data CryptoVote, expects it to return true in validation
	validate, err := boInstance.ValidateCryptoVote(in)
	assert.NotNil(t, err, "err should not be nil")
	assert.False(t, validate, "validate should be false")
}
