package test

import (
	"testing"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/bo"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/utils"
	"github.com/stretchr/testify/assert"
)

func TestUpdateOneCryptoVoteByFilter(t *testing.T) {
	/*
		Instância que permite acessar os metodos implementados em bo.CryptoVoteBO
	*/
	var cryptoVoteBO bo.InterfaceCryptoVoteBO = bo.CryptoVoteBO{}

	testUpdateOneCryptoVoteByFilter0_Config(t, cryptoVoteBO)
	testUpdateOneCryptoVoteByFilter1_FilterByNameValidData(t, cryptoVoteBO)
	testUpdateOneCryptoVoteByFilter2_FilterBySymbolValidData(t, cryptoVoteBO)
	testUpdateOneCryptoVoteByFilter3_FilterMissMatch(t, cryptoVoteBO)
	testUpdateOneCryptoVoteByFilter4_FilterArgsEmpty(t, cryptoVoteBO)
	testUpdateOneCryptoVoteByFilter5_FilterNull(t, cryptoVoteBO)
	testUpdateOneCryptoVoteByFilter6_DataEmptyArgs(t, cryptoVoteBO)
	testUpdateOneCryptoVoteByFilter7_DataNull(t, cryptoVoteBO)
	testUpdateOneCryptoVoteByFilter8_DuplicatedData(t, cryptoVoteBO)
	testUpdateOneCryptoVoteByFilter9_MissingNameData(t, cryptoVoteBO)
	testUpdateOneCryptoVoteByFilter10_MissingSymbolData(t, cryptoVoteBO)
	testUpdateOneCryptoVoteByFilter11_SymbolTolargeData(t, cryptoVoteBO)
	testUpdateOneCryptoVoteByFilter12_NameTolargeData(t, cryptoVoteBO)
	testUpdateOneCryptoVoteByFilter13_NotUniqueSymbolData(t, cryptoVoteBO)
	testUpdateOneCryptoVoteByFilter14_NotUniqueNameData(t, cryptoVoteBO)
	testUpdateOneCryptoVoteByFilter15_UpvoteNegativeData(t, cryptoVoteBO)
	testUpdateOneCryptoVoteByFilter16_DownvoteNegativeData(t, cryptoVoteBO)
}

/*
	0
	configurando antes do teste
	remove todos os dados
	insere 3 dados
*/
func testUpdateOneCryptoVoteByFilter0_Config(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
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
	atualiza usando filter by name
	faz a atualização
*/
func testUpdateOneCryptoVoteByFilter1_FilterByNameValidData(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   "Klever",
		Symbol: "",
	}

	newCryptoData := model.CryptoVote{
		Name:   "New Klever",
		Symbol: "NKLV",
	}

	updatedCryptoVote, err := cryptoVoteBO.UpdateOneCryptoVoteByFilter(filterCryptoVote, newCryptoData)
	assert.Nil(t, err, "err should be nil")
	assert.False(t, updatedCryptoVote.Id.IsZero(), "should be false")
}

/*
	2
	atualiza usando filter by symbol
	faz a atualização
*/
func testUpdateOneCryptoVoteByFilter2_FilterBySymbolValidData(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	// criar um filtro de busca para fazer atualizações
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   "",
		Symbol: "NKLV",
	}

	newCryptoData := model.CryptoVote{
		Name:   "Klever",
		Symbol: "KLV",
	}

	updatedCryptoVote, err := cryptoVoteBO.UpdateOneCryptoVoteByFilter(filterCryptoVote, newCryptoData)
	assert.Nil(t, err, "err should be nil")
	assert.False(t, updatedCryptoVote.Id.IsZero(), "should be false")
}

/*
	3
	atualiza usando filter miss match
	não faz a atualização
*/
func testUpdateOneCryptoVoteByFilter3_FilterMissMatch(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	// criar um filtro miss match  para fazer atualizações
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   "Klever",
		Symbol: "NeeeeeeeeeeeKLV",
	}

	newCryptoData := model.CryptoVote{
		Name:   "MissMatch Klever",
		Symbol: "MM KLV",
	}

	updatedCryptoVote, err := cryptoVoteBO.UpdateOneCryptoVoteByFilter(filterCryptoVote, newCryptoData)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, updatedCryptoVote.Id.IsZero(), "should be true")
}

