package clans

import (
	"errors"
	"fmt"

	"aftermath.link/repo/am-wg-proxy/wargaming/client"
	"github.com/byvko-dev/am-types/wargaming/generic/api"
	"github.com/byvko-dev/am-types/wargaming/v1/clans"
)

type ClanInfoResponse struct {
	api.Response
	Data map[string]clans.CompleteProfile `json:"data"`
}

func GetClanInfo(bucket, realm string, clanId int) (clans.CompleteProfile, error) {
	var response ClanInfoResponse
	_, err := client.WargamingRequest(bucket, realm, fmt.Sprintf("clans/info/?clan_id=%v", clanId), "GET", nil, &response)
	if err != nil {
		return clans.CompleteProfile{}, err
	}
	if response.Error.Code != 0 {
		return clans.CompleteProfile{}, errors.New(response.Error.Message)
	}

	info, ok := response.Data[fmt.Sprint(clanId)]
	if !ok {
		return info, errors.New("account not found")
	}
	return info, nil
}
