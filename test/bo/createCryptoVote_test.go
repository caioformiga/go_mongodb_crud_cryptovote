package test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/bo"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/utils"
)

func TestCreateCryptoVote(t *testing.T) {
	/*
		Instância que permite acessar os metodos implementados em bo.CryptoVoteBO
	*/
	var cryptoVoteBO bo.InterfaceCryptoVoteBO = bo.CryptoVoteBO{}

	testCreateCryptoVote0_Config(t, cryptoVoteBO)
	testCreateCryptoVote1_ValidatedData(t, cryptoVoteBO)
	testCreateCryptoVote2_DuplicatedData(t, cryptoVoteBO)
	testCreateCryptoVote3_MissingNameData(t, cryptoVoteBO)
	testCreateCryptoVote4_MissingSymbolData(t, cryptoVoteBO)
	testCreateCryptoVote5_SymbolTolargeData(t, cryptoVoteBO)
	testCreateCryptoVote6_NameTolargeData(t, cryptoVoteBO)
	testCreateCryptoVote7_NotUniqueSymbolData(t, cryptoVoteBO)
	testCreateCryptoVote8_NotUniqueNameData(t, cryptoVoteBO)
	testCreateCryptoVote9_UpvoteNegativeData(t, cryptoVoteBO)
	testCreateCryptoVote10_DownvoteNegativeData(t, cryptoVoteBO)
	testCreateCryptoVote11_BadJsonData(t, cryptoVoteBO)
}

/*
	configurando antes do teste
	limpar todos os dados
*/
func testCreateCryptoVote0_Config(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	// limpa a coleção
	_, err := cryptoVoteBO.DeleteAllCryptoVote()
	assert.Nil(t, err, "err should be nil")
}

/*
	1
	tentativa com dados validos correto
	espera que todos os registros sejam iguais a saida
*/
func testCreateCryptoVote1_ValidatedData(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	// carrega json data com 3 CrypytoVotes
	listIn, _ := utils.Load_data(utils.JsonInData)
	listOut, _ := utils.Load_data(utils.JsonOutData)
	assert.Equal(t, len(listIn), len(listOut), "they should be equal")

	tam := len(listIn)
	for i := 0; i < tam; i++ {
		savedCrypoVote, err := cryptoVoteBO.CreateCryptoVote(listIn[i])
		validCrypo := listOut[i]
		assert.Nil(t, err, "err should be nil")
		assert.False(t, savedCrypoVote.Id.IsZero(), "savedCrypoVote.Id.IsZero() should not be false")
		assert.Equal(t, savedCrypoVote.Name, validCrypo.Name, "they should be equal")
		assert.Equal(t, savedCrypoVote.Symbol, validCrypo.Symbol, "they should be equal")
		assert.Equal(t, savedCrypoVote.Qtd_Upvote, validCrypo.Qtd_Upvote, "they should be equal")
		assert.Equal(t, savedCrypoVote.Qtd_Downvote, validCrypo.Qtd_Downvote, "they should be equal")
	}
}

/*
	2
	tentativa com os mesmos dados
	espera que retorne erro de validação
*/
func testCreateCryptoVote2_DuplicatedData(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	var listIn []model.CryptoVote
	var cryptoVote model.CryptoVote

	tam := len(listIn)
	for i := 0; i < tam; i++ {
		cryptoVote = listIn[i]

		_, err := cryptoVoteBO.CreateCryptoVote(cryptoVote)
		assert.NotNil(t, err, "err should not be nil")
	}
}

/*
	3
	tentativa com empty name CryptoVote
	espera que retorne erro de validação
*/
func testCreateCryptoVote3_MissingNameData(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	cryptoVote, err := cryptoVoteBO.CreateCryptoVote(model.CryptoVote{
		Name:         "",
		Symbol:       "FORM",
		Qtd_Upvote:   0,
		Qtd_Downvote: 0,
	})
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, cryptoVote.Id.IsZero(), "cryptoVote.Id.IsZero() should be true")
}

