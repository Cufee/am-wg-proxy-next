package accounts

import (
	"errors"
	"fmt"

	"github.com/cufee/am-wg-proxy-next/internal/wargaming/client"
	"github.com/cufee/am-wg-proxy-next/types"
)

type VehiclesResponse struct {
	types.WgResponse
	Data map[string][]types.VehicleStatsFrame `json:"data"`
}

type VehicleAchievementsResponse struct {
	types.WgResponse
	Data map[string]types.AchievementsFrame `json:"data"`
}

func GetAccountVehicles(realm string, id string) ([]types.VehicleStatsFrame, error) {
	var response VehiclesResponse
	_, err := client.WargamingRequest(realm, fmt.Sprintf("tanks/stats/?account_id=%s", id), "GET", nil, &response)
	if err != nil {
		return nil, err
	}
	if response.Error.Code != 0 {
		return nil, errors.New(response.Error.Message)
	}

	info, ok := response.Data[id]
	if !ok {
		return info, errors.New("account not found")
	}
	return info, nil
}
