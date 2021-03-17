package shared

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoConn(strConn string, ctx context.Context) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(strConn))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return client, nil

}
