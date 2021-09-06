package test

import (
	"testing"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/bo"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestDeleteAllCryptoVote(t *testing.T) {
	var boInstance bo.CryptoVoteBO
	var mockDAO *mock.MockInterfaceCryptoVoteDAO
	var err error
	const ONE_COUNT int64 = 0

	boInstance, mockDAO = ConfigBOmockedDAO(t)
	if mockDAO != nil {
		filter := bson.M{}
		// prepare mock simulation at below order
		gomock.InOrder(
			mockDAO.EXPECT().DeleteMany(filter).Times(1).Return(ONE_COUNT, err),
		)
	}

	deletedCount, err := boInstance.DeleteAllCryptoVote()
	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, deletedCount, ONE_COUNT, "they should be equal")
}
