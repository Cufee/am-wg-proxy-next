package accounts

import (
	"errors"
	"fmt"
	"strings"

	"github.com/byvko-dev/am-types/wargaming/generic/api"
	"github.com/byvko-dev/am-types/wargaming/v1/statistics"
	"github.com/cufee/am-wg-proxy-next/internal/wargaming/client"
)

type VehiclesResponse struct {
	api.Response
	Data map[string][]statistics.VehicleStatsFrame `json:"data"`
}

type VehicleAchievementsResponse struct {
	api.Response
	Data map[string]statistics.AchievementsFrame `json:"data"`
}

func GetAccountVehicles(realm string, id string) ([]statistics.VehicleStatsFrame, error) {
	vehiclesMap, err := GetBulkAccountsVehicles(realm, id)
	if err != nil {
		return nil, err
	}

	info, ok := vehiclesMap[id]
	if !ok {
		return info, errors.New("account not found")
	}
	return info, nil
}

func GetBulkAccountsVehicles(realm string, ids ...string) (map[string][]statistics.VehicleStatsFrame, error) {
	var response VehiclesResponse
	_, err := client.WargamingRequest(realm, fmt.Sprintf("tanks/stats/?account_id=%s", strings.Join(ids, ",")), "GET", nil, &response)
	if err != nil {
		return nil, err
	}
	if response.Error.Code != 0 {
		return nil, errors.New(response.Error.Message)
	}
	return response.Data, nil
}
