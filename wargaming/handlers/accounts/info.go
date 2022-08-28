package accounts

import (
	"errors"
	"fmt"
	"strings"

	"aftermath.link/repo/am-wg-proxy/wargaming/client"
	"github.com/byvko-dev/am-types/wargaming/generic/api"
	"github.com/byvko-dev/am-types/wargaming/v1/accounts"
)

type InfoResponse struct {
	api.Response
	Data map[string]accounts.CompleteProfile `json:"data"`
}

func GetAccountInfo(realm string, id string) (accounts.CompleteProfile, error) {
	accountsMap, err := GetBulkAccountsInfo(realm, id)
	if err != nil {
		return accounts.CompleteProfile{}, err
	}

	info, ok := accountsMap[id]
	if !ok {
		return info, errors.New("account not found")
	}
	return info, nil
}

func GetBulkAccountsInfo(realm string, ids ...string) (map[string]accounts.CompleteProfile, error) {
	var response InfoResponse
	_, err := client.WargamingRequest(realm, fmt.Sprintf("account/info/?account_id=%s&extra=statistics.rating", strings.Join(ids, ",")), "GET", nil, &response)
	if err != nil {
		return nil, err
	}
	if response.Error.Code != 0 {
		return nil, errors.New(response.Error.Message)
	}
	return response.Data, nil
}
