package interfaces

import (
	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"
)

/*
	In go, there is no explicit declaration (keyword) to set an inheritance relationship between classes,
	as occurs in Java or Python, for example. There is the keyword interface, but this only indicates that
	it is an abstract struct. A class that wants to be an implementation of an interface-type struct needs
	to implement all the methods of the given interface.

	InterfaceCryptoVoteBO has the signature of 11 methods. The class which implements this interface
	will define the business rules (BUSNESS OBJECT) for model-related actions, such as: validation, steps
	before bank-related actions. It works to prepare data before Create One, Retrieve One or Many, Update
	One, Delete One or Many, Add Vote (Down, means deslike or Up, as like) and so on...

	In this implementation there is no access to the bank. All bank-related actions are managed by a DAO
	object defined at constructor method NewCryptoVoteBO(), more details for this topic see SetCryptoVoteDAO.
*/
type InterfaceCryptoVoteBO interface {
	/*
		Only used in test environments, to set a diferent DAO that might be Mock of the CryptoVoteDAO Interface,
		to simulate interaction with the bank. At production en environments, is used the constructor method
		NewCryptoVoteBO()
	*/
	SetCryptoVoteDAO(dao InterfaceCryptoVoteDAO)

	AddDownVote(filterCryptoVote model.FilterCryptoVote) (model.CryptoVote, error)
	AddUpVote(filterCryptoVote model.FilterCryptoVote) (model.CryptoVote, error)

	CreateCryptoVote(cryptoVote model.CryptoVote) (model.CryptoVote, error)

	DeleteAllCryptoVoteByFilter(filterCryptoVote model.FilterCryptoVote) (int64, error)
	DeleteAllCryptoVote() (int64, error)

	RetrieveAllCryptoVoteByFilter(filterCryptoVote model.FilterCryptoVote) ([]model.CryptoVote, error)
	RetrieveOneCryptoVote(name string, symbol string) (model.CryptoVote, error)

	SumaryAllCryptoVote(pageSize int64) ([]model.SumaryCryptoVote, error)

	UpdateOneCryptoVoteByFilter(filterCryptoVote model.FilterCryptoVote, cryptoNewData model.CryptoVote) (model.CryptoVote, error)

	ValidateCryptoVote(crypto model.CryptoVote) (bool, error)
}
