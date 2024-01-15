package accounts

import (
	"errors"
	"fmt"

	"github.com/cufee/am-wg-proxy-next/internal/wargaming/client"
	"github.com/cufee/am-wg-proxy-next/types"
	"github.com/rs/zerolog/log"
)

type SearchResponse struct {
	types.WgResponse
	Data []types.Account `json:"data"`
}

func SearchAccounts(realm, query string) ([]types.Account, error) {
	var response SearchResponse
	_, err := client.WargamingRequest(realm, fmt.Sprintf("account/list/?search=%v&limit=3", query), "GET", nil, &response)
	if err != nil {
		return nil, err
	}
	if response.Error.Code != 0 {
		log.Error().Str("realm", realm).Str("query", query).Msg("Error while searching accounts")
		return nil, errors.New(response.Error.Message)
	}

	return response.Data, nil
}
