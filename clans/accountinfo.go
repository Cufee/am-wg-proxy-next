package clans

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/cufee/am-wg-proxy-next/internal/wargaming/client"
	"github.com/cufee/am-wg-proxy-next/types"
)

type accountClanInfoResponse struct {
	types.WgResponse
	Data map[string]types.ClanMember `json:"data"`
}

func GetAccountClanInfo(realm string, playerId string, fields ...string) (*types.ClanMember, error) {
	data, err := GetBulkAccountClanInfo(realm, []string{playerId}, fields...)
	if err != nil {
		return nil, err
	}

	info, ok := data[fmt.Sprint(playerId)]
	if !ok || info.ClanID == 0 {
		return nil, errors.New("clan not found")
	}
	return nil, nil
}

func GetBulkAccountClanInfo(realm string, ids []string, fields ...string) (map[string]types.ClanMember, error) {
	var response accountClanInfoResponse
	query := url.Values{}
	if len(fields) > 0 {
		query.Set("fields", strings.Join(fields, ","))
	}
	query.Set("extra", "clan")
	query.Set("account_id", strings.Join(ids, ","))

	_, err := client.WargamingRequest(realm, fmt.Sprintf("clans/accountinfo/?%s", query.Encode()), "GET", nil, &response)
	if err != nil {
		return nil, err
	}

	if response.Error.Code != 0 {
		return nil, errors.New(response.Error.Message)
	}
	return response.Data, nil
}
