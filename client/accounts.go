package client

import (
	"errors"

	e "github.com/byvko-dev/am-types/errors/v2"
	"github.com/byvko-dev/am-types/wargaming/v1/accounts"
	"github.com/byvko-dev/am-types/wargaming/v1/clans"
	"github.com/byvko-dev/am-types/wargaming/v1/statistics"
)

func (c *Client) SearchAccounts(realm, query string, useSlowProxy bool) (accounts.BaseProfile, *e.Error) {
	opts := defaultRequestOptions
	opts.Query.Add("query", query)

	var target []accounts.BaseProfile
	err := c.sendRequest(realm, accountsSearchEndpoint, &target, opts)
	if err != nil {
		return accounts.BaseProfile{}, err
	}
	if len(target) == 0 {
		return accounts.BaseProfile{}, e.Input(errors.New("no results found"), "No results found")
	}
	return target[0], nil
}

func (c *Client) GetAccountByID(id int, useSlowProxy bool) (accounts.CompleteProfile, *e.Error) {
	var target accounts.CompleteProfile
	return target, c.sendRequest(realmFromPlayerID(id), accountsGetEndpointFMT.Fmt(id), &target, defaultRequestOptions)

}

func (c *Client) GetAccountClan(id int, useSlowProxy bool) (clans.MemberProfile, *e.Error) {
	var target clans.MemberProfile
	return target, c.sendRequest(realmFromPlayerID(id), accountClanGetEndpointFMT.Fmt(id), &target, defaultRequestOptions)
}

func (c *Client) GetAccountVehicles(id int, useSlowProxy bool) ([]statistics.VehicleStatsFrame, *e.Error) {
	var target []statistics.VehicleStatsFrame
	return target, c.sendRequest(realmFromPlayerID(id), accountGetVehiclesEndpointFMT.Fmt(id), &target, defaultRequestOptions)
}

func (c *Client) GetAccountAchievements(id int, useSlowProxy bool) (statistics.AchievementsFrame, *e.Error) {
	var target statistics.AchievementsFrame
	return target, c.sendRequest(realmFromPlayerID(id), accountGetAchievementsEndpointFMT.Fmt(id), &target, defaultRequestOptions)
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
		return "ASIA"
	}
}
