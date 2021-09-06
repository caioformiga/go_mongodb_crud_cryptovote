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

func TestDeleteAllCryptoVoteByFilterEmptyData(t *testing.T) {
	var boInstance bo.CryptoVoteBO
	var mockDAO *mock.MockInterfaceCryptoVoteDAO
	var out []model.CryptoVote
	var filterCryptoVote model.FilterCryptoVote
	var err error
	const ZERO_COUNT int64 = 0

	filterCryptoVote = utils.LoadOneNewEmptyFilterCryptoVote()
	out, _ = utils.LoadManyCryptoVoteDataFromJson(utils.JsonOutDataSorted)
	totalCount := int64(len(out))

	boInstance, mockDAO = ConfigBOmockedDAO(t)
	if mockDAO != nil {

		// prepare mock simulation at below order
		gomock.InOrder(
			mockDAO.EXPECT().DeleteMany(filterCryptoVote).Times(1).Return(totalCount, err),
		)
	}

	deletedCount, err := boInstance.DeleteAllCryptoVoteByFilter(filterCryptoVote)
	assert.NotNil(t, err, "err should not be nil")
	assert.Equal(t, deletedCount, ZERO_COUNT, "they should be equal")
}
