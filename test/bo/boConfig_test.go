package test

import (
	"testing"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/bo"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/mock"
	"github.com/golang/mock/gomock"
)

/*
	ConfigBOmockedDAO is a helper function to create a boInstance within a mockDAO to simulate
	database-related actions.
*/
func ConfigBOmockedDAO(t *testing.T) (bo.CryptoVoteBO, *mock.MockInterfaceCryptoVoteDAO) {
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
