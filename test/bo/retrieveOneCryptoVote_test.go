package test

import (
	"testing"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/bo"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/utils"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestRetrieveOneCryptoVote(t *testing.T) {
	testRetrieveOneCryptoVote0_Config(t)
	testRetrieveOneCryptoVote1_FilterByName(t)
	testRetrieveOneCryptoVote2_FilterBySymbol(t)
	testRetrieveOneCryptoVote3_FilterMissMatch(t)
	testRetrieveOneCryptoVote4_FormatArg(t)
	testRetrieveOneCryptoVote5_EmptyArgs(t)
}

/*
	0
	configurando antes do teste
	remove todos os dados
	insere 3 dados
*/
func testRetrieveOneCryptoVote0_Config(t *testing.T) {
	// cria filtro vazio para remover todos os registros
	filter := bson.M{}
	_, err := bo.DeleteAllCryptoVoteByFilter(filter)
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
	recupera todos os registros usando filter
	name = "Klever"
*/
func testRetrieveOneCryptoVote1_FilterByName(t *testing.T) {
	name := "Klever"
	symbol := ""
	crypto, err := bo.RetrieveOneCryptoVote(name, symbol)

	assert.Nil(t, err, "err should be nil")
	assert.NotNil(t, crypto.Id.IsZero(), "err should nnot be nil")
	assert.Equal(t, crypto.Name, name, "they should be equal")
}

/*
	2
	recupera todos os registros usando filter
	Symbol = "KLV"
*/
func testRetrieveOneCryptoVote2_FilterBySymbol(t *testing.T) {
	name := ""
	symbol := "KLV"
	crypto, err := bo.RetrieveOneCryptoVote(name, symbol)

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
func testRetrieveOneCryptoVote3_FilterMissMatch(t *testing.T) {
	name := "Bitcoin"
	symbol := "KLV"
	cryptoVote, err := bo.RetrieveOneCryptoVote(name, symbol)

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
func testRetrieveOneCryptoVote4_FormatArg(t *testing.T) {
	name := "BitCOIN"
	symbol := "btc"
	cryptoVote, err := bo.RetrieveOneCryptoVote(name, symbol)

	assert.Nil(t, err, "err should be nil")
	assert.False(t, cryptoVote.Id.IsZero(), " should be true")
}

/*
	5
	recupera todos os registros usando filter com args (name e symbol) empty
	erro de validação
*/
func testRetrieveOneCryptoVote5_EmptyArgs(t *testing.T) {
	name := ""
	symbol := ""
	cryptoVote, err := bo.RetrieveOneCryptoVote(name, symbol)

	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, cryptoVote.Id.IsZero(), " should be true")
}
