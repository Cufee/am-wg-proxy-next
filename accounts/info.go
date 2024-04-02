package accounts

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/cufee/am-wg-proxy-next/internal/wargaming/client"
	"github.com/cufee/am-wg-proxy-next/types"
)

type infoResponse struct {
	types.WgResponse
	Data map[string]types.ExtendedAccount `json:"data"`
}

func GetAccountInfo(realm string, id string, fields ...string) (*types.ExtendedAccount, error) {
	accountsMap, err := GetBulkAccountsInfo(realm, []string{id}, fields...)
	if err != nil {
		return nil, err
	}

	info, ok := accountsMap[id]
	if !ok || info.ID == 0 {
		return nil, errors.New("account not found")
	}
	return &info, nil
}

func GetBulkAccountsInfo(realm string, ids []string, fields ...string) (map[string]types.ExtendedAccount, error) {
	var response infoResponse
	query := url.Values{}
	if len(fields) > 0 {
		query.Set("fields", strings.Join(fields, ","))
	}
	query.Set("extra", "statistics.rating")
	query.Set("account_id", strings.Join(ids, ","))

	_, err := client.WargamingRequest(realm, fmt.Sprintf("account/info/?%s", query.Encode()), "GET", nil, &response)
	if err != nil {
		return nil, err
	}
	if response.Error.Code != 0 {
		return nil, errors.New(response.Error.Message)
	}
	return response.Data, nil
}
