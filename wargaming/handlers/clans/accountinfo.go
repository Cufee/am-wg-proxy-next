package clans

import (
	"errors"
	"fmt"

	"aftermath.link/repo/am-wg-proxy/wargaming/client"
	"aftermath.link/repo/am-wg-proxy/wargaming/handlers"
)

type AccountClanInfoResponse struct {
	handlers.WargamingBaseResponse
	Data map[string]AccountClanInfo `json:"data"`
}

type AccountClanInfo struct {
	Clan        AccountClan `json:"clan"`
	AccountID   int64       `json:"account_id"`
	JoinedAt    int64       `json:"joined_at"`
	ClanID      int64       `json:"clan_id"`
	Role        string      `json:"role"`
	AccountName string      `json:"account_name"`
}

type AccountClan struct {
	BasicClanDetails
	EmblemSetID int64 `json:"emblem_set_id"`
}

func GetAccountClanInfo(bucket, realm string, playerId int) (AccountClanInfo, error) {
	var response AccountClanInfoResponse
	_, err := client.WargamingRequest(bucket, realm, fmt.Sprintf("clans/accountinfo/?account_id=%v&extra=clan", playerId), "GET", nil, &response)
	if err != nil {
		return AccountClanInfo{}, err
	}
	if response.Error.Code != 0 {
		return AccountClanInfo{}, errors.New(response.Error.Message)
	}

	info, ok := response.Data[fmt.Sprint(playerId)]
	if !ok {
		return info, errors.New("account not found")
	}
	return info, nil
}
