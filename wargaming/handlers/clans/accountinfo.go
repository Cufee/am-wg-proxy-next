package clans

import (
	"errors"
	"fmt"

	"aftermath.link/repo/am-wg-proxy/wargaming/client"
	"github.com/byvko-dev/am-types/wargaming/generic/api"
	"github.com/byvko-dev/am-types/wargaming/v1/clans"
)

type AccountClanInfoResponse struct {
	api.Response
	Data map[string]clans.MemberProfile `json:"data"`
}

func GetAccountClanInfo(bucket, realm string, playerId int) (clans.MemberProfile, error) {
	var response AccountClanInfoResponse
	_, err := client.WargamingRequest(bucket, realm, fmt.Sprintf("clans/accountinfo/?account_id=%v&extra=clan", playerId), "GET", nil, &response)
	if err != nil {
		return clans.MemberProfile{}, err
	}
	if response.Error.Code != 0 {
		return clans.MemberProfile{}, errors.New(response.Error.Message)
	}

	info, ok := response.Data[fmt.Sprint(playerId)]
	if !ok {
		return info, errors.New("account not found")
	}
	return info, nil
}
