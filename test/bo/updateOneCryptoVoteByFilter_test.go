package test

import (
	"testing"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/bo"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/utils"
	"github.com/stretchr/testify/assert"
)

func TestUpdateOneCryptoVoteByFilter(t *testing.T) {
	testUpdateOneCryptoVoteByFilter0_Config(t)
	testUpdateOneCryptoVoteByFilter1_FilterByNameValidData(t)
	testUpdateOneCryptoVoteByFilter2_FilterBySymbolValidData(t)
	testUpdateOneCryptoVoteByFilter3_FilterMissMatch(t)
	testUpdateOneCryptoVoteByFilter4_FilterArgsEmpty(t)
	testUpdateOneCryptoVoteByFilter5_FilterNull(t)
	testUpdateOneCryptoVoteByFilter6_DataEmptyArgs(t)
	testUpdateOneCryptoVoteByFilter7_DataNull(t)
	testUpdateOneCryptoVoteByFilter8_DuplicatedData(t)
	testUpdateOneCryptoVoteByFilter9_MissingNameData(t)
	testUpdateOneCryptoVoteByFilter10_MissingSymbolData(t)
	testUpdateOneCryptoVoteByFilter11_SymbolTolargeData(t)
	testUpdateOneCryptoVoteByFilter12_NameTolargeData(t)
	testUpdateOneCryptoVoteByFilter13_NotUniqueSymbolData(t)
	testUpdateOneCryptoVoteByFilter14_NotUniqueNameData(t)
	testUpdateOneCryptoVoteByFilter15_UpvoteNegativeData(t)
	testUpdateOneCryptoVoteByFilter16_DownvoteNegativeData(t)
}

/*
	0
	configurando antes do teste
	remove todos os dados
	insere 3 dados
*/
func testUpdateOneCryptoVoteByFilter0_Config(t *testing.T) {
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
	atualiza usando filter by name
	faz a atualização
*/
func testUpdateOneCryptoVoteByFilter1_FilterByNameValidData(t *testing.T) {
	// cria model vazio que sera convertido para filtro vazio
	var filterCryptoVote = model.CryptoVote{
		Name:   "Klever",
		Symbol: "",
	}

	newCryptoData := model.CryptoVote{
		Name:   "New Klever",
		Symbol: "NKLV",
	}

	updatedCryptoVote, err := bo.UpdateOneCryptoVoteByFilter(filterCryptoVote, newCryptoData)
	assert.Nil(t, err, "err should be nil")
	assert.False(t, updatedCryptoVote.Id.IsZero(), "should be false")
}

/*
	2
	atualiza usando filter by symbol
	faz a atualização
*/
func testUpdateOneCryptoVoteByFilter2_FilterBySymbolValidData(t *testing.T) {
	// criar um filtro de busca para fazer atualizações
	var filterCryptoVote = model.CryptoVote{
		Name:   "",
		Symbol: "NKLV",
	}

	newCryptoData := model.CryptoVote{
		Name:   "Klever",
		Symbol: "KLV",
	}

	updatedCryptoVote, err := bo.UpdateOneCryptoVoteByFilter(filterCryptoVote, newCryptoData)
	assert.Nil(t, err, "err should be nil")
	assert.False(t, updatedCryptoVote.Id.IsZero(), "should be false")
}

/*
	3
	atualiza usando filter miss match
	não faz a atualização
*/
func testUpdateOneCryptoVoteByFilter3_FilterMissMatch(t *testing.T) {
	// criar um filtro miss match  para fazer atualizações
	var filterCryptoVote = model.CryptoVote{
		Name:   "Klever",
		Symbol: "NeeeeeeeeeeeKLV",
	}

	newCryptoData := model.CryptoVote{
		Name:   "MissMatch Klever",
		Symbol: "MM KLV",
	}

	updatedCryptoVote, err := bo.UpdateOneCryptoVoteByFilter(filterCryptoVote, newCryptoData)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, updatedCryptoVote.Id.IsZero(), "should be true")
}

/*
	4
	atualiza usando filter com os args (name e symbol) empty
	não faz a atualização
*/
func testUpdateOneCryptoVoteByFilter4_FilterArgsEmpty(t *testing.T) {
	// criar um filtro miss match  para fazer atualizações
	var filterCryptoVote = model.CryptoVote{
		Name:   "",
		Symbol: "",
	}

	newCryptoData := model.CryptoVote{
		Name:   "MissMatch Klever",
		Symbol: "MM KLV",
	}

	updatedCryptoVote, err := bo.UpdateOneCryptoVoteByFilter(filterCryptoVote, newCryptoData)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, updatedCryptoVote.Id.IsZero(), "should be true")
}

