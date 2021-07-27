package dao

import (
	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
	Não existe uma declaração explicita (palavra chave)) para definir que um tipo (struct) implementa uma interface.
	Isso acontece quando um tipo implementa todo os metodos da interface, como no caso da definição da interface
	dao.InterfaceMongodbDAO que possui a assinatura de 6 metodos. Qualquer typo que implementar estes metodos faz
	uma implementação de dao.InterfaceMongodbDAO, a exemplo de dao.CryptoVoteDAO
*/
type InterfaceMongodbDAO interface {
	CreateCryptoVote(mongodbClient *mongo.Client, cryptoVote model.CryptoVote) (*mongo.InsertOneResult, error)

	FindOneCryptoVote(mongodbClient *mongo.Client, filter bson.M) (model.CryptoVote, error)
	FindManyCryptoVote(mongodbClient *mongo.Client, filter bson.M) ([]model.CryptoVote, error)
	FindManyCryptoVoteLimitedSortedByField(mongodbClient *mongo.Client, filter bson.M, limit int64, sortFieldName string, orderType int) ([]model.CryptoVote, error)

	UpdateOneCryptoVote(mongodbClient *mongo.Client, filter bson.M, newData bson.M) (model.CryptoVote, error)

	DeleteManyCryptoVote(mongodbClient *mongo.Client, filter bson.M) (*mongo.DeleteResult, error)
}
