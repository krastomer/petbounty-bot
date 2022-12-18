package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repositories struct {
	client *mongo.Client
}

func InitializeRepositories() Repositories {
	uri := "mongodb+srv://admin:AwesomePassword@cluster.d1xadz9.mongodb.net/?retryWrites=true&w=majority"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return Repositories{
		client: client,
	}
}