/*
	5
	atualiza usando filter nul
	não faz a atualização
*/
func testUpdateOneCryptoVoteByFilter5_FilterNull(t *testing.T) {
	// criar um filtro de busca para fazer atualizações
	var filterCryptoVote model.CryptoVote

	newCryptoData := model.CryptoVote{
		Name:   "MissMatch Klever",
		Symbol: "MM KLV",
	}

	updatedCryptoVote, err := bo.UpdateOneCryptoVoteByFilter(filterCryptoVote, newCryptoData)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, updatedCryptoVote.Id.IsZero(), "should be true")
}

/*
	6
	atualiza usando data (name e symbol) empty
	não faz a atualização
*/
func testUpdateOneCryptoVoteByFilter6_DataEmptyArgs(t *testing.T) {
	// criar um filtro válido de busca para fazer atualizações
	var filterCryptoVote = model.CryptoVote{
		Name:   "",
		Symbol: "btc",
	}

	newCryptoData := model.CryptoVote{
		Name:   "",
		Symbol: "",
	}

	updatedCryptoVote, err := bo.UpdateOneCryptoVoteByFilter(filterCryptoVote, newCryptoData)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, updatedCryptoVote.Id.IsZero(), "should be true")
}

/*
	7
	atualiza usando null data
	não faz a atualização
*/
func testUpdateOneCryptoVoteByFilter7_DataNull(t *testing.T) {
	// criar um filtro válido de busca para fazer atualizações
	var filterCryptoVote = model.CryptoVote{
		Name:   "",
		Symbol: "btc",
	}

	// nil cryptoData sem nada
	var nilCryptoData model.CryptoVote

	updatedCryptoVote, err := bo.UpdateOneCryptoVoteByFilter(filterCryptoVote, nilCryptoData)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, updatedCryptoVote.Id.IsZero(), "should be true")
}

/*
	8
	atualiza usando null dados duplicados
	não faz a atualização
*/
func testUpdateOneCryptoVoteByFilter8_DuplicatedData(t *testing.T) {
	// criar um filtro válido de busca para fazer atualizações
	var filterCryptoVote = model.CryptoVote{
		Name:   "Bitcoin",
		Symbol: "",
	}

	// carrega json data com 3 CrypytoVotes
	listIn, err := utils.Load_data(utils.JsonInData)
	assert.Nil(t, err, "err should be nil")

	cryptoVote := listIn[0]

	updatedCryptoVote, err := bo.UpdateOneCryptoVoteByFilter(filterCryptoVote, cryptoVote)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, updatedCryptoVote.Id.IsZero(), "should be true")
}

/*
	9
	atualiza com empty name CryptoVote
	não faz a atualização
*/
func testUpdateOneCryptoVoteByFilter9_MissingNameData(t *testing.T) {
	// criar um filtro válido de busca para fazer atualizações
	var filterCryptoVote = model.CryptoVote{
		Name:   "Bitcoin",
		Symbol: "",
	}

	newCryptoData := model.CryptoVote{
		Name:         "",
		Symbol:       "FORM",
		Qtd_Upvote:   0,
		Qtd_Downvote: 0,
	}

	updatedCryptoVote, err := bo.UpdateOneCryptoVoteByFilter(filterCryptoVote, newCryptoData)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, updatedCryptoVote.Id.IsZero(), "should be true")
}

/*
	10
	atualiza com empty symbol CryptoVote
	não faz a atualização
*/
func testUpdateOneCryptoVoteByFilter10_MissingSymbolData(t *testing.T) {
	// criar um filtro válido de busca para fazer atualizações
	var filterCryptoVote = model.CryptoVote{
		Name:   "Bitcoin",
		Symbol: "",
	}

	newCryptoData := model.CryptoVote{
		Name:         "FormiCOIN",
		Symbol:       "",
		Qtd_Upvote:   0,
		Qtd_Downvote: 0,
	}

	updatedCryptoVote, err := bo.UpdateOneCryptoVoteByFilter(filterCryptoVote, newCryptoData)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, updatedCryptoVote.Id.IsZero(), "should be true")
}

