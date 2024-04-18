package DbApi

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ColFacade struct {
	Col *mongo.Collection
}

func NewColFacade(cl *mongo.Client, colName string) ColFacade {
	return ColFacade{Col: cl.Database("sqout").Collection(colName)}
}

func InitDB(c context.Context) *mongo.Client {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = "mongodb://root:example@localhost:27017"
	}
	fmt.Printf("Connecting to mongo using: \"%s\"\n", uri)
	client, err := mongo.Connect(c, options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Printf("Error connecting to mongo: %v\n", err)
	}
	return client
}
