package test

import (
	"strings"
	"testing"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/bo"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/utils"
	"github.com/stretchr/testify/assert"
)

func TestAddUpVote(t *testing.T) {
	/*
		Instância que permite acessar os metodos implementados em bo.CryptoVoteBO
	*/
	var cryptoVoteBO bo.InterfaceCryptoVoteBO = bo.CryptoVoteBO{}

	testAddUpVote0_Config(t, cryptoVoteBO)
	testAddUpVote1_FilterByName(t, cryptoVoteBO)
	testAddUpVote2_FilterBySymbol(t, cryptoVoteBO)
	testAddUpVote3_FilterMissMatch(t, cryptoVoteBO)
	testAddUpVote4_FormatArg(t, cryptoVoteBO)
	testAddUpVote5_EmptyArg(t, cryptoVoteBO)
	testAddUpVote6_SumaryValid(t, cryptoVoteBO)
	testAddUpVote7_SumaryWrong(t, cryptoVoteBO)
}

/*
	0
	configurando antes do teste
	remove todos os dados
	insere 3 dados
*/
func testAddUpVote0_Config(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
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
	adiciona 1 voto em Upvote usando filtro de name
	confirma voto
*/
func testAddUpVote1_FilterByName(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	name := "Klever"
	symbol := ""
	beforeCrypto, err := cryptoVoteBO.RetrieveOneCryptoVote(name, symbol)
	assert.Nil(t, err, "err should be nil")
	assert.NotNil(t, beforeCrypto.Symbol, "should not be nil")

	// popular os args no padrão para criar o filtro
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   strings.Title(strings.ToLower(strings.TrimSpace(name))),
		Symbol: strings.ToUpper(strings.TrimSpace(symbol)),
	}

	cryptoVote, err := cryptoVoteBO.AddUpVote(filterCryptoVote)
	assert.Nil(t, err, "err should be nil")
	assert.NotNil(t, cryptoVote.Symbol, "should not be nil")

	assert.Equal(t, cryptoVote.Name, beforeCrypto.Name, "they should be equal")

	newQtd := beforeCrypto.Qtd_Upvote + 1
	assert.Equal(t, cryptoVote.Qtd_Upvote, newQtd, "they should be equal")
}

/*
	2
	adiciona 1 voto em Upvote usando filtro de symbol
	confirma voto
*/
func testAddUpVote2_FilterBySymbol(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	name := ""
	symbol := "KLV"
	beforeCrypto, err := cryptoVoteBO.RetrieveOneCryptoVote(name, symbol)
	assert.Nil(t, err, "err should be nil")
	assert.NotNil(t, beforeCrypto.Name, "should not be nil")

	// popular os args no padrão para criar o filtro
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   strings.Title(strings.ToLower(strings.TrimSpace(name))),
		Symbol: strings.ToUpper(strings.TrimSpace(symbol)),
	}

	cryptoVote, err := cryptoVoteBO.AddUpVote(filterCryptoVote)
	assert.Nil(t, err, "err should be nil")
	assert.NotNil(t, cryptoVote.Name, "should not be nil")

	assert.Equal(t, cryptoVote.Symbol, beforeCrypto.Symbol, "they should be equal")

	newQtd := beforeCrypto.Qtd_Upvote + 1
	assert.Equal(t, cryptoVote.Qtd_Upvote, newQtd, "they should be equal")
}

/*
	3
	adiciona 1 voto em Upvote usando filter Miss Match
	name = "Bitcoin"
	e
	symbol := "KLV"
	erro nenhum documento localizado
*/
func testAddUpVote3_FilterMissMatch(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	name := "Bitcoin"
	symbol := "KLV"
	beforeCrypto, err := cryptoVoteBO.RetrieveOneCryptoVote(name, symbol)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, beforeCrypto.Id.IsZero(), " should be true")

	// popular os args no padrão para criar o filtro
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   strings.Title(strings.ToLower(strings.TrimSpace(name))),
		Symbol: strings.ToUpper(strings.TrimSpace(symbol)),
	}

	cryptoVote, err := cryptoVoteBO.AddUpVote(filterCryptoVote)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, cryptoVote.Id.IsZero(), " should be true")
}

