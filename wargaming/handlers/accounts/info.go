package accounts

import (
	"errors"
	"fmt"

	"aftermath.link/repo/am-wg-proxy/wargaming/client"
	"github.com/byvko-dev/am-types/wargaming/generic/api"
	"github.com/byvko-dev/am-types/wargaming/v1/accounts"
)

type InfoResponse struct {
	api.Response
	Data map[string]accounts.CompleteProfile `json:"data"`
}

func GetAccountInfo(bucket, realm string, playerId int) (accounts.CompleteProfile, error) {
	var response InfoResponse
	_, err := client.WargamingRequest(bucket, realm, fmt.Sprintf("account/info/?account_id=%v&extra=statistics.rating", playerId), "GET", nil, &response)
	if err != nil {
		return accounts.CompleteProfile{}, err
	}
	if response.Error.Code != 0 {
		return accounts.CompleteProfile{}, errors.New(response.Error.Message)
	}

	info, ok := response.Data[fmt.Sprint(playerId)]
	if !ok {
		return info, errors.New("account not found")
	}
	return info, nil
}
