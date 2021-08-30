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

func TestAddDownVoteValidFilterByNameAndSymbol(t *testing.T) {
	var boInstance bo.CryptoVoteBO
	var mockDAO *mock.MockInterfaceCryptoVoteDAO
	var in = utils.LoadOneNewCryptoVoteDataFromArgs("Formiga", "FORM")
	var out = utils.LoadOneNewCryptoVoteDataFromArgs("Formiga", "FORM")
	var retrievedCryptoVote model.CryptoVote
	var err error

	boInstance, mockDAO = configBOmockedDAO(t)
	if mockDAO != nil {

		out.Qtd_Downvote = out.Qtd_Downvote + 1
		out.Sum = in.Qtd_Upvote - out.Qtd_Downvote
		out.SumAbsolute = utils.Abs(out.Sum)

		// prepare mock simulation at below order
		gomock.InOrder(
			mockDAO.EXPECT().FindOne(gomock.Any()).Times(1).Return(in, nil),
			mockDAO.EXPECT().UpdateOne(gomock.Any(), gomock.Any()).Times(1).Return(out, nil),
		)
	}

	var filterCryptoVote = model.FilterCryptoVote{
		Name:   "Formiga",
		Symbol: "FORM",
	}

	retrievedCryptoVote, err = boInstance.AddDownVote(filterCryptoVote)
	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, retrievedCryptoVote.Name, in.Name, "they should be equal")
	assert.Equal(t, retrievedCryptoVote.Symbol, in.Symbol, "they should be equal")
	newQtd := in.Qtd_Downvote + 1
	assert.Equal(t, retrievedCryptoVote.Qtd_Downvote, newQtd, "they should be equal")
	newSum := in.Qtd_Upvote - newQtd
	newSumAbsolute := utils.Abs(newSum)
	assert.Equal(t, retrievedCryptoVote.Sum, newSum, "they should be equal")
	assert.Equal(t, retrievedCryptoVote.SumAbsolute, newSumAbsolute, "they should be equal")
}

func TestAddDownVoteValidFilterByName(t *testing.T) {
	var boInstance bo.CryptoVoteBO
	var mockDAO *mock.MockInterfaceCryptoVoteDAO
	var in = utils.LoadOneNewCryptoVoteDataFromArgs("Formiga", "FORM")
	var out = utils.LoadOneNewCryptoVoteDataFromArgs("Formiga", "FORM")
	var retrievedCryptoVote model.CryptoVote
	var err error

	boInstance, mockDAO = configBOmockedDAO(t)
	if mockDAO != nil {

		out.Qtd_Downvote = out.Qtd_Downvote + 1
		// prepare mock simulation at below order
		gomock.InOrder(
			mockDAO.EXPECT().FindOne(gomock.Any()).Times(1).Return(in, nil),
			mockDAO.EXPECT().UpdateOne(gomock.Any(), gomock.Any()).Times(1).Return(out, nil),
		)
	}

	var filterCryptoVote = model.FilterCryptoVote{
		Name:   "Formiga",
		Symbol: "",
	}

	retrievedCryptoVote, err = boInstance.AddDownVote(filterCryptoVote)
	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, retrievedCryptoVote.Name, in.Name, "they should be equal")
	assert.Equal(t, retrievedCryptoVote.Symbol, in.Symbol, "they should be equal")
	newQtd := in.Qtd_Downvote + 1
	assert.Equal(t, retrievedCryptoVote.Qtd_Downvote, newQtd, "they should be equal")
}

func TestAddDownVoteValidFilterBySymbol(t *testing.T) {
	var boInstance bo.CryptoVoteBO
	var mockDAO *mock.MockInterfaceCryptoVoteDAO
	var in = utils.LoadOneNewCryptoVoteDataFromArgs("Formiga", "FORM")
	var out = utils.LoadOneNewCryptoVoteDataFromArgs("Formiga", "FORM")
	var retrievedCryptoVote model.CryptoVote
	var err error

	boInstance, mockDAO = configBOmockedDAO(t)
	if mockDAO != nil {

		out.Qtd_Downvote = out.Qtd_Downvote + 1
		// prepare mock simulation at below order
		gomock.InOrder(
			mockDAO.EXPECT().FindOne(gomock.Any()).Times(1).Return(in, nil),
			mockDAO.EXPECT().UpdateOne(gomock.Any(), gomock.Any()).Times(1).Return(out, nil),
		)
	}

	var filterCryptoVote = model.FilterCryptoVote{
		Name:   "",
		Symbol: "FORM",
	}

	retrievedCryptoVote, err = boInstance.AddDownVote(filterCryptoVote)
	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, retrievedCryptoVote.Name, in.Name, "they should be equal")
	assert.Equal(t, retrievedCryptoVote.Symbol, in.Symbol, "they should be equal")
	newQtd := in.Qtd_Downvote + 1
	assert.Equal(t, retrievedCryptoVote.Qtd_Downvote, newQtd, "they should be equal")
}

func TestAddDownVoteInvalidFilter(t *testing.T) {
	var boInstance bo.CryptoVoteBO
	var mockDAO *mock.MockInterfaceCryptoVoteDAO
	var retrievedCryptoVote model.CryptoVote
	var err error

	boInstance, mockDAO = configBOmockedDAO(t)
	if mockDAO != nil {
		// prepare mock simulation at below order
		gomock.InOrder(
			mockDAO.EXPECT().FindOne(gomock.Any()).Times(1).Return(model.CryptoVote{}, nil),
		)
	}

	retrievedCryptoVote, err = boInstance.AddDownVote(model.FilterCryptoVote{})
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, retrievedCryptoVote.Id.IsZero(), " should be true")
}