/*
	4
	adiciona 1 voto em Upvote usando filter com parametros mau formatados
	name := "BitCOIN"
	symbol := "btc"
	encontra Bitcoin BTC
*/
func testAddUpVote4_FormatArg(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	name := "BitCOIN"
	symbol := "btc"
	beforeCrypto, err := cryptoVoteBO.RetrieveOneCryptoVote(name, symbol)
	assert.Nil(t, err, "err should be nil")
	assert.NotNil(t, beforeCrypto.Id.IsZero(), "err should not be nil")

	// popular os args no padrão para criar o filtro
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   strings.Title(strings.ToLower(strings.TrimSpace(name))),
		Symbol: strings.ToUpper(strings.TrimSpace(symbol)),
	}

	cryptoVote, err := cryptoVoteBO.AddUpVote(filterCryptoVote)
	assert.Nil(t, err, "err should be nil")
	assert.NotNil(t, cryptoVote.Id.IsZero(), "err should not be nil")

	assert.Equal(t, beforeCrypto.Symbol, cryptoVote.Symbol, "they should be equal")

	newQtd := beforeCrypto.Qtd_Upvote + 1
	assert.Equal(t, cryptoVote.Qtd_Upvote, newQtd, "they should be equal")
}

/*
	5
	adiciona 1 voto em Upvote usando filter com args (name e symbol) empty
	erro de validação
*/
func testAddUpVote5_EmptyArg(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	name := ""
	symbol := ""
	beforeCrypto, err := cryptoVoteBO.RetrieveOneCryptoVote(name, symbol)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, beforeCrypto.Id.IsZero(), " should be true")

	// popular os args no padrão para criar o filtro
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   strings.Title(strings.ToLower(strings.TrimSpace(name))),
		Symbol: strings.ToUpper(strings.TrimSpace(symbol)),
	}

	cryptoVote, err := cryptoVoteBO.AddUpVote(filterCryptoVote)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, cryptoVote.Id.IsZero(), " should be true")
}

/*
	6
	adiciona votos alterar sumary
*/
func testAddUpVote6_SumaryValid(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	// cria model vazio que sera convertido para filtro vazio
	var filterCryptoVote = model.FilterCryptoVote{}
	filterCryptoVote.Name = "Crypto Para Teste de Sum"
	filterCryptoVote.Symbol = "CCC"

	//tenta localizar uma CCC
	cccCryptoVote, err := cryptoVoteBO.RetrieveOneCryptoVote("", filterCryptoVote.Symbol)
	if err != nil {
		cccCryptoVote = model.CryptoVote{
			Id:           [12]byte{},
			Name:         filterCryptoVote.Name,
			Symbol:       filterCryptoVote.Symbol,
			Qtd_Upvote:   0,
			Qtd_Downvote: 0,
			Sum:          0,
		}
		_, err := cryptoVoteBO.CreateCryptoVote(cccCryptoVote)
		assert.Nil(t, err, "err should be nil")
	}

	// adiciona 10 votos para crypto
	for i := 0; i < 10; i++ {
		_, err := cryptoVoteBO.AddUpVote(filterCryptoVote)
		assert.Nil(t, err, "err should be nil")
	}
	crypto, err := cryptoVoteBO.RetrieveOneCryptoVote("", filterCryptoVote.Symbol)
	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, crypto.Sum, 10, "Sum should be equal")

	// remove a crypyo
	_, err = cryptoVoteBO.DeleteAllCryptoVoteByFilter(filterCryptoVote)
	assert.Nil(t, err, "err should be nil")
}

/*
	7
	adiciona votos alterar sumary
*/
func testAddUpVote7_SumaryWrong(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	// cria model vazio que sera convertido para filtro vazio
	var filterCryptoVote = model.FilterCryptoVote{}
	filterCryptoVote.Name = "Crypto Para Teste de Sum"
	filterCryptoVote.Symbol = "CCC"

	//tenta localizar uma CCC
	cccCryptoVote, err := cryptoVoteBO.RetrieveOneCryptoVote("", filterCryptoVote.Symbol)
	if err != nil {
		cccCryptoVote = model.CryptoVote{
			Id:           [12]byte{},
			Name:         filterCryptoVote.Name,
			Symbol:       filterCryptoVote.Symbol,
			Qtd_Upvote:   0,
			Qtd_Downvote: 0,
			Sum:          0,
		}
		_, err := cryptoVoteBO.CreateCryptoVote(cccCryptoVote)
		assert.Nil(t, err, "err should be nil")
	}

	// adiciona 5 votos para crypto
	for i := 0; i < 5; i++ {
		_, err := cryptoVoteBO.AddUpVote(filterCryptoVote)
		assert.Nil(t, err, "err should be nil")
	}
	crypto, err := cryptoVoteBO.RetrieveOneCryptoVote("", filterCryptoVote.Symbol)
	assert.Nil(t, err, "err should be nil")
	assert.NotEqual(t, crypto.Sum, 10, "Sum should not be equal")

	// remove a crypyo
	_, err = cryptoVoteBO.DeleteAllCryptoVoteByFilter(filterCryptoVote)
	assert.Nil(t, err, "err should be nil")
}
