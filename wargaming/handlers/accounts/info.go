package accounts

import (
	"errors"
	"fmt"

	"aftermath.link/repo/am-wg-proxy/wargaming/client"
	"aftermath.link/repo/am-wg-proxy/wargaming/handlers"
)

type InfoResponse struct {
	handlers.WargamingBaseResponse
	Data map[string]DetailedInfo `json:"data"`
}

type DetailedInfo struct {
	BaseDetails
	Statistics struct {
		All    StatsFrame `json:"all"`
		Rating StatsFrame `json:"rating"`
	} `json:"statistics"`
	CreatedAt      int64 `json:"created_at"`
	UpdatedAt      int64 `json:"updated_at"`
	LastBattleTime int64 `json:"last_battle_time"`
}

func GetAccountInfo(bucket, realm string, playerId int) (DetailedInfo, error) {
	var response InfoResponse
	_, err := client.WargamingRequest(bucket, realm, fmt.Sprintf("account/info/?account_id=%v&extra=statistics.rating", playerId), "GET", nil, &response)
	if err != nil {
		return DetailedInfo{}, err
	}
	if response.Error.Code != 0 {
		return DetailedInfo{}, errors.New(response.Error.Message)
	}

	info, ok := response.Data[fmt.Sprint(playerId)]
	if !ok {
		return info, errors.New("account not found")
	}
	return info, nil
}
