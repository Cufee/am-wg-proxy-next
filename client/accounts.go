package client

import (
	"errors"
	"strings"

	"github.com/cufee/am-wg-proxy-next/types"
)

func (c *Client) SearchAccounts(realm, query string, fields ...string) (types.Account, error) {
	opts := newDefaultRequestOptions()
	opts.Query.Add("query", query)
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}

	var target []types.Account
	err := c.sendRequest(realm, accountsSearchEndpoint, &target, opts)
	if err != nil {
		return types.Account{}, err
	}
	if len(target) == 0 {
		return types.Account{}, errors.New("no results found")
	}
	return target[0], nil
}

func (c *Client) GetAccountByID(id int, fields ...string) (types.ExtendedAccount, error) {
	opts := newDefaultRequestOptions()
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}

	var target types.ExtendedAccount
	return target, c.sendRequest(RealmFromPlayerID(id), accountsGetEndpointFMT.Fmt(id), &target, opts)
}

func (c *Client) GetAccountClan(id int, fields ...string) (types.ClanMember, error) {
	opts := newDefaultRequestOptions()
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}

	var target types.ClanMember
	return target, c.sendRequest(RealmFromPlayerID(id), accountClanGetEndpointFMT.Fmt(id), &target, opts)
}

func (c *Client) GetAccountVehicles(id int, fields ...string) ([]types.VehicleStatsFrame, error) {
	opts := newDefaultRequestOptions()
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}

	var target []types.VehicleStatsFrame
	return target, c.sendRequest(RealmFromPlayerID(id), accountGetVehiclesEndpointFMT.Fmt(id), &target, opts)
}

func (c *Client) GetAccountAchievements(id int, fields ...string) (types.AchievementsFrame, error) {
	opts := newDefaultRequestOptions()
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}

	var target types.AchievementsFrame
	return target, c.sendRequest(RealmFromPlayerID(id), accountGetAchievementsEndpointFMT.Fmt(id), &target, opts)
}

func RealmFromPlayerID(id int) string {
	switch {
	case id == 0:
		return ""
	case id < 500000000:
		return "RU"
	case id < 1000000000:
		return "EU"
	case id < 2000000000:
		return "NA"
	default:
		return "AS"
	}
}
