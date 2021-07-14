package dao

import (
	"context"
	"log"
	"time"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	DB_NAME         = "teste"
	COLLECTION_NAME = "cryptovotes"
)

// variável usada para recuperar um registro do tipo model.CryptoVote
var objCryptoVote model.CryptoVote

// variável usada para recuperar uma array de registro do tipo model.CryptoVote
var sliceObjCryptoVote []model.CryptoVote

/*
	Função para criar um registro
	usa na entrada um model.CryptoVote strcut criado em outra camada
*/
func CreateCryptoVote(mongodbClient *mongo.Client, CryptoVote model.CryptoVote) (*mongo.InsertOneResult, error) {
	//criar um contexto com deadline de 10 segundos
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cryptoVoteCollection := mongodbClient.Database(DB_NAME).Collection(COLLECTION_NAME)

	insertResult, err := cryptoVoteCollection.InsertOne(mongoContext, CryptoVote)
	return insertResult, err
}

//Função para deletar
/*
	Função para deletar varios registros
	c
*/
func DeleteManyCryptoVote(mongodbClient *mongo.Client, filter bson.M) (*mongo.DeleteResult, error) {
	//criar um contexto com deadline de 10 segundos
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cryptoVoteCollection := mongodbClient.Database(DB_NAME).Collection(COLLECTION_NAME)

	deleteResult, err := cryptoVoteCollection.DeleteMany(mongoContext, filter)
	return deleteResult, err
}

//Função para buscar
/*
	Função para buscar um model.CryptoVote
	usa na entrada filter := bson.M{"key": "value"} criado em outra camada
*/
func FindOneCryptoVoteByIdHex(mongodbClient *mongo.Client, idHex string) (model.CryptoVote, error) {
	//criar um contexto com deadline de 10 segundos
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cryptoVoteCollection := mongodbClient.Database(DB_NAME).Collection(COLLECTION_NAME)

	// cria os parametros do filtro sem restrições
	primitiveObjectID, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": primitiveObjectID}

	err = cryptoVoteCollection.FindOne(mongoContext, filter).Decode(&objCryptoVote)
	return objCryptoVote, err
}

/*
	Função para buscar um model.CryptoVote
	usa na entrada filter := bson.M{"key": "value"} criado em outra camada
*/
func FindOneCryptoVote(mongodbClient *mongo.Client, filter bson.M) (model.CryptoVote, error) {
	//criar um contexto com deadline de 10 segundos
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cryptoVoteCollection := mongodbClient.Database(DB_NAME).Collection(COLLECTION_NAME)

	err := cryptoVoteCollection.FindOne(mongoContext, filter).Decode(&objCryptoVote)
	return objCryptoVote, err
}

/*
	Função para buscar vários registros de model.CryptoVote
	usa na entrada filter := bson.M{"key": "value"} criado em outra camada
*/
func FindManyCryptoVote(mongodbClient *mongo.Client, filter bson.M) ([]model.CryptoVote, error) {
	//criar um contexto com deadline de 10 segundos
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cryptoVoteCollection := mongodbClient.Database(DB_NAME).Collection(COLLECTION_NAME)

	cursor, err := cryptoVoteCollection.Find(mongoContext, filter)

	// chamada ao banco cryptoVoteCollection.Find(mongoContext, filter) tem erro
	if err != nil {
		defer cancel()

	} else {
		// se a chamada ao banco estiver ok
		sliceObjCryptoVote = nil
		for cursor.Next(mongoContext) {
			err = cursor.Decode(&objCryptoVote)
			objCryptoVote.IdHex = objCryptoVote.Id.Hex()
			sliceObjCryptoVote = append(sliceObjCryptoVote, objCryptoVote)
		}
	}
	return sliceObjCryptoVote, err
}

//Função para atualizar
/*
	Função para deletar um registro
	usa na entrada filter := bson.M{"last_name": "silva"}

	usa na entrada newData := bson.M{"age": "93"}
*/
func UpdateOneCryptoVote(mongodbClient *mongo.Client, filter bson.M, newData bson.M) (model.CryptoVote, error) {
	//criar um contexto com deadline de 10 segundos
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cryptoVoteCollection := mongodbClient.Database(DB_NAME).Collection(COLLECTION_NAME)

	//se o documento não exisitr não faz nada
	//se alterar para true cria um novo documento caso não seja localizado
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	updateResult := cryptoVoteCollection.FindOneAndUpdate(mongoContext, filter, newData, &opt)
	err := updateResult.Decode(&objCryptoVote)
	return objCryptoVote, err
}
