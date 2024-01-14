package clans

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cufee/am-wg-proxy-next/internal/wargaming/client"
	"github.com/cufee/am-wg-proxy-next/types"
)

type AccountClanInfoResponse struct {
	types.WgResponse
	Data map[string]types.ClanMember `json:"data"`
}

func GetAccountClanInfo(realm string, playerId string) (*types.ClanMember, error) {
	data, err := GetBulkAccountClanInfo(realm, playerId)
	if err != nil {
		return nil, err
	}

	info, ok := data[fmt.Sprint(playerId)]
	if !ok || info.ClanID == 0 {
		return nil, errors.New("clan not found")
	}
	return nil, nil
}

func GetBulkAccountClanInfo(realm string, ids ...string) (map[string]types.ClanMember, error) {
	var response AccountClanInfoResponse
	_, err := client.WargamingRequest(realm, fmt.Sprintf("clans/accountinfo/?account_id=%v&extra=clan", strings.Join(ids, ",")), "GET", nil, &response)
	if err != nil {
		return nil, err
	}

	if response.Error.Code != 0 {
		return nil, errors.New(response.Error.Message)
	}
	return response.Data, nil
}
