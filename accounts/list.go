package accounts

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/cufee/am-wg-proxy-next/internal/wargaming/client"
	"github.com/cufee/am-wg-proxy-next/types"
	"github.com/rs/zerolog/log"
)

type searchResponse struct {
	types.WgResponse
	Data []types.Account `json:"data"`
}

func SearchAccounts(realm, search string, fields ...string) ([]types.Account, error) {
	var response searchResponse
	query := url.Values{}
	if len(fields) > 0 {
		query.Set("fields", strings.Join(fields, ","))
	}
	query.Set("search", search)
	query.Set("limit", "3")

	_, err := client.WargamingRequest(realm, fmt.Sprintf("account/list/?%s", query.Encode()), "GET", nil, &response)
	if err != nil {
		return nil, err
	}
	if response.Error.Code != 0 {
		log.Error().Str("realm", realm).Str("query", search).Msg("Error while searching accounts")
		return nil, errors.New(response.Error.Message)
	}

	return response.Data, nil
}
