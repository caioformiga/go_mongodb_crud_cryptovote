package test

import (
	"testing"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/bo"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/utils"
	"github.com/stretchr/testify/assert"
)

func TestSumaryAllCryptoVote(t *testing.T) {
	/*
		Instância que permite acessar os metodos implementados em bo.CryptoVoteBO
	*/
	var cryptoVoteBO bo.InterfaceCryptoVoteBO = bo.CryptoVoteBO{}

	testSumaryAllCryptoVote0_Config(t, cryptoVoteBO)
	testSumaryAllCryptoVote1_FilterNull(t, cryptoVoteBO)
	testSumaryAllCryptoVote2_SortedDesc(t, cryptoVoteBO)
	testSumaryAllCryptoVote3_Limit(t, cryptoVoteBO)
}

/*
	0
	configurando antes do teste
	remove todos os dados
	insere 3 dados
*/
func testSumaryAllCryptoVote0_Config(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
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
*/
func testSumaryAllCryptoVote1_FilterNull(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	list, err := cryptoVoteBO.SumaryAllCryptoVote(10)

	totalCount := int32(3)
	retrivedCount := int32(len(list))

	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, retrivedCount, totalCount, "they should be equal")
}

/*
	2
	todos os registros usando filter null
	lista está ordenada descrescente
*/
func testSumaryAllCryptoVote2_SortedDesc(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	// cria model vazio que sera convertido para filtro vazio
	var filterCryptoVote = model.FilterCryptoVote{}

	filterCryptoVote.Symbol = "BTC"
	// adiciona 10 votos para crypto
	for i := 0; i < 10; i++ {
		_, err := cryptoVoteBO.AddUpVote(filterCryptoVote)
		assert.Nil(t, err, "err should be nil")
	}

	filterCryptoVote.Symbol = "ETH"
	// adiciona 5 votos para crypto
	for i := 0; i < 5; i++ {
		_, err := cryptoVoteBO.AddUpVote(filterCryptoVote)
		assert.Nil(t, err, "err should be nil")
	}

	filterCryptoVote.Symbol = "KLV"
	// adiciona 500 votos para crypto
	for i := 0; i < 500; i++ {
		_, err := cryptoVoteBO.AddUpVote(filterCryptoVote)
		assert.Nil(t, err, "err should be nil")
	}

	list, err := cryptoVoteBO.SumaryAllCryptoVote(10)
	assert.Nil(t, err, "err should be nil")

	prim := list[0]
	assert.Equal(t, prim.Token, "Klever/KLV", "they should be equal")

	seg := list[1]
	assert.Equal(t, seg.Token, "Bitcoin/BTC", "they should be equal")

	ult := list[2]
	assert.Equal(t, ult.Token, "Ethereum/ETH", "they should be equal")
}

/*
	3
	todos os registros usando filter null
	lista está ordenada descrescente
*/
func testSumaryAllCryptoVote3_Limit(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	// cria model vazio que sera convertido para filtro vazio
	var filterCryptoVote = model.FilterCryptoVote{}

	filterCryptoVote.Symbol = "BTC"
	// adiciona 10 votos para crypto
	for i := 0; i < 10; i++ {
		_, err := cryptoVoteBO.AddUpVote(filterCryptoVote)
		assert.Nil(t, err, "err should be nil")
	}

	filterCryptoVote.Symbol = "ETH"
	// adiciona 5 votos para crypto
	for i := 0; i < 5; i++ {
		_, err := cryptoVoteBO.AddUpVote(filterCryptoVote)
		assert.Nil(t, err, "err should be nil")
	}

	filterCryptoVote.Symbol = "KLV"
	// adiciona 500 votos para crypto
	for i := 0; i < 500; i++ {
		_, err := cryptoVoteBO.AddUpVote(filterCryptoVote)
		assert.Nil(t, err, "err should be nil")
	}

	pagSize := int64(2)
	list, err := cryptoVoteBO.SumaryAllCryptoVote(pagSize)
	assert.Nil(t, err, "err should be nil")
	tam := int64(len(list))
	assert.Equal(t, tam, pagSize, "should be equal")
}
