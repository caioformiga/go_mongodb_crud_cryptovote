package dao

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	//outra opção para conectar seria usado a DB_URI := "mongodb://localhost:27017"
	DB_URI = "mongodb://127.0.0.1:27017"
)

//padrão de projeto singleton
var mongoClientInstance *mongo.Client

var mongoClientInstanceError error

var mongoOnce sync.Once

var collection_name string = "cryptovotes"
var db_name string = "prod"

func SetDtabaseName(name string) {
	db_name = name
}

func GetDtabaseName() string {
	return db_name
}

func SetCollectionName(name string) {
	collection_name = name
}

func GetCollectionName() string {
	return collection_name
}

//retorna a instancia de conexão singleton
func GetMongoClientInstance() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		//faz uma conexão
		client, err := connect()
		if err != nil {
			mongoClientInstanceError = err
		}
		mongoClientInstance = client
	})
	return mongoClientInstance, mongoClientInstanceError
}

func connect() (*mongo.Client, error) {
	//configurando as opções de conexão
	mongoOptions := options.Client().ApplyURI(DB_URI)

	//criar um contexto com deadline de 10 segundos
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//conecta no MongoDB
	mongoClient, err := mongo.Connect(mongoContext, mongoOptions)
	if err != nil {
		log.Println("[cryptovote.mongodb] Erro na conexão, ver: mongo.Connect(mongoContext, client_mongo)")
	}
	log.Println("[cryptovote.mongodb] Conexão com mongoDB foi feita com sucesso...")

	//verifica a conexão
	err = mongoClient.Ping(mongoContext, readpref.Primary())
	if err != nil {
		log.Println("[cryptovote.mongodb] Perdeu a conexão, ver: mongoClient.Ping(mongoContext, readpref.Primary())")
		diconnect(mongoClient)
	}
	log.Println("[cryptovote.mongodb] Ping feito com sucesso!")
	return mongoClient, err
}

func diconnect(mongoClient *mongo.Client) {
	//criar um contexto com deadline de 10 segundos
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//fechar a conexão
	err := mongoClient.Disconnect(mongoContext)
	if err != nil {
		log.Fatalf("[cryptovote.mongodb] +%v", err)
	}
	log.Println("[cryptovote.mongodb] Fechando a conexão com mongoDB")
}
