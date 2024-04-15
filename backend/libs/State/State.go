package State

import (
	"context"
	"sqout/libs/DbApi"
	"go.mongodb.org/mongo-driver/mongo"
)

type State struct {
	dbClient *mongo.Client
	ModulesCol DbApi.ColFacade
	ProbesCol  DbApi.ColFacade
}


func InitState(c context.Context) *State {
	s := new(State)
	s.dbClient = DbApi.InitDB(c)
	s.ModulesCol = DbApi.NewColFacade(s.dbClient, "modules")
	s.ProbesCol = DbApi.NewColFacade(s.dbClient, "probes")
	return s
}