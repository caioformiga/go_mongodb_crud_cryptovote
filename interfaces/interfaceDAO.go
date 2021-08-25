package interfaces

import (
	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"

	"go.mongodb.org/mongo-driver/bson"
)

/*
	In go, there is no explicit declaration (keyword) to set an inheritance relationship between classes,
	as occurs in Java or Python, for example. There is the keyword interface, but this only indicates that
	it is an abstract struct. A class that wants to be an implementation of an interface-type struct needs
	to implement all the methods of the given interface.

	InterfaceCryptoVoteDAO has the signature of 8 methods. The class which implements this interface will
	define the bank-related actions, using one of its 6 CRUD a model data methods. Besides this, it creates
	2 subscription methods so that you can define other database services.
*/
type InterfaceCryptoVoteDAO interface {
	GetDbService() InterfaceDbService
	SetDbService(dbService InterfaceDbService)

	Create(cryptoVote model.CryptoVote) (model.CryptoVote, error)

	FindOne(filter bson.M) (model.CryptoVote, error)
	FindMany(filter bson.M) ([]model.CryptoVote, error)
	FindManyLimitedSortedByField(filter bson.M, limit int64, sortFieldName string, orderType int) ([]model.CryptoVote, error)

	UpdateOne(filter bson.M, newData bson.M) (model.CryptoVote, error)

	DeleteMany(filter bson.M) (int64, error)
}

type InterfaceDbService interface {
	SetDatabaseName(db string)
	GetDatabaseName() string

	SetCollectionName(collection string)
	GetCollectionName() string

	SetDatabaseUri(uri string)
	GetDatabaseUri() string
}
