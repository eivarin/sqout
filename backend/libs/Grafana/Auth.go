package Grafana

import (
	"encoding/base64"
	"fmt"
	"os"
)

type AuthDetails struct {
	user       string
	pass       string
	ServAccId  int
	AuthHeader *string
}

func NewAuthDetails(authHeaderPTR *string) AuthDetails {
	ad := AuthDetails{}
	ad.AuthHeader = authHeaderPTR
	ad.setBasicAuthHeader()
	return ad
}

func (ad *AuthDetails) setBasicAuthHeader() {
	ad.user = os.Getenv("GRAPHANA_USER")
	if ad.user == "" {
		ad.user = "admin"
	}
	ad.pass = os.Getenv("GRAPHANA_PASSWORD")
	if ad.pass == "" {
		ad.pass = "admin"
	}
	*ad.AuthHeader = "Basic " + base64.StdEncoding.EncodeToString([]byte(ad.user+":"+ad.pass))
}

func (gs *GrafanaState) setServiceAccountId() {
	createServiceAccountBody := struct {
		Name string `json:"name"`
		Role string `json:"role"`
	}{
		Name: "sqout",
		Role: "Admin",
	}
	resp := gs.http("/api/serviceaccounts/search?query="+createServiceAccountBody.Name, "GET", createServiceAccountBody)
	responseData := DecodeObjectJSON(resp)
	var saJSON map[string]interface{}
	if int(responseData["totalCount"].(float64)) == 0 {
		resp = gs.http("/api/serviceaccounts", "POST", createServiceAccountBody)
		if resp.StatusCode != 201 {
			fmt.Printf("Error: %s\n", resp.Status)
			panic("Failed to create service account")
		}
		saJSON = DecodeObjectJSON(resp)
	} else {
		saJSON = responseData["serviceAccounts"].([]interface{})[0].(map[string]interface{})
	}
	gs.AuthDet.ServAccId = int(saJSON["id"].(float64))
}

func (gs *GrafanaState) setServAccTokenAuth() {
	gs.deleteAllServAccTokens()
	getServiceAccountTokenBody := struct {
		Name string `json:"name"`
	}{
		Name: "sqout-sa",
	}
	resp := gs.http(fmt.Sprintf("/api/serviceaccounts/%d/tokens", gs.AuthDet.ServAccId), "POST", getServiceAccountTokenBody)
	if resp.StatusCode != 200 {
		fmt.Printf("Error: %s\n", resp.Status)
		panic("Failed to get service account token")
	}
	responseData := DecodeObjectJSON(resp)
	*gs.AuthDet.AuthHeader = "Bearer " + responseData["key"].(string)
}

func (gs *GrafanaState) deleteAllServAccTokens() {
	resp := gs.http(fmt.Sprintf("/api/serviceaccounts/%d/tokens", gs.AuthDet.ServAccId), "GET", nil)
	responseData := DecodeArrayJSON(resp)
	for _, token := range responseData {
		id := int(token["id"].(float64))
		resp = gs.http(fmt.Sprintf("/api/serviceaccounts/%d/tokens/%d", gs.AuthDet.ServAccId, id), "DELETE", nil)
		if resp.StatusCode != 200 {
			fmt.Printf("Error: %s\n", resp.Status)
			panic("Failed to delete service account token")
		}
	}
}
