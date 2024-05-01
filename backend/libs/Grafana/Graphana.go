package Grafana

import (
	"net/http"
	"os"
	"time"
)

type GrafanaState struct {
	AuthDet    AuthDetails
	GrafanaApi GrafanaHTTP
	BackendUrl string
}

func NewGrafanaState() *GrafanaState {
	gs := GrafanaState{}
	var authHeader string
	gs.BackendUrl = os.Getenv("BACKEND_DOCKER_URL")
	if gs.BackendUrl == "" {
		gs.BackendUrl = "http://backend:8080"
	}
	gs.AuthDet = NewAuthDetails(&authHeader)
	gs.GrafanaApi = NewGrafanaHTTP(&authHeader)
	gs.auth()
	return &gs
}

func (gs *GrafanaState) auth() {
	gs.WaitForGrafanaUp()
	gs.setServiceAccountId()
	gs.setServAccTokenAuth()
}

func (gs *GrafanaState) WaitForGrafanaUp() {
	for {
		resp := gs.http("/api/search", "GET", nil)
		if resp.StatusCode == 200 {
			break
		}
		time.Sleep(2 * time.Second)	
	}
}

func (gs *GrafanaState) http(endpoint string, method string, body interface{}) *http.Response {
	return gs.GrafanaApi.makeRequest(endpoint, method, body)
}