/*
	4
	tentativa com empty symbol CryptoVote
	espera que retorne erro de validação em um campo
*/
func testCreateCryptoVote4_MissingSymbolData(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	cryptoVote, err := cryptoVoteBO.CreateCryptoVote(model.CryptoVote{
		Name:         "FormiCOIN",
		Symbol:       "",
		Qtd_Upvote:   0,
		Qtd_Downvote: 0,
	})
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, cryptoVote.Id.IsZero(), "should be true")
}

/*
	5
	tentativa com campo symbol > 6 CryptoVote
	espera que retorne erro de validação do tamanho do campo symbol
*/
func testCreateCryptoVote5_SymbolTolargeData(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	cryptoVote, err := cryptoVoteBO.CreateCryptoVote(model.CryptoVote{
		Name:         "Bitcoin",
		Symbol:       "BTCBTCBTC",
		Qtd_Upvote:   0,
		Qtd_Downvote: 0,
	})
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, cryptoVote.Id.IsZero(), "should be true")
}

/*
	6
	tentativa com campo name > 30 CryptoVote
	espera que retorne erro de validação do tamanho do campo symbol
*/
func testCreateCryptoVote6_NameTolargeData(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	cryptoVote, err := cryptoVoteBO.CreateCryptoVote(model.CryptoVote{
		Name:         "Formitcoinhjauheauhuehuahueuauehuahuheuahuehua",
		Symbol:       "FORM",
		Qtd_Upvote:   0,
		Qtd_Downvote: 0,
	})
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, cryptoVote.Id.IsZero(), "should be true")
}

/*
	7
	tentativa com os dados de symbol igual e nome unique
	espera que retorne erro de validação
*/
func testCreateCryptoVote7_NotUniqueSymbolData(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	listIn, _ := utils.Load_data(utils.JsonInData)
	cryptoVote := listIn[0]

	cryptoVote.Name = "Cacau Coin"
	_, err := cryptoVoteBO.CreateCryptoVote(cryptoVote)
	assert.NotNil(t, err, "err should not be nil")
}

/*
	8
	tentativa com os dados de nome igual e symbol unique
	espera que retorne erro de validação
*/
func testCreateCryptoVote8_NotUniqueNameData(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	listIn, _ := utils.Load_data(utils.JsonInData)
	cryptoVote := listIn[0]

	cryptoVote.Symbol = "CC"
	_, err := cryptoVoteBO.CreateCryptoVote(cryptoVote)
	assert.NotNil(t, err, "err should not be nil")
}

/*
	9
	tentativa com os dados de Qtd_Upvote menor que zero
	espera que retorne erro de validação
*/
func testCreateCryptoVote9_UpvoteNegativeData(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	var cryptoVote = model.CryptoVote{
		Id:           [12]byte{},
		Name:         "Cacau Coin",
		Symbol:       "CC",
		Qtd_Upvote:   -1,
		Qtd_Downvote: 0,
	}

	_, err := cryptoVoteBO.CreateCryptoVote(cryptoVote)
	assert.NotNil(t, err, "err should not be nil")
}

/*
	10
	tentativa com os dados de Qtd_Downvote menor que zero
	espera que retorne erro de validação
*/
func testCreateCryptoVote10_DownvoteNegativeData(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	var cryptoVote = model.CryptoVote{
		Name:         "Cacau Coin",
		Symbol:       "CC",
		Qtd_Upvote:   0,
		Qtd_Downvote: -1,
	}

	_, err := cryptoVoteBO.CreateCryptoVote(cryptoVote)
	assert.NotNil(t, err, "err should not be nil")
}

/*
	11
	tentativa com os dados malfomatados no json
	espera que retorne erro de validação
*/
func testCreateCryptoVote11_BadJsonData(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	var ptr *[]model.CryptoVote
	err := json.Unmarshal(utils.JsonBadData, &ptr)
	assert.NotNil(t, err, "err should not be nil")
}
