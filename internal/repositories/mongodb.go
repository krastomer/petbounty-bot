package repositories

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repositories struct {
	client *mongo.Client
}

func InitializeRepositories() Repositories {
	uri := os.Getenv("DATABASE_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return Repositories{
		client: client,
	}
}

func (r *Repositories) GetDatabase() *mongo.Database {
	return r.client.Database("petbounty-bot")
}
