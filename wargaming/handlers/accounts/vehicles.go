package accounts

import (
	"errors"
	"fmt"

	"aftermath.link/repo/am-wg-proxy/wargaming/client"
	"github.com/byvko-dev/am-types/wargaming/generic/api"
	"github.com/byvko-dev/am-types/wargaming/v1/statistics"
)

type VehiclesResponse struct {
	api.Response
	Data map[string][]statistics.VehicleStatsFrame `json:"data"`
}

func GetAccountVehicles(bucket, realm string, playerId int) ([]statistics.VehicleStatsFrame, error) {
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