/*
	4
	atualiza usando filter com os args (name e symbol) empty
	não faz a atualização
*/
func testUpdateOneCryptoVoteByFilter4_FilterArgsEmpty(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	// criar um filtro miss match  para fazer atualizações
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   "",
		Symbol: "",
	}

	newCryptoData := model.CryptoVote{
		Name:   "MissMatch Klever",
		Symbol: "MM KLV",
	}

	updatedCryptoVote, err := cryptoVoteBO.UpdateOneCryptoVoteByFilter(filterCryptoVote, newCryptoData)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, updatedCryptoVote.Id.IsZero(), "should be true")
}

/*
	5
	atualiza usando filter nul
	não faz a atualização
*/
func testUpdateOneCryptoVoteByFilter5_FilterNull(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	// criar um filtro de busca para fazer atualizações
	var filterCryptoVote model.FilterCryptoVote

	newCryptoData := model.CryptoVote{
		Name:   "MissMatch Klever",
		Symbol: "MM KLV",
	}

	updatedCryptoVote, err := cryptoVoteBO.UpdateOneCryptoVoteByFilter(filterCryptoVote, newCryptoData)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, updatedCryptoVote.Id.IsZero(), "should be true")
}

/*
	6
	atualiza usando data (name e symbol) empty
	não faz a atualização
*/
func testUpdateOneCryptoVoteByFilter6_DataEmptyArgs(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	// criar um filtro válido de busca para fazer atualizações
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   "",
		Symbol: "btc",
	}

	newCryptoData := model.CryptoVote{
		Name:   "",
		Symbol: "",
	}

	updatedCryptoVote, err := cryptoVoteBO.UpdateOneCryptoVoteByFilter(filterCryptoVote, newCryptoData)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, updatedCryptoVote.Id.IsZero(), "should be true")
}

/*
	7
	atualiza usando null data
	não faz a atualização
*/
func testUpdateOneCryptoVoteByFilter7_DataNull(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	// criar um filtro válido de busca para fazer atualizações
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   "",
		Symbol: "btc",
	}

	// nil cryptoData sem nada
	var nilCryptoData model.CryptoVote

	updatedCryptoVote, err := cryptoVoteBO.UpdateOneCryptoVoteByFilter(filterCryptoVote, nilCryptoData)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, updatedCryptoVote.Id.IsZero(), "should be true")
}

/*
	8
	atualiza usando null dados duplicados
	não faz a atualização
*/
func testUpdateOneCryptoVoteByFilter8_DuplicatedData(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	// criar um filtro válido de busca para fazer atualizações
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   "Bitcoin",
		Symbol: "",
	}

	// carrega json data com 3 CrypytoVotes
	listIn, err := utils.Load_data(utils.JsonInData)
	assert.Nil(t, err, "err should be nil")

	cryptoVote := listIn[0]

	updatedCryptoVote, err := cryptoVoteBO.UpdateOneCryptoVoteByFilter(filterCryptoVote, cryptoVote)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, updatedCryptoVote.Id.IsZero(), "should be true")
}

/*
	9
	atualiza com empty name CryptoVote
	não faz a atualização
*/
func testUpdateOneCryptoVoteByFilter9_MissingNameData(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	// criar um filtro válido de busca para fazer atualizações
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   "Bitcoin",
		Symbol: "",
	}

	newCryptoData := model.CryptoVote{
		Name:         "",
		Symbol:       "FORM",
		Qtd_Upvote:   0,
		Qtd_Downvote: 0,
	}

	updatedCryptoVote, err := cryptoVoteBO.UpdateOneCryptoVoteByFilter(filterCryptoVote, newCryptoData)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, updatedCryptoVote.Id.IsZero(), "should be true")
}

/*
	10
	atualiza com empty symbol CryptoVote
	não faz a atualização
*/
func testUpdateOneCryptoVoteByFilter10_MissingSymbolData(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	// criar um filtro válido de busca para fazer atualizações
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   "Bitcoin",
		Symbol: "",
	}

	newCryptoData := model.CryptoVote{
		Name:         "FormiCOIN",
		Symbol:       "",
		Qtd_Upvote:   0,
		Qtd_Downvote: 0,
	}

	updatedCryptoVote, err := cryptoVoteBO.UpdateOneCryptoVoteByFilter(filterCryptoVote, newCryptoData)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, updatedCryptoVote.Id.IsZero(), "should be true")
}

