package accounts

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cufee/am-wg-proxy-next/internal/wargaming/client"
	"github.com/cufee/am-wg-proxy-next/types"
)

type InfoResponse struct {
	types.WgResponse
	Data map[string]types.ExtendedAccount `json:"data"`
}

func GetAccountInfo(realm string, id string) (*types.ExtendedAccount, error) {
	accountsMap, err := GetBulkAccountsInfo(realm, id)
	if err != nil {
		return nil, err
	}

	info, ok := accountsMap[id]
	if !ok || info.ID == 0 {
		return nil, errors.New("account not found")
	}
	return &info, nil
}

func GetBulkAccountsInfo(realm string, ids ...string) (map[string]types.ExtendedAccount, error) {
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
