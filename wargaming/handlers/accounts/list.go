package accounts

import (
	"errors"
	"fmt"

	"aftermath.link/repo/am-wg-proxy/logs"
	"aftermath.link/repo/am-wg-proxy/wargaming/client"
	"github.com/byvko-dev/am-types/wargaming/generic/api"
	"github.com/byvko-dev/am-types/wargaming/v1/accounts"
)

type SearchResponse struct {
	api.Response
	Data []accounts.BaseProfile `json:"data"`
}

func SearchAccounts(bucket, realm, query string) ([]accounts.BaseProfile, error) {
	var response SearchResponse
	_, err := client.WargamingRequest(bucket, realm, fmt.Sprintf("account/list/?search=%v&limit=3", query), "GET", nil, &response)
	if err != nil {
		return nil, err
	}
	if response.Error.Code != 0 {
		logs.Error("Error while searching accounts: %+v", response.Error)
		return nil, errors.New(response.Error.Message)
	}

	return response.Data, nil
}
