package accounts

import (
	"errors"
	"fmt"

	"aftermath.link/repo/am-wg-proxy/wargaming/client"
	"aftermath.link/repo/am-wg-proxy/wargaming/handlers"
)

type VehiclesResponse struct {
	handlers.WargamingBaseResponse
	Data map[string][]VehicleStats `json:"data"`
}

func GetAccountVehicles(bucket, realm string, playerId int) ([]VehicleStats, error) {
	var response VehiclesResponse
	_, err := client.WargamingRequest(bucket, realm, fmt.Sprintf("tanks/stats/?account_id=%v", playerId), "GET", nil, &response)
	if err != nil {
		return nil, err
	}
	if response.Error.Code != 0 {
		return nil, errors.New(response.Error.Message)
	}

	info, ok := response.Data[fmt.Sprint(playerId)]
	if !ok {
		return info, errors.New("account not found")
	}
	return info, nil
}
