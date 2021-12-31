package workers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var wargamingServerStatusURL = "https://api.worldoftanks.com/wgn/servers/info/?application_id=%v"

// This method is used to validate the API key.
func getWargamingServerStatus(appID string) (bool, error) {
	var serverStatus map[string]interface{}
	var err error

	resp, err := http.Get(fmt.Sprintf(wargamingServerStatusURL, appID))
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&serverStatus)
	if err != nil {
		return false, err
	}

	if serverStatus["status"] == "ok" {
		return true, nil
	}

	return false, nil
}
