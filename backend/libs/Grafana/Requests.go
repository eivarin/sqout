package Grafana

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type GrafanaHTTP struct {
	Url        string
	AuthHeader *string
}

func NewGrafanaHTTP(authHeader *string) GrafanaHTTP {
	var gAPI GrafanaHTTP
	gAPI.Url = os.Getenv("GRAFANA_DOCKER_URL")
	if gAPI.Url == "" {
		gAPI.Url = "http://grafana:3000"
	}
	gAPI.AuthHeader = authHeader
	return gAPI
}

func (gAPI *GrafanaHTTP) makeRequest(endpoint string, method string, body interface{}) *http.Response {
	fmt.Printf("Making %s request to %s\n", method, endpoint)
	client := &http.Client{}
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	uri := gAPI.Url + endpoint
	req, err := http.NewRequest(method, uri, bytes.NewBuffer(bodyJSON))
	req.Header.Set("Authorization", *gAPI.AuthHeader)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%s - %s %s with body:\n%s\n\n", err, method, uri, bodyJSON)
	}
	return resp
}

func DecodeObjectJSON(resp *http.Response) map[string]interface{} {
	var responseData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		panic(err)
	}
	return responseData
}

func DecodeArrayJSON(resp *http.Response) []map[string]interface{} {
	var responseData []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		panic(err)
	}
	return responseData
}