/*
	11
	atualiza com campo symbol > 6 CryptoVote
	não faz a atualização
*/
func testUpdateOneCryptoVoteByFilter11_SymbolTolargeData(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	// criar um filtro válido de busca para fazer atualizações
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   "Bitcoin",
		Symbol: "",
	}

	newCryptoData := model.CryptoVote{
		Name:         "FormiCOIN",
		Symbol:       "FORMFORMFORMFORM",
		Qtd_Upvote:   0,
		Qtd_Downvote: 0,
	}

	updatedCryptoVote, err := cryptoVoteBO.UpdateOneCryptoVoteByFilter(filterCryptoVote, newCryptoData)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, updatedCryptoVote.Id.IsZero(), "should be true")
}

/*
	12
	atualiza com campo name > 30 CryptoVote
	não faz a atualização
*/
func testUpdateOneCryptoVoteByFilter12_NameTolargeData(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	// criar um filtro válido de busca para fazer atualizações
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   "Bitcoin",
		Symbol: "",
	}

	newCryptoData := model.CryptoVote{
		Name:         "Formitcoinhjauheauhuehuahueuauehuahuheuahuehua",
		Symbol:       "FORM",
		Qtd_Upvote:   0,
		Qtd_Downvote: 0,
	}

	updatedCryptoVote, err := cryptoVoteBO.UpdateOneCryptoVoteByFilter(filterCryptoVote, newCryptoData)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, updatedCryptoVote.Id.IsZero(), "should be true")
}

/*
	13
	atualiza com os dados de symbol not unique
	não faz a atualização
*/
func testUpdateOneCryptoVoteByFilter13_NotUniqueSymbolData(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	// criar um filtro válido de busca para fazer atualizações
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   "Bitcoin",
		Symbol: "",
	}

	// carrega json data com 3 CrypytoVotes
	listIn, err := utils.Load_data(utils.JsonInData)
	assert.Nil(t, err, "err should be nil")

	newCryptoData := listIn[0]
	newCryptoData.Name = "Cacau Coin"

	updatedCryptoVote, err := cryptoVoteBO.UpdateOneCryptoVoteByFilter(filterCryptoVote, newCryptoData)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, updatedCryptoVote.Id.IsZero(), "should be true")
}

/*
	14
	atualiza com os dados de name not unique
	não faz a atualização
*/
func testUpdateOneCryptoVoteByFilter14_NotUniqueNameData(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	// criar um filtro válido de busca para fazer atualizações
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   "Bitcoin",
		Symbol: "",
	}

	// carrega json data com 3 CrypytoVotes
	listIn, err := utils.Load_data(utils.JsonInData)
	assert.Nil(t, err, "err should be nil")

	newCryptoData := listIn[0]
	newCryptoData.Symbol = "CC"

	updatedCryptoVote, err := cryptoVoteBO.UpdateOneCryptoVoteByFilter(filterCryptoVote, newCryptoData)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, updatedCryptoVote.Id.IsZero(), "should be true")
}

/*
	15
	atualiza com os dados de Qtd_Upvote menor que zero
	não faz a atualização
*/
func testUpdateOneCryptoVoteByFilter15_UpvoteNegativeData(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	// criar um filtro válido de busca para fazer atualizações
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   "Bitcoin",
		Symbol: "",
	}

	newCryptoData := model.CryptoVote{
		Name:         "Cacau Coin",
		Symbol:       "CC",
		Qtd_Upvote:   -1,
		Qtd_Downvote: 0,
	}

	updatedCryptoVote, err := cryptoVoteBO.UpdateOneCryptoVoteByFilter(filterCryptoVote, newCryptoData)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, updatedCryptoVote.Id.IsZero(), "should be true")
}

/*
	16
	atualiza com os dados de Qtd_Downvote menor que zero
	não faz a atualização
*/
func testUpdateOneCryptoVoteByFilter16_DownvoteNegativeData(t *testing.T, cryptoVoteBO bo.InterfaceCryptoVoteBO) {
	// criar um filtro válido de busca para fazer atualizações
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   "Bitcoin",
		Symbol: "",
	}

	newCryptoData := model.CryptoVote{
		Name:         "Cacau Coin",
		Symbol:       "CC",
		Qtd_Upvote:   0,
		Qtd_Downvote: -1,
	}

	updatedCryptoVote, err := cryptoVoteBO.UpdateOneCryptoVoteByFilter(filterCryptoVote, newCryptoData)
	assert.NotNil(t, err, "err should not be nil")
	assert.True(t, updatedCryptoVote.Id.IsZero(), "should be true")
}
