package test

import (
	"testing"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/bo"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/utils"
	"github.com/stretchr/testify/assert"
)

func TestRetrieveAllCryptoVoteByFilter(t *testing.T) {
	/*
		Instância que permite acessar os metodos implementados em bo.CryptoVoteBO
	*/
	var cryptoVoteBO bo.InterfaceCryptoVoteBO = bo.CryptoVoteBO{}

	testRetrieveAllCryptoVoteByFilter0_Config(t, cryptoVoteBO)
	testRetrieveAllCryptoVoteByFilter1_FilterNull(t, cryptoVoteBO)
}

/*
	0
	configurando antes do teste
	remove todos os dados
	insere 3 dados
*/
func testRetrieveAllCryptoVoteByFilter0_Config(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	// limpa a coleção
	_, err := cryptoVoteBO.DeleteAllCryptoVote()
	assert.Nil(t, err, "err should be nil")

	// carrega json data com 3 CrypytoVotes
	listIn, err := utils.Load_data(utils.JsonInData)
	assert.Nil(t, err, "err should be nil")

	tam := len(listIn)
	for i := 0; i < tam; i++ {
		cryptoVote := listIn[i]

		_, err := cryptoVoteBO.CreateCryptoVote(cryptoVote)
		assert.Nil(t, err, "err should be nil")
	}
}

/*
	1
	recupera todos os registros usando filter null
	compara totalCount com retrivedCount
*/
func testRetrieveAllCryptoVoteByFilter1_FilterNull(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	// cria model vazio que sera convertido para filtro vazio
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   "",
		Symbol: "",
	}
	list, err := cryptoVoteBO.RetrieveAllCryptoVoteByFilter(filterCryptoVote)

	totalCount := int32(3)
	retrivedCount := int32(len(list))

	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, retrivedCount, totalCount, "they should be equal")
}
