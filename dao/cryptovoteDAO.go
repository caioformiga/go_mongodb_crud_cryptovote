package dao

import (
	"context"
	"time"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
)

/*
	Função para criar um registro
	usa na entrada um model.CryptoVote struct criado em outra camada
*/
func CreateCryptoVote(mongodbClient *mongo.Client, cryptoVote model.CryptoVote) (*mongo.InsertOneResult, error) {
	// antes de salvar no mongo faz a soma
	cryptoVote.Sum = cryptoVote.Qtd_Upvote - cryptoVote.Qtd_Downvote

	//criar um contexto com deadline de 10 segundos
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cryptoVoteCollection := mongodbClient.Database(db_name).Collection(collection_name)

	insertResult, err := cryptoVoteCollection.InsertOne(mongoContext, cryptoVote)
	return insertResult, err
}

//Função para deletar
/*
	Função para deletar varios registros
	usa na entrada filter := bson.M{"symbol": "KLV"}
*/
func DeleteManyCryptoVote(mongodbClient *mongo.Client, filter bson.M) (*mongo.DeleteResult, error) {
	//criar um contexto com deadline de 10 segundos
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cryptoVoteCollection := mongodbClient.Database(db_name).Collection(collection_name)

	deleteResult, err := cryptoVoteCollection.DeleteMany(mongoContext, filter)
	return deleteResult, err
}

//Função para buscar
/*
	Função para recuperar vários registros de model.CryptoVote
	usa na entrada filter := bson.M{"symbol": "KLV"}
*/
func FindOneCryptoVote(mongodbClient *mongo.Client, filter bson.M) (model.CryptoVote, error) {
	var oneCryptoVote model.CryptoVote

	//criar um contexto com deadline de 10 segundos
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cryptoVoteCollection := mongodbClient.Database(db_name).Collection(collection_name)

	singleResult := cryptoVoteCollection.FindOne(mongoContext, filter)
	err := singleResult.Decode(&oneCryptoVote)
	return oneCryptoVote, err
}

/*
	Função para recuperar vários registros de model.CryptoVote
	usa na entrada filter := bson.M{"symbol": "KLV"}
*/
func FindManyCryptoVote(mongodbClient *mongo.Client, filter bson.M) ([]model.CryptoVote, error) {
	var oneCryptoVote model.CryptoVote
	var manyCryptoVotes []model.CryptoVote

	//criar um contexto com deadline de 10 segundos
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cryptoVoteCollection := mongodbClient.Database(db_name).Collection(collection_name)

	cursor, err := cryptoVoteCollection.Find(mongoContext, filter)

	// chamada ao banco cryptoVoteCollection.Find(mongoContext, filter) tem erro
	if err != nil {
		defer cancel()

	} else {
		// se a chamada ao banco estiver ok
		manyCryptoVotes = nil
		for cursor.Next(mongoContext) {
			err = cursor.Decode(&oneCryptoVote)
			manyCryptoVotes = append(manyCryptoVotes, oneCryptoVote)
		}
	}
	return manyCryptoVotes, err
}

/*
	Função para recuperar vários registros de model.CryptoVote
	usa na entrada filter := bson.M{"symbol": "KLV"}
*/
func FindManyCryptoVoteLimitedSortedByField(mongodbClient *mongo.Client, filter bson.M, limit int64, field string, orderType int) ([]model.CryptoVote, error) {
	var oneCryptoVote model.CryptoVote
	var manyCryptoVotes []model.CryptoVote

	//criar um contexto com deadline de 10 segundos
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cryptoVoteCollection := mongodbClient.Database(db_name).Collection(collection_name)

	opt := options.Find()
	opt.SetLimit(limit)
	// Sort by field usinng orderType para ascending = 1 / descending = -1
	opt.SetSort(bson.D{{Key: field, Value: orderType}})

	cursor, err := cryptoVoteCollection.Find(mongoContext, filter, opt)

	// chamada ao banco cryptoVoteCollection.Find(mongoContext, filter) tem erro
	if err != nil {
		defer cancel()

	} else {
		// se a chamada ao banco estiver ok
		manyCryptoVotes = nil
		for cursor.Next(mongoContext) {
			err = cursor.Decode(&oneCryptoVote)
			manyCryptoVotes = append(manyCryptoVotes, oneCryptoVote)
		}
	}
	return manyCryptoVotes, err
}

//Função para atualizar
/*
	Função para deletar um registro
	usa na entrada filter := bson.M{"symbol": "KLV"}

	usa na entrada newData := bson.M{"name": "New Data"}
*/
func UpdateOneCryptoVote(mongodbClient *mongo.Client, filter bson.M, newData bson.M) (model.CryptoVote, error) {
	var oneCryptoVote model.CryptoVote

	//criar um contexto com deadline de 10 segundos
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cryptoVoteCollection := mongodbClient.Database(db_name).Collection(collection_name)

	//se o documento não exisitr não faz nada
	//se alterar para true cria um novo documento caso não seja localizado
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	updateResult := cryptoVoteCollection.FindOneAndUpdate(mongoContext, filter, newData, &opt)
	err := updateResult.Decode(&oneCryptoVote)
	return oneCryptoVote, err
}
