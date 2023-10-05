package mongo

import (
	"awesomeProject/infra/config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetConnection(ctx context.Context) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Env.MongoUrl))
	if err != nil {
		return client, err
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return client, err
	}

	return client, nil
}

func GetCollection(client *mongo.Client, name string) *mongo.Collection {
	return client.Database(config.Env.MongoDatabase).Collection(name)
}
