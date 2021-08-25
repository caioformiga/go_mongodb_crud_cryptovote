package dao

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Singleton design pattern
var mongoClientInstance *mongo.Client

var mongoClientInstanceError error

var mongoOnce sync.Once

var collection_name string = "cryptovotes"
var db_name string = os.Getenv("MONGODB_CRYPTOVOTE")
var db_uri = os.Getenv("MONGODB_URI")

/*
	Não existe uma declaração explicita de que um tipo (struct) implementa uma interface.
	Isso acontece quando um tipo implementa todo os metodos da interface, como este tipo
	dao.Mongodb implementa os 7 metodos de dao.InterfaceMongodbServie
*/
type MongodbService struct {
}

func (m MongodbService) SetDatabaseName(db string) {
	db_name = db
}

func (m MongodbService) GetDatabaseName() string {
	return db_name
}

func (m MongodbService) SetCollectionName(collection string) {
	collection_name = collection
}

func (m MongodbService) GetCollectionName() string {
	return collection_name
}

func (m MongodbService) SetDatabaseUri(uri string) {
	db_uri = uri
}

func (m MongodbService) GetDatabaseUri() string {
	return db_uri
}

// returns the singleton connection instantiation
func GetMongoClientInstance(dbService interfaces.InterfaceDbService) (*mongo.Client, error) {
	mongoOnce.Do(func() {
		// configure connection options
		mongoOptions := options.Client().ApplyURI(dbService.GetDatabaseUri())

		// create a connection but dosn't starts it
		client, err := connect(mongoOptions, dbService)
		if err != nil {
			mongoClientInstanceError = err
		}
		mongoClientInstance = client
	})
	return mongoClientInstance, mongoClientInstanceError
}

func connect(mongoOptions *options.ClientOptions, dbService interfaces.InterfaceDbService) (*mongo.Client, error) {

	// create a context with 10-second deadline
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// prepares connection with MongoDB
	mongoClient, err := mongo.Connect(mongoContext, mongoOptions)
	if err != nil {
		log.Println("[cryptovote.mongodb] Connection erro, see: mongo.Connect(mongoContext, client_mongo)")
		log.Printf("[cryptovote.mongodb] %v", err)
		return mongoClient, err
	}

	// test connection
	err = mongoClient.Ping(mongoContext, readpref.Primary())
	if err != nil {
		log.Printf("[cryptovote.mongodb] Connection refused at %v", dbService.GetDatabaseUri())
		log.Printf("[cryptovote.mongodb] %v", err)
		diconnect(mongoClient)
		return mongoClient, err
	}
	log.Println("[cryptovote.mongodb] Ping done successfully!")

	if err != nil {
		log.Printf("[cryptovote.mongodb] %v", err)
		return mongoClient, err
	}
	log.Printf("[cryptovote.mongodb] Connection doen successfully at %v database (%v)...", dbService.GetDatabaseUri(), dbService.GetDatabaseName())

	return mongoClient, err
}

func diconnect(mongoClient *mongo.Client) {
	// create a context with 10-second deadline
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// close connection
	err := mongoClient.Disconnect(mongoContext)
	if err != nil {
		log.Fatalf("[cryptovote.mongodb] %v", err)
	}
	log.Println("[cryptovote.mongodb] Closing connection from mongoDB")
}
