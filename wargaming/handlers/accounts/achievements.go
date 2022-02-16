package accounts

import (
	"errors"
	"fmt"

	"aftermath.link/repo/am-wg-proxy/wargaming/client"
	"aftermath.link/repo/am-wg-proxy/wargaming/handlers"
)

type AchievementsResponse struct {
	handlers.WargamingBaseResponse
	Data map[string]AchievementsFrame `json:"data"`
}

func GetAccountAchievements(bucket, realm string, playerId int) (AchievementsFrame, error) {
	var response AchievementsResponse
	_, err := client.WargamingRequest(bucket, realm, fmt.Sprintf("account/achievements/?account_id=%v&fields=achievements", playerId), "GET", nil, &response)
	if err != nil {
		return AchievementsFrame{}, err
	}
	if response.Error.Code != 0 {
		return AchievementsFrame{}, errors.New(response.Error.Message)
	}

	info, ok := response.Data[fmt.Sprint(playerId)]
	if !ok {
		return info, errors.New("account not found")
	}
	return info, nil
}
