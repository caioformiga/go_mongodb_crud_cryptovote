package bo

import (
	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"
)

type InterfaceCryptoVoteBO interface {
	CreateCryptoVote(cryptoVote model.CryptoVote) (model.CryptoVote, error)

	RetrieveAllCryptoVoteByFilter(filterCryptoVote model.FilterCryptoVote) ([]model.CryptoVote, error)
	RetrieveOneCryptoVote(name string, symbol string) (model.CryptoVote, error)

	UpdateOneCryptoVoteByFilter(filterCryptoVote model.FilterCryptoVote, cryptoNewData model.CryptoVote) (model.CryptoVote, error)

	DeleteAllCryptoVoteByFilter(filterCryptoVote model.FilterCryptoVote) (int64, error)
	DeleteAllCryptoVote() (int64, error)

	AddUpVote(filterCryptoVote model.FilterCryptoVote) (model.CryptoVote, error)
	AddDownVote(filterCryptoVote model.FilterCryptoVote) (model.CryptoVote, error)

	SumaryAllCryptoVote(pageSize int64) ([]model.SumaryCryptoVote, error)
}
