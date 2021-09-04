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

func TestSumaryAllCryptoVoteZeroPageSize(t *testing.T) {
	var boInstance bo.CryptoVoteBO
	var mockDAO *mock.MockInterfaceCryptoVoteDAO
	var sumaryCryptoVotes []model.SumaryCryptoVote
	var out []model.CryptoVote
	var err error

	boInstance, mockDAO = ConfigBOmockedDAO(t)
	if mockDAO != nil {

		out, _ = utils.LoadManyCryptoVoteDataFromJson(utils.JsonOutDataSorted)
		filter, _ := utils.MarshalFilterCryptoVoteToBsonFilter(utils.LoadOneNewEmptyFilterCryptoVote())
		orderType := -1
		sortFieldName := "sum_absolute"
		pageSize := int64(10)

		// prepare mock simulation at below order
		gomock.InOrder(
			mockDAO.EXPECT().FindManyLimitedSortedByField(filter, pageSize, sortFieldName, orderType).Times(1).Return(out, err),
		)
	}

	zeroPageSize := int64(0)
	sumaryCryptoVotes, err = boInstance.SumaryAllCryptoVote(zeroPageSize)
	assert.Nil(t, err, "err should be nil")
	assert.NotNil(t, sumaryCryptoVotes, "err should not be nil")

	totalCount := int64(3)
	retrivedCount := int64(len(sumaryCryptoVotes))
	assert.Equal(t, retrivedCount, totalCount, "they should be equal")
}

func TestSumaryAllCryptoVoteSorted(t *testing.T) {
	var boInstance bo.CryptoVoteBO
	var mockDAO *mock.MockInterfaceCryptoVoteDAO
	var sumaryCryptoVotes []model.SumaryCryptoVote
	var out []model.CryptoVote
	var err error

	boInstance, mockDAO = ConfigBOmockedDAO(t)
	if mockDAO != nil {

		out, _ = utils.LoadManyCryptoVoteDataFromJson(utils.JsonOutDataSorted)
		filter, _ := utils.MarshalFilterCryptoVoteToBsonFilter(utils.LoadOneNewEmptyFilterCryptoVote())
		orderType := -1
		sortFieldName := "sum_absolute"
		pageSize := int64(10)

		// prepare mock simulation at below order
		gomock.InOrder(
			mockDAO.EXPECT().FindManyLimitedSortedByField(filter, pageSize, sortFieldName, orderType).Times(1).Return(out, err),
		)
	}

	zeroPageSize := int64(0)
	sumaryCryptoVotes, err = boInstance.SumaryAllCryptoVote(zeroPageSize)
	assert.Nil(t, err, "err should be nil")
	assert.NotNil(t, sumaryCryptoVotes, "err should not be nil")

	prim := sumaryCryptoVotes[0]
	assert.Equal(t, prim.Token, "Klever/KLV", "they should be equal")

	seg := sumaryCryptoVotes[1]
	assert.Equal(t, seg.Token, "Bitcoin/BTC", "they should be equal")

	ult := sumaryCryptoVotes[2]
	assert.Equal(t, ult.Token, "Ethereum/ETH", "they should be equal")
}
