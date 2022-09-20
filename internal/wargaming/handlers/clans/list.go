package clans

import (
	"errors"
	"fmt"

	"github.com/byvko-dev/am-types/wargaming/generic/api"
	"github.com/byvko-dev/am-types/wargaming/v1/clans"
	"github.com/cufee/am-wg-proxy-next/internal/wargaming/client"
)

type ClanSearchResponse struct {
	api.Response
	Data []clans.BasicProfile `json:"data"`
}

func SearchClans(realm, search string) ([]clans.BasicProfile, error) {
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
