package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Client, error) {
	// Define the mongodb client URL
	uri := "mongodb://admin:password@mongo:27017/manager?authSource=admin&authMechanism=SCRAM-SHA-1"

	// Establish the connection
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("Problem ping mongo", err.Error())
	}

	// Create go routine to defer the closure
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			fmt.Println("Disconnect mongo")
		}
	}()
	log.Println("returning mongo client")
	return client, nil
}

// https://github.com/wpcodevo/golang-mongodb-api/blob/master/cmd/server/main.go
