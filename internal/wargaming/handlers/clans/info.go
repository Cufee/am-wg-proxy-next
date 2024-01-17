package clans

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cufee/am-wg-proxy-next/internal/wargaming/client"
	"github.com/cufee/am-wg-proxy-next/types"
)

type ClanInfoResponse struct {
	types.WgResponse
	Data map[string]types.ExtendedClan `json:"data"`
}

func GetClanInfo(realm string, clanId string) (*types.ExtendedClan, error) {
	data, err := GetBulkClanInfo(realm, clanId)
	if err != nil {
		return nil, err
	}

	info, ok := data[clanId]
	if !ok || info.ID == 0 {
		return nil, errors.New("clan not found")
	}
	return &info, nil
}

func GetBulkClanInfo(realm string, ids ...string) (map[string]types.ExtendedClan, error) {
	var response ClanInfoResponse
	_, err := client.WargamingRequest(realm, fmt.Sprintf("clans/info/?clan_id=%s", strings.Join(ids, ",")), "GET", nil, &response)
	if err != nil {
		return nil, err
	}
	if response.Error.Code != 0 {
		return nil, errors.New(response.Error.Message)
	}

	return response.Data, nil
}
