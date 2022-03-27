package accounts

import (
	"errors"
	"fmt"

	"aftermath.link/repo/am-wg-proxy/wargaming/client"
	"github.com/byvko-dev/am-types/wargaming/generic/api"
	"github.com/byvko-dev/am-types/wargaming/v1/statistics"
)

type AchievementsResponse struct {
	api.Response
	Data map[string]statistics.AchievementsFrame `json:"data"`
}

func GetAccountAchievements(bucket, realm string, playerId int) (statistics.AchievementsFrame, error) {
	var response AchievementsResponse
	_, err := client.WargamingRequest(bucket, realm, fmt.Sprintf("account/achievements/?account_id=%v&fields=achievements", playerId), "GET", nil, &response)
	if err != nil {
		return statistics.AchievementsFrame{}, err
	}
	if response.Error.Code != 0 {
		return statistics.AchievementsFrame{}, errors.New(response.Error.Message)
	}

	info, ok := response.Data[fmt.Sprint(playerId)]
	if !ok {
		return info, errors.New("account not found")
	}
	return info, nil
}
