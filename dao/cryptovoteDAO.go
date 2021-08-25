package dao

import (
	"context"
	"errors"
	"time"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/interfaces"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
)

/*
	In go, there is no explicit declaration (keyword) to set an inheritance relationship between classes,
	as occurs in Java or Python, for example. There is the keyword interface, but this only indicates that
	it is an abstract struct. A class that wants to be an implementation of an interface-type struct needs
	to implement all the methods of the given interface.

	CryptoVoteDAO implements all methods of interfaces.InterfaceCryptoVoteDAO
*/
type CryptoVoteDAO struct {
	dbService interfaces.InterfaceDbService
}

func NewCryptoVoteDAO() CryptoVoteDAO {
	var d CryptoVoteDAO = CryptoVoteDAO{
		// dao.MongodbService{} is a struct that makes the connection to the databse Mongodb
		dbService: MongodbService{},
	}
	return d
}

func (c CryptoVoteDAO) GetDbService() interfaces.InterfaceDbService {
	return c.dbService
}

func (c CryptoVoteDAO) SetDbService(service interfaces.InterfaceDbService) {
	c.dbService = service
}

/*
	Function uses a model.CryptoVote struct created at another layer and persist ir on database,
	using mongodbClient.
*/
func (c CryptoVoteDAO) Create(cryptoVote model.CryptoVote) (model.CryptoVote, error) {
	var savedCryptoVote model.CryptoVote

	// uses function from dao package to start the mongo client
	mongodbClient, err := GetMongoClientInstance(c.dbService)
	if err != nil {
		z := "[cryptovote.mongodb] Problems using GetMongoClientInstance: " + err.Error()
		err = errors.New(z)
		return savedCryptoVote, err
	}

	// before saving, do the math
	cryptoVote.Sum = cryptoVote.Qtd_Upvote - cryptoVote.Qtd_Downvote

	// create a context with 10-second deadline
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cryptoVoteCollection := mongodbClient.Database(db_name).Collection(collection_name)

	insertResult, err := cryptoVoteCollection.InsertOne(mongoContext, cryptoVote)
	if err != nil {
		return savedCryptoVote, err
	}

	// creates filter with id to retrieve data
	filter := bson.M{"_id": insertResult.InsertedID}

	// uses function from dao package to start the mongo client
	savedCryptoVote, err = c.FindOne(filter)
	if err != nil {
		z := "[cryptovote.mongodb] Problemas na execução de c.FindOne: " + err.Error()
		err = errors.New(z)
		return savedCryptoVote, err
	}
	return savedCryptoVote, err
}

/*
	Function to delete multiple records using a filter.
	filter := bson. M{"symbol": "KLV"}
*/
func (c CryptoVoteDAO) DeleteMany(filter bson.M) (int64, error) {
	var deletedCount int64 = int64(0)

	// uses function from dao package to start the mongo client
	mongodbClient, err := GetMongoClientInstance(c.dbService)
	if err != nil {
		z := "[cryptovote.mongodb] Problems using GetMongoClientInstance: " + err.Error()
		err = errors.New(z)
		return deletedCount, err
	}

	// create a context with 10-second deadline
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cryptoVoteCollection := mongodbClient.Database(db_name).Collection(collection_name)

	deleteResult, err := cryptoVoteCollection.DeleteMany(mongoContext, filter)
	if err != nil {
		z := "[cryptovote.mongodb] Problems using DeleteMany: " + err.Error()
		err = errors.New(z)
		return deletedCount, err
	}
	deletedCount = deleteResult.DeletedCount
	return deletedCount, err
}

/*
	Function to find and retrieve one record using a filter.
	filter := bson. M{"symbol": "KLV"}
*/
func (c CryptoVoteDAO) FindOne(filter bson.M) (model.CryptoVote, error) {
	var oneCryptoVote model.CryptoVote
	var err error

	// uses function from dao package to start the mongo client
	mongodbClient, err := GetMongoClientInstance(c.dbService)
	if err != nil {
		z := "[cryptovote.mongodb] Problems using GetMongoClientInstance: " + err.Error()
		err = errors.New(z)
		return oneCryptoVote, err
	}

	// create a context with 10-second deadline
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cryptoVoteCollection := mongodbClient.Database(db_name).Collection(collection_name)

	singleResult := cryptoVoteCollection.FindOne(mongoContext, filter)
	err = singleResult.Decode(&oneCryptoVote)
	return oneCryptoVote, err
}

