package State

import (
	"context"
	"sqout/libs/DbApi"
	"sqout/libs/Grafana"
	"sqout/libs/ModuleConfig"
	"sqout/libs/Probe"
	"sqout/libs/TimersMap"

	"go.mongodb.org/mongo-driver/mongo"
)

type State struct {
	DbClient   *mongo.Client
	ModulesCol DbApi.ColFacade
	ProbesCol  DbApi.ColFacade
	Timers     TimersMap.TimersMap
	GS         *Grafana.GrafanaState
}

func InitState(c context.Context) *State {
	s := new(State)
	s.DbClient = DbApi.InitDB(c)
	s.ModulesCol = DbApi.NewColFacade(s.DbClient, "modules")
	s.ProbesCol = DbApi.NewColFacade(s.DbClient, "probes")
	s.Timers = TimersMap.NewTimersMap()
	s.GS = Grafana.NewGrafanaState()
	ModuleConfig.SanitizeModulesByDB(c, &s.ModulesCol)
	Probe.RestartAllProbes(c, &s.ModulesCol, &s.ProbesCol, &s.Timers, s.GS)
	return s
}
