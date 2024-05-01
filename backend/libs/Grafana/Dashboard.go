package Grafana

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
)

type Dashboard struct {
	Annotations          interface{} `json:"annotations"`
	Editable             bool        `json:"editable"`
	FiscalYearStartMonth int         `json:"fiscalYearStartMonth"`
	GraphTooltip         int         `json:"graphTooltip"`
	Id                   int         `json:"id"`
	Links                interface{} `json:"links"`
	Panels               interface{} `json:"panels"`
	Refresh              string      `json:"refresh"`
	SchemaVersion        int         `json:"schemaVersion"`
	Tags                 interface{} `json:"tags"`
	Templating           interface{} `json:"templating"`
	Time                 interface{} `json:"time"`
	Timepicker           interface{} `json:"timepicker"`
	Timezone             string      `json:"timezone"`
	Title                string      `json:"title"`
	Uid                  string      `json:"uid"`
	Version              int         `json:"version"`
	WeekStart            string      `json:"weekStart"`
}

func (gs *GrafanaState) LoadDashboardForProbe(path, probeName string) Dashboard {
	gs.DeleteDashboardOnGrafana(probeName)
	dashText := GetStringJsonFromFile(path)
	dashText = gs.ReplaceUrlInJsonString(dashText, probeName)
	var dash Dashboard
	if err := json.Unmarshal([]byte(dashText), &dash); err != nil {
		panic(err)
	}
	dash.Title = probeName
	dash.Uid = probeName
	dash.Id = 0
	gs.createDashboardOnGrafana(dash)
	return dash
}

func GetStringJsonFromFile(filePath string) string {
	file, _ := os.Open(filePath)
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func (gs *GrafanaState) ReplaceUrlInJsonString(json string, ProbeName string) string {
	r, _ := regexp.Compile(`"url":\s".*"`)
	link := fmt.Sprintf(`"url": "%s/probes/%s?includeResults=true"`, gs.BackendUrl, ProbeName)
	r1, _ := regexp.Compile(`"datasource": \{\n\s+"type": "yesoreyeram-infinity-datasource",?(\n.*){1,2}\}`)
	datasource := `"datasource": {
		"type": "yesoreyeram-infinity-datasource"
	}`
	return r1.ReplaceAllString(r.ReplaceAllString(json, link), datasource)
}

func (gs *GrafanaState) createDashboardOnGrafana(d Dashboard) {
	body := struct {
		Dashboard Dashboard `json:"dashboard"`
		Overwrite bool      `json:"overwrite"`
		Message   string    `json:"message"`
	}{
		Dashboard: d,
		Overwrite: true,
		Message:   "Created by sqout",
	}
	resp := gs.http("/api/dashboards/db", "POST", body)
	if resp.StatusCode != 200 {
		fmt.Printf("Error: %s\n", resp.Status)
		fmt.Println("Failed to create dashboard")
	}
}

func (gs *GrafanaState) DeleteDashboardOnGrafana(probeName string) {
	resp := gs.http(fmt.Sprintf("/api/dashboards/uid/%s", probeName), "DELETE", nil)
	if resp.StatusCode != 200 {
		fmt.Printf("Error: %s\n", resp.Status)
		fmt.Println("Failed to delete dashboard")
	}
}
