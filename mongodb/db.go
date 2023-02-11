package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDb struct {

	Client *mongo.Client

}

var instance *MongoDb = &MongoDb{}

func Instance() *MongoDb {
	if instance == nil {
		instance = &MongoDb{}
		instance.Conect()
	}
	instance.Conect()
	return instance
}

func (m * MongoDb) Conect() error {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGO_URI");
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
			panic(err)
	}
	fmt.Println("Connect")
	m.Client = client
	return err
}

func (m * MongoDb) Disconect() {
	err := m.Client.Disconnect(context.TODO())
	if(err != nil) {
		panic("error disconnect MongoDB")
	}
	fmt.Println("disconectado Mongo")
}