package clans

import (
	"errors"
	"fmt"

	"github.com/byvko-dev/am-types/wargaming/generic/api"
	"github.com/byvko-dev/am-types/wargaming/v1/clans"
	"github.com/cufee/am-wg-proxy-next/internal/wargaming/client"
)

type ClanInfoResponse struct {
	api.Response
	Data map[string]clans.CompleteProfile `json:"data"`
}

func GetClanInfo(realm string, clanId string) (clans.CompleteProfile, error) {
	data, err := GetBulkClanInfo(realm, clanId)
	if err != nil {
		return clans.CompleteProfile{}, err
	}

	info, ok := data[clanId]
	if !ok {
		return info, errors.New("account not found")
	}
	return info, nil
}

func GetBulkClanInfo(realm string, ids ...string) (map[string]clans.CompleteProfile, error) {
	var response ClanInfoResponse
	_, err := client.WargamingRequest(realm, fmt.Sprintf("clans/info/?clan_id=%v", ids), "GET", nil, &response)
	if err != nil {
		return nil, err
	}
	if response.Error.Code != 0 {
		return nil, errors.New(response.Error.Message)
	}

	return response.Data, nil
}
