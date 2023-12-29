package client

import (
	"errors"

	"github.com/cufee/am-wg-proxy-next/types"
)

func (c *Client) SearchAccounts(realm, query string) (types.Account, error) {
	opts := newDefaultRequestOptions()
	opts.Query.Add("query", query)

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

func (c *Client) GetAccountByID(id int) (types.ExtendedAccount, error) {
	var target types.ExtendedAccount
	return target, c.sendRequest(realmFromPlayerID(id), accountsGetEndpointFMT.Fmt(id), &target, newDefaultRequestOptions())
}

func (c *Client) GetAccountClan(id int) (types.ClanMember, error) {
	var target types.ClanMember
	return target, c.sendRequest(realmFromPlayerID(id), accountClanGetEndpointFMT.Fmt(id), &target, newDefaultRequestOptions())
}

func (c *Client) GetAccountVehicles(id int) ([]types.VehicleStatsFrame, error) {
	var target []types.VehicleStatsFrame
	return target, c.sendRequest(realmFromPlayerID(id), accountGetVehiclesEndpointFMT.Fmt(id), &target, newDefaultRequestOptions())
}

func (c *Client) GetAccountAchievements(id int) (types.AchievementsFrame, error) {
	var target types.AchievementsFrame
	return target, c.sendRequest(realmFromPlayerID(id), accountGetAchievementsEndpointFMT.Fmt(id), &target, newDefaultRequestOptions())
}

func realmFromPlayerID(id int) string {
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
