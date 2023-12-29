package clans

import (
	"errors"
	"fmt"

	"github.com/cufee/am-wg-proxy-next/internal/wargaming/client"
	"github.com/cufee/am-wg-proxy-next/types"
)

type ClanSearchResponse struct {
	types.WgResponse
	Data []types.Clan `json:"data"`
}

func SearchClans(realm, search string) ([]types.Clan, error) {
	var response ClanSearchResponse
	_, err := client.WargamingRequest(realm, fmt.Sprintf("clans/list/?search=%v&limit=3", search), "GET", nil, &response)
	if err != nil {
		return nil, err
	}
	if response.Error.Code != 0 {
		return nil, errors.New(response.Error.Message)
	}
	return response.Data, nil
}
