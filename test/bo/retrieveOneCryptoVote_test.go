package test

import (
	"testing"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/bo"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/utils"
	"github.com/stretchr/testify/assert"
)

func TestRetrieveOneCryptoVote(t *testing.T) {
	/*
		Instância que permite acessar os metodos implementados em bo.CryptoVoteBO
	*/
	var cryptoVoteBO bo.InterfaceCryptoVoteBO = bo.CryptoVoteBO{}

	testRetrieveOneCryptoVote0_Config(t, cryptoVoteBO)
	testRetrieveOneCryptoVote1_FilterByName(t, cryptoVoteBO)
	testRetrieveOneCryptoVote2_FilterBySymbol(t, cryptoVoteBO)
	testRetrieveOneCryptoVote3_FilterMissMatch(t, cryptoVoteBO)
	testRetrieveOneCryptoVote4_FormatArg(t, cryptoVoteBO)
	testRetrieveOneCryptoVote5_EmptyArgs(t, cryptoVoteBO)
}

/*
	0
	configurando antes do teste
	remove todos os dados
	insere 3 dados
*/
func testRetrieveOneCryptoVote0_Config(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
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
	recupera todos os registros usando filter
	name = "Klever"
*/
func testRetrieveOneCryptoVote1_FilterByName(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	name := "Klever"
	symbol := ""
	crypto, err := cryptoVoteBO.RetrieveOneCryptoVote(name, symbol)

	assert.Nil(t, err, "err should be nil")
	assert.NotNil(t, crypto.Id.IsZero(), "err should nnot be nil")
	assert.Equal(t, crypto.Name, name, "they should be equal")
}

/*
	2
	recupera todos os registros usando filter
	Symbol = "KLV"
*/
func testRetrieveOneCryptoVote2_FilterBySymbol(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	name := ""
	symbol := "KLV"
	crypto, err := cryptoVoteBO.RetrieveOneCryptoVote(name, symbol)

	assert.Nil(t, err, "err should be nil")
	assert.NotNil(t, crypto.Id.IsZero(), "should not be nil")
	assert.Equal(t, crypto.Symbol, symbol, "they should be equal")
}

/*
	3
	recupera todos os registros usando filter Miss Match
	name = "Bitcoin"
	e
	symbol := "KLV"
	erro nenhum documento localizado
*/
func testRetrieveOneCryptoVote3_FilterMissMatch(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	name := "Bitcoin"
	symbol := "KLV"
	cryptoVote, err := cryptoVoteBO.RetrieveOneCryptoVote(name, symbol)

	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, cryptoVote.Id.IsZero(), " should be true")
}

/*
	4
	recupera todos os registros usando filter com parametros mau formatados
	name := "BitCOIN"
	symbol := "btc"
	encontra Bitcoin BTC
*/
func testRetrieveOneCryptoVote4_FormatArg(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	name := "BitCOIN"
	symbol := "btc"
	cryptoVote, err := cryptoVoteBO.RetrieveOneCryptoVote(name, symbol)

	assert.Nil(t, err, "err should be nil")
	assert.False(t, cryptoVote.Id.IsZero(), " should be true")
}

/*
	5
	recupera todos os registros usando filter com args (name e symbol) empty
	erro de validação
*/
func testRetrieveOneCryptoVote5_EmptyArgs(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	name := ""
	symbol := ""
	cryptoVote, err := cryptoVoteBO.RetrieveOneCryptoVote(name, symbol)

	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, cryptoVote.Id.IsZero(), " should be true")
}