/*
	Function to find and retrieve many record using a filter.
	filter := bson. M{"symbol": "KLV"}
*/
func (c CryptoVoteDAO) FindMany(filter bson.M) ([]model.CryptoVote, error) {
	var oneCryptoVote model.CryptoVote
	var manyCryptoVotes []model.CryptoVote

	// uses function from dao package to start the mongo client
	mongodbClient, err := GetMongoClientInstance(c.dbService)
	if err != nil {
		z := "[cryptovote.mongodb] Problems using GetMongoClientInstance: " + err.Error()
		err = errors.New(z)
		return manyCryptoVotes, err
	}

	// create a context with 10-second deadline
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cryptoVoteCollection := mongodbClient.Database(db_name).Collection(collection_name)

	cursor, err := cryptoVoteCollection.Find(mongoContext, filter)

	// previous call cryptoVoteCollection.Find(mongoContext, filter) has database errors
	if err != nil {
		defer cancel()

	} else {
		manyCryptoVotes = nil
		for cursor.Next(mongoContext) {
			err = cursor.Decode(&oneCryptoVote)
			manyCryptoVotes = append(manyCryptoVotes, oneCryptoVote)
		}
	}
	return manyCryptoVotes, err
}

/*
	Function to find and retrieve many record using filter and other options.
*/
func (c CryptoVoteDAO) FindManyLimitedSortedByField(filter bson.M, limit int64, sortFieldName string, orderType int) ([]model.CryptoVote, error) {
	var oneCryptoVote model.CryptoVote
	var manyCryptoVotes []model.CryptoVote
	var err error

	// uses function from dao package to start the mongo client
	mongodbClient, err := GetMongoClientInstance(c.dbService)
	if err != nil {
		z := "[cryptovote.mongodb] Problems using GetMongoClientInstance: " + err.Error()
		err = errors.New(z)
		return manyCryptoVotes, err
	}

	// create a context with 10-second deadline
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cryptoVoteCollection := mongodbClient.Database(db_name).Collection(collection_name)

	opt := options.Find()
	opt.SetLimit(limit)
	// Sort by field usinng orderType
	opt.SetSort(bson.D{{Key: sortFieldName, Value: orderType}})

	cursor, err := cryptoVoteCollection.Find(mongoContext, filter, opt)

	// previous call cryptoVoteCollection.Find(mongoContext, filter) has database errors
	if err != nil {
		defer cancel()

	} else {
		manyCryptoVotes = nil
		for cursor.Next(mongoContext) {
			err = cursor.Decode(&oneCryptoVote)
			manyCryptoVotes = append(manyCryptoVotes, oneCryptoVote)
		}
	}
	return manyCryptoVotes, err
}

/*
	Function to update one record using filter and newData.
	filter := bson.M{"symbol": "KLV"}
	newData := bson.M{"name": "New Data"}
*/
func (c CryptoVoteDAO) UpdateOne(filter bson.M, newData bson.M) (model.CryptoVote, error) {
	var oneCryptoVote model.CryptoVote
	var err error

	// uses function from dao package to start the mongo client
	mongodbClient, err := GetMongoClientInstance(c.dbService)
	if err != nil {
		z := "[cryptovote.mongodb] Problems using GetMongoClientInstance: " + err.Error()
		err = errors.New(z)
		return oneCryptoVote, err
	}

	// create a context with 10-second deadline
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cryptoVoteCollection := mongodbClient.Database(db_name).Collection(collection_name)

	// only updates if the document exists, if upsert := true, creates a new document if it not located
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	updateResult := cryptoVoteCollection.FindOneAndUpdate(mongoContext, filter, newData, &opt)
	err = updateResult.Decode(&oneCryptoVote)
	return oneCryptoVote, err
}