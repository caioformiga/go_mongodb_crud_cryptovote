package test

import (
	"testing"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/bo"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/utils"
	"github.com/stretchr/testify/assert"
)

func TestRetrieveAllCryptoVoteByFilter(t *testing.T) {
	testRetrieveAllCryptoVoteByFilter0_Config(t)
	testRetrieveAllCryptoVoteByFilter1_FilterNull(t)
}

/*
	0
	configurando antes do teste
	remove todos os dados
	insere 3 dados
*/
func testRetrieveAllCryptoVoteByFilter0_Config(t *testing.T) {
	// limpa a coleção
	_, err := bo.DeleteAllCryptoVote()
	assert.Nil(t, err, "err should be nil")

	// carrega json data com 3 CrypytoVotes
	listIn, err := utils.Load_data(utils.JsonInData)
	assert.Nil(t, err, "err should be nil")

	tam := len(listIn)
	for i := 0; i < tam; i++ {
		cryptoVote := listIn[i]

		_, err := bo.CreateCryptoVote(cryptoVote)
		assert.Nil(t, err, "err should be nil")
	}
}

/*
	1
	recupera todos os registros usando filter null
	compara totalCount com retrivedCount
*/
func testRetrieveAllCryptoVoteByFilter1_FilterNull(t *testing.T) {
	// cria model vazio que sera convertido para filtro vazio
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   "",
		Symbol: "",
	}
	list, err := bo.RetrieveAllCryptoVoteByFilter(filterCryptoVote)

	totalCount := int32(3)
	retrivedCount := int32(len(list))

	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, retrivedCount, totalCount, "they should be equal")
}
