package test

import (
	"strings"
	"testing"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/bo"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/utils"
	"github.com/stretchr/testify/assert"
)

func TestAddDownVote(t *testing.T) {
	testAddDownVote0_Config(t)
	testAddDownVote1_FilterByName(t)
	testAddDownVote2_FilterBySymbol(t)
	testAddDownVote3_FilterMissMatch(t)
	testAddDownVote4_FormatArgs(t)
	testAddDownVote5_EmptyArg(t)
}

/*
	0
	configurando antes do teste
	remove todos os dados
	insere 3 dados
*/
func testAddDownVote0_Config(t *testing.T) {
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
	adiciona 1 voto em Downvote usando filtro de name
	confirma voto
*/
func testAddDownVote1_FilterByName(t *testing.T) {
	name := "Klever"
	symbol := ""

	beforeCrypto, err := bo.RetrieveOneCryptoVote(name, symbol)
	assert.Nil(t, err, "err should be nil")
	assert.NotNil(t, beforeCrypto.Symbol, "should not be nil")

	// popular os args no padrão para criar o filtro
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   strings.Title(strings.ToLower(strings.TrimSpace(name))),
		Symbol: strings.ToUpper(strings.TrimSpace(symbol)),
	}

	cryptoVote, err := bo.AddDownVote(filterCryptoVote)
	assert.Nil(t, err, "err should be nil")
	assert.NotNil(t, cryptoVote.Symbol, "should not be nil")

	assert.Equal(t, cryptoVote.Name, beforeCrypto.Name, "they should be equal")

	newQtd := beforeCrypto.Qtd_Downvote + 1
	assert.Equal(t, cryptoVote.Qtd_Downvote, newQtd, "they should be equal")
}

/*
	2
	adiciona 1 voto em Downvote usando filtro de symbol
	confirma voto
*/
func testAddDownVote2_FilterBySymbol(t *testing.T) {
	name := ""
	symbol := "KLV"
	beforeCrypto, err := bo.RetrieveOneCryptoVote(name, symbol)
	assert.Nil(t, err, "err should be nil")
	assert.NotNil(t, beforeCrypto.Name, "should not be nil")

	// popular os args no padrão para criar o filtro
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   strings.Title(strings.ToLower(strings.TrimSpace(name))),
		Symbol: strings.ToUpper(strings.TrimSpace(symbol)),
	}

	cryptoVote, err := bo.AddDownVote(filterCryptoVote)
	assert.Nil(t, err, "err should be nil")
	assert.NotNil(t, cryptoVote.Name, "should not be nil")

	assert.Equal(t, cryptoVote.Symbol, beforeCrypto.Symbol, "they should be equal")

	newQtd := beforeCrypto.Qtd_Downvote + 1
	assert.Equal(t, cryptoVote.Qtd_Downvote, newQtd, "they should be equal")
}

/*
	3
	adiciona 1 voto em Downvote usando filter Miss Match
	name = "Bitcoin"
	e
	symbol := "KLV"
	erro nenhum documento localizado
*/
func testAddDownVote3_FilterMissMatch(t *testing.T) {
	name := "Bitcoin"
	symbol := "KLV"
	beforeCrypto, err := bo.RetrieveOneCryptoVote(name, symbol)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, beforeCrypto.Id.IsZero(), " should be true")

	// popular os args no padrão para criar o filtro
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   strings.Title(strings.ToLower(strings.TrimSpace(name))),
		Symbol: strings.ToUpper(strings.TrimSpace(symbol)),
	}

	cryptoVote, err := bo.AddDownVote(filterCryptoVote)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, cryptoVote.Id.IsZero(), " should be true")
}

/*
	4
	adiciona 1 voto em Downvote usando filter com parametros mau formatados
	name := "BitCOIN"
	symbol := "btc"
	encontra Bitcoin BTC
*/
func testAddDownVote4_FormatArgs(t *testing.T) {
	name := "BitCOIN"
	symbol := "btc"
	beforeCrypto, err := bo.RetrieveOneCryptoVote(name, symbol)
	assert.Nil(t, err, "err should be nil")
	assert.NotNil(t, beforeCrypto.Id.IsZero(), "err should not be nil")

	// popular os args no padrão para criar o filtro
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   strings.Title(strings.ToLower(strings.TrimSpace(name))),
		Symbol: strings.ToUpper(strings.TrimSpace(symbol)),
	}

	cryptoVote, err := bo.AddDownVote(filterCryptoVote)
	assert.Nil(t, err, "err should be nil")
	assert.NotNil(t, cryptoVote.Id.IsZero(), "err should not be nil")

	assert.Equal(t, cryptoVote.Symbol, beforeCrypto.Symbol, "they should be equal")

	newQtd := beforeCrypto.Qtd_Downvote + 1
	assert.Equal(t, cryptoVote.Qtd_Downvote, newQtd, "they should be equal")
}

/*
	5
	adiciona 1 voto em Downvote usando filter com args (name e symbol) empty
	erro de validação
*/
func testAddDownVote5_EmptyArg(t *testing.T) {
	name := ""
	symbol := ""
	beforeCrypto, err := bo.RetrieveOneCryptoVote(name, symbol)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, beforeCrypto.Id.IsZero(), " should be true")

	// popular os args no padrão para criar o filtro
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   strings.Title(strings.ToLower(strings.TrimSpace(name))),
		Symbol: strings.ToUpper(strings.TrimSpace(symbol)),
	}

	cryptoVote, err := bo.AddDownVote(filterCryptoVote)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, cryptoVote.Id.IsZero(), " should be true")
}
