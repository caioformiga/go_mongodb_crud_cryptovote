package test

import (
	"testing"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/bo"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/utils"
	"github.com/stretchr/testify/assert"
)

func TestDeleteAllCryptoVoteByFilter(t *testing.T) {
	testDeleteAllCryptoVoteByFilter0_Config(t)
	testDeleteAllCryptoVoteByFilter1_FilterNull(t)
}

/*
	0
	configurando antes do teste
	insere dados
*/
func testDeleteAllCryptoVoteByFilter0_Config(t *testing.T) {
	// carrega json data com 3 CrypytoVotes
	// usa func do arquivo default_data.go
	listIn, err := utils.Load_data(utils.JsonInData)
	assert.Nil(t, err, "err should be nil")

	totalCount := 3
	assert.Equal(t, len(listIn), totalCount, "they should be equal")

	tam := len(listIn)
	for i := 0; i < tam; i++ {
		cryptoVote := listIn[i]

		_, err := bo.CreateCryptoVote(cryptoVote)
		assert.NotNil(t, err, "err should not be nil")
	}
}

/*
	1
	remove todos os registros usando filter null
	compara totalCount com deletedCount
*/
func testDeleteAllCryptoVoteByFilter1_FilterNull(t *testing.T) {
	// cria model vazio que sera convertido para filtro vazio
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   "",
		Symbol: "",
	}
	list, _ := bo.RetrieveAllCryptoVoteByFilter(filterCryptoVote)
	totalCount := int64(len(list))

	// limpa a coleção
	deletedCount, err := bo.DeleteAllCryptoVote()
	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, deletedCount, totalCount, "they should be equal")
}
