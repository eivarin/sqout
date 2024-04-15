package DbApi

import (
	"context"
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
	client, _ := mongo.Connect(c, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	return client
}