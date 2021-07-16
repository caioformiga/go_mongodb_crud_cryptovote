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

// variável usada para recuperar um registro do tipo model.CryptoCurrency
var objCryptoCurrency model.CryptoCurrency

// variável usada para recuperar uma array de registro do tipo model.CryptoCurrency
var sliceObjCryptoCurrency []model.CryptoCurrency

//Função para criar
/*
	Função para criar um registro
	usa na entrada um model.CryptoCurrency strcut criado em outra camada
*/
func CreateCryptoCurrency(mongodbClient *mongo.Client, CryptoCurrency model.CryptoCurrency) (*mongo.InsertOneResult, error) {
	//criar um contexto com deadline de 10 segundos
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	CryptoCurrencyCollection := mongodbClient.Database(DB_NAME).Collection(collection_name)

	insertResult, err := CryptoCurrencyCollection.InsertOne(mongoContext, CryptoCurrency)
	return insertResult, err
}

//Função para deletar
/*
	Função para deletar varios registros
	usa na entrada filter := bson.M{"key": "value"} criado em outra camada
*/
func DeleteManyCryptoCurrency(mongodbClient *mongo.Client, filter bson.M) (*mongo.DeleteResult, error) {
	//criar um contexto com deadline de 10 segundos
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	CryptoCurrencyCollection := mongodbClient.Database(DB_NAME).Collection(collection_name)

	deleteResult, err := CryptoCurrencyCollection.DeleteMany(mongoContext, filter)
	return deleteResult, err
}

//Função para buscar
/*
	Função para buscar um model.CryptoCurrency
	usa na entrada string do id
*/
func FindOneCryptoCurrencyByIdHex(mongodbClient *mongo.Client, idHex string) (model.CryptoCurrency, error) {
	//criar um contexto com deadline de 10 segundos
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	CryptoCurrencyCollection := mongodbClient.Database(DB_NAME).Collection(collection_name)

	// cria os parametros do filtro sem restrições
	primitiveObjectID, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": primitiveObjectID}

	err = CryptoCurrencyCollection.FindOne(mongoContext, filter).Decode(&objCryptoCurrency)
	return objCryptoCurrency, err
}

/*
	Função para buscar um model.CryptoCurrency
	usa na entrada filter := bson.M{"key": "value"} criado em outra camada
*/
func FindOneCryptoCurrency(mongodbClient *mongo.Client, filter bson.M) (model.CryptoCurrency, error) {
	//criar um contexto com deadline de 10 segundos
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	CryptoCurrencyCollection := mongodbClient.Database(DB_NAME).Collection(collection_name)

	err := CryptoCurrencyCollection.FindOne(mongoContext, filter).Decode(&objCryptoCurrency)
	return objCryptoCurrency, err
}

/*
	Função para buscar vários registros de model.CryptoCurrency
	usa na entrada filter := bson.M{"key": "value"} criado em outra camada
*/
func FindManyCryptoCurrency(mongodbClient *mongo.Client, filter bson.M) ([]model.CryptoCurrency, error) {
	//criar um contexto com deadline de 10 segundos
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	CryptoCurrencyCollection := mongodbClient.Database(DB_NAME).Collection(collection_name)

	cursor, err := CryptoCurrencyCollection.Find(mongoContext, filter)

	// chamada ao banco CryptoCurrencyCollection.Find(mongoContext, filter) tem erro
	if err != nil {
		defer cancel()

	} else {
		// se a chamada ao banco estiver ok
		sliceObjCryptoCurrency = nil
		for cursor.Next(mongoContext) {
			err = cursor.Decode(&objCryptoCurrency)
			objCryptoCurrency.IdHex = objCryptoCurrency.Id.Hex()
			sliceObjCryptoCurrency = append(sliceObjCryptoCurrency, objCryptoCurrency)
		}
	}
	return sliceObjCryptoCurrency, err
}

//Função para atualizar
/*
	Função para deletar um registro
	usa na entrada filter := bson.M{"last_name": "silva"}

	usa na entrada newData := bson.M{"age": "93"}
*/
func UpdateOneCryptoCurrency(mongodbClient *mongo.Client, filter bson.M, newData bson.M) (model.CryptoCurrency, error) {
	//criar um contexto com deadline de 10 segundos
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	CryptoCurrencyCollection := mongodbClient.Database(DB_NAME).Collection(collection_name)

	//se o documento não exisitr não faz nada
	//se alterar para true cria um novo documento caso não seja localizado
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	updateResult := CryptoCurrencyCollection.FindOneAndUpdate(mongoContext, filter, newData, &opt)
	err := updateResult.Decode(&objCryptoCurrency)
	return objCryptoCurrency, err
}
