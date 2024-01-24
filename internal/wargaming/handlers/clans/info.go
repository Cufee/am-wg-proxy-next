package clans

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/cufee/am-wg-proxy-next/internal/wargaming/client"
	"github.com/cufee/am-wg-proxy-next/types"
)

type ClanInfoResponse struct {
	types.WgResponse
	Data map[string]types.ExtendedClan `json:"data"`
}

func GetClanInfo(realm string, clanId string, fields ...string) (*types.ExtendedClan, error) {
	data, err := GetBulkClanInfo(realm, []string{clanId}, fields...)
	if err != nil {
		return nil, err
	}

	info, ok := data[clanId]
	if !ok || info.ID == 0 {
		return nil, errors.New("clan not found")
	}
	return &info, nil
}

func GetBulkClanInfo(realm string, ids []string, fields ...string) (map[string]types.ExtendedClan, error) {
	var response ClanInfoResponse
	var query url.Values
	if len(fields) > 0 {
		query.Set("fields", strings.Join(fields, ","))
	}
	query.Set("clan_id", strings.Join(ids, ","))

	_, err := client.WargamingRequest(realm, fmt.Sprintf("clans/info/?%s", query.Encode()), "GET", nil, &response)
	if err != nil {
		return nil, err
	}
	if response.Error.Code != 0 {
		return nil, errors.New(response.Error.Message)
	}

	return response.Data, nil
}