/*
	11
	atualiza com campo symbol > 6 CryptoVote
	não faz a atualização
*/
func testUpdateOneCryptoVoteByFilter11_SymbolTolargeData(t *testing.T) {
	// criar um filtro válido de busca para fazer atualizações
	var filterCryptoVote = model.CryptoVote{
		Name:   "Bitcoin",
		Symbol: "",
	}

	newCryptoData := model.CryptoVote{
		Name:         "FormiCOIN",
		Symbol:       "FORMFORMFORMFORM",
		Qtd_Upvote:   0,
		Qtd_Downvote: 0,
	}

	updatedCryptoVote, err := bo.UpdateOneCryptoVoteByFilter(filterCryptoVote, newCryptoData)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, updatedCryptoVote.Id.IsZero(), "should be true")
}

/*
	12
	atualiza com campo name > 30 CryptoVote
	não faz a atualização
*/
func testUpdateOneCryptoVoteByFilter12_NameTolargeData(t *testing.T) {
	// criar um filtro válido de busca para fazer atualizações
	var filterCryptoVote = model.CryptoVote{
		Name:   "Bitcoin",
		Symbol: "",
	}

	newCryptoData := model.CryptoVote{
		Name:         "Formitcoinhjauheauhuehuahueuauehuahuheuahuehua",
		Symbol:       "FORM",
		Qtd_Upvote:   0,
		Qtd_Downvote: 0,
	}

	updatedCryptoVote, err := bo.UpdateOneCryptoVoteByFilter(filterCryptoVote, newCryptoData)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, updatedCryptoVote.Id.IsZero(), "should be true")
}

/*
	13
	atualiza com os dados de symbol not unique
	não faz a atualização
*/
func testUpdateOneCryptoVoteByFilter13_NotUniqueSymbolData(t *testing.T) {
	// criar um filtro válido de busca para fazer atualizações
	var filterCryptoVote = model.CryptoVote{
		Name:   "Bitcoin",
		Symbol: "",
	}

	// carrega json data com 3 CrypytoVotes
	listIn, err := utils.Load_data(utils.JsonInData)
	assert.Nil(t, err, "err should be nil")

	newCryptoData := listIn[0]
	newCryptoData.Name = "Cacau Coin"

	updatedCryptoVote, err := bo.UpdateOneCryptoVoteByFilter(filterCryptoVote, newCryptoData)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, updatedCryptoVote.Id.IsZero(), "should be true")
}

/*
	14
	atualiza com os dados de name not unique
	não faz a atualização
*/
func testUpdateOneCryptoVoteByFilter14_NotUniqueNameData(t *testing.T) {
	// criar um filtro válido de busca para fazer atualizações
	var filterCryptoVote = model.CryptoVote{
		Name:   "Bitcoin",
		Symbol: "",
	}

	// carrega json data com 3 CrypytoVotes
	listIn, err := utils.Load_data(utils.JsonInData)
	assert.Nil(t, err, "err should be nil")

	newCryptoData := listIn[0]
	newCryptoData.Symbol = "CC"

	updatedCryptoVote, err := bo.UpdateOneCryptoVoteByFilter(filterCryptoVote, newCryptoData)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, updatedCryptoVote.Id.IsZero(), "should be true")
}

/*
	15
	atualiza com os dados de Qtd_Upvote menor que zero
	não faz a atualização
*/
func testUpdateOneCryptoVoteByFilter15_UpvoteNegativeData(t *testing.T) {
	// criar um filtro válido de busca para fazer atualizações
	var filterCryptoVote = model.CryptoVote{
		Name:   "Bitcoin",
		Symbol: "",
	}

	newCryptoData := model.CryptoVote{
		Name:         "Cacau Coin",
		Symbol:       "CC",
		Qtd_Upvote:   -1,
		Qtd_Downvote: 0,
	}

	updatedCryptoVote, err := bo.UpdateOneCryptoVoteByFilter(filterCryptoVote, newCryptoData)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, updatedCryptoVote.Id.IsZero(), "should be true")
}

/*
	16
	atualiza com os dados de Qtd_Downvote menor que zero
	não faz a atualização
*/
func testUpdateOneCryptoVoteByFilter16_DownvoteNegativeData(t *testing.T) {
	// criar um filtro válido de busca para fazer atualizações
	var filterCryptoVote = model.CryptoVote{
		Name:   "Bitcoin",
		Symbol: "",
	}

	newCryptoData := model.CryptoVote{
		Name:         "Cacau Coin",
		Symbol:       "CC",
		Qtd_Upvote:   0,
		Qtd_Downvote: -1,
	}

	updatedCryptoVote, err := bo.UpdateOneCryptoVoteByFilter(filterCryptoVote, newCryptoData)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, updatedCryptoVote.Id.IsZero(), "should be true")
}
