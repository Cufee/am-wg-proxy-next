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

func GetAccountClanInfo(realm string, playerId string) (clans.MemberProfile, error) {
	data, err := GetBulkAccountClanInfo(realm, playerId)
	if err != nil {
		return clans.MemberProfile{}, err
	}

	info, ok := data[fmt.Sprint(playerId)]
	if !ok {
		return info, errors.New("account not found")
	}
	return info, nil
}

func GetBulkAccountClanInfo(realm string, ids ...string) (map[string]clans.MemberProfile, error) {
	var response AccountClanInfoResponse
	_, err := client.WargamingRequest(realm, fmt.Sprintf("clans/accountinfo/?account_id=%v&extra=clan", ids), "GET", nil, &response)
	if err != nil {
		return nil, err
	}
	if response.Error.Code != 0 {
		return nil, errors.New(response.Error.Message)
	}
	return response.Data, nil
}
