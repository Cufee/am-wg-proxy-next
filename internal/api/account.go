package api

import (
	"errors"
	"strings"

	"github.com/cufee/am-wg-proxy-next/v2/types"
)

func (c *Client) SearchAccounts(realm, query string, fields ...string) ([]types.Account, error) {
	opts := newDefaultRequestOptions()
	opts.Query.Add("query", query)
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}

	var target []types.Account
	err := c.sendRequest(realm, accountsSearchEndpoint, &target, opts)
	if err != nil {
		return nil, err
	}
	if len(target) == 0 {
		return nil, errors.New("no results found")
	}
	return target, nil
}

func (c *Client) AccountByID(realm string, id string, fields ...string) (types.ExtendedAccount, error) {
	opts := newDefaultRequestOptions()
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}

	var target types.ExtendedAccount
	return target, c.sendRequest(realm, accountsGetEndpointFMT.Fmt(id), &target, opts)
}

// bulk.Get("/accounts/info", query.BulkAccountsInfoHandler)
func (c *Client) BatchAccountByID(realm string, ids []string, fields ...string) (map[string]types.ExtendedAccount, error) {
	var target map[string]types.ExtendedAccount

	opts := newDefaultRequestOptions()
	opts.Query.Add("ids", strings.Join(ids, ","))
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}

	return target, c.sendRequest(realm, bulkAccountInfoEndpoint, &target, opts)
}

func (c *Client) AccountClan(realm string, id string, fields ...string) (types.ClanMember, error) {
	opts := newDefaultRequestOptions()
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}

	var target types.ClanMember
	return target, c.sendRequest(realm, accountClanGetEndpointFMT.Fmt(id), &target, opts)
}

// bulk.Get("/accounts/clan", query.BulkAccountClanInfoHandler)
func (c *Client) BatchAccountClan(realm string, ids []string, fields ...string) (map[string]types.ClanMember, error) {
	var target map[string]types.ClanMember

	opts := newDefaultRequestOptions()
	opts.Query.Add("ids", strings.Join(ids, ","))
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}

	return target, c.sendRequest(realm, bulkAccountClanInfoEndpoint, &target, opts)
}

func (c *Client) AccountVehicles(realm string, id string, fields ...string) ([]types.VehicleStatsFrame, error) {
	opts := newDefaultRequestOptions()
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}

	var target []types.VehicleStatsFrame
	return target, c.sendRequest(realm, accountGetVehiclesEndpointFMT.Fmt(id), &target, opts)
}

func (c *Client) AccountAchievements(realm string, id string, fields ...string) (types.AchievementsFrame, error) {
	opts := newDefaultRequestOptions()
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}

	var target types.AchievementsFrame
	return target, c.sendRequest(realm, accountGetAchievementsEndpointFMT.Fmt(id), &target, opts)
}

// bulk.Get("/accounts/achievements", query.BulkAccountsAchievementsHandler)
func (c *Client) BatchAccountAchievements(realm string, ids []string, fields ...string) (map[string]types.AchievementsFrame, error) {
	var target map[string]types.AchievementsFrame

	opts := newDefaultRequestOptions()
	opts.Query.Add("ids", strings.Join(ids, ","))
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}

	return target, c.sendRequest(realm, bulkAccountAchievementsEndpoint, &target, opts)
}
