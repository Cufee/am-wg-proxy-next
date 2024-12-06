package api

import (
	"context"
	"errors"
	"strings"

	"github.com/cufee/am-wg-proxy-next/v2/types"
)

func (c *Client) SearchAccounts(ctx context.Context, realm types.Realm, query string, options ...types.Option) ([]types.Account, error) {
	opts := newDefaultRequestOptions(options)
	opts.Query.Add("query", query)

	var target []types.Account
	err := c.sendRequest(ctx, realm, accountsSearchEndpoint, &target, opts)
	if err != nil {
		return nil, err
	}
	if len(target) == 0 {
		return nil, errors.New("no results found")
	}
	return target, nil
}

func (c *Client) AccountByID(ctx context.Context, realm types.Realm, id string, options ...types.Option) (types.ExtendedAccount, error) {
	opts := newDefaultRequestOptions(options)
	var target types.ExtendedAccount
	return target, c.sendRequest(ctx, realm, accountsGetEndpointFMT.Fmt(id), &target, opts)
}

// bulk.Get("/accounts/info", query.BulkAccountsInfoHandler)
func (c *Client) BatchAccountByID(ctx context.Context, realm types.Realm, ids []string, options ...types.Option) (map[string]types.ExtendedAccount, error) {
	var target map[string]types.ExtendedAccount
	opts := newDefaultRequestOptions(options)
	opts.Query.Add("ids", strings.Join(ids, ","))
	return target, c.sendRequest(ctx, realm, bulkAccountInfoEndpoint, &target, opts)
}

func (c *Client) AccountClan(ctx context.Context, realm types.Realm, id string, options ...types.Option) (types.ClanMember, error) {
	opts := newDefaultRequestOptions(options)
	var target types.ClanMember
	return target, c.sendRequest(ctx, realm, accountClanGetEndpointFMT.Fmt(id), &target, opts)
}

// bulk.Get("/accounts/clan", query.BulkAccountClanInfoHandler)
func (c *Client) BatchAccountClan(ctx context.Context, realm types.Realm, ids []string, options ...types.Option) (map[string]types.ClanMember, error) {
	var target map[string]types.ClanMember

	opts := newDefaultRequestOptions(options)
	opts.Query.Add("ids", strings.Join(ids, ","))
	return target, c.sendRequest(ctx, realm, bulkAccountClanInfoEndpoint, &target, opts)
}

func (c *Client) AccountVehicles(ctx context.Context, realm types.Realm, id string, vehicles []string, options ...types.Option) ([]types.VehicleStatsFrame, error) {
	opts := newDefaultRequestOptions(options)
	if len(vehicles) > 0 {
		opts.Query.Add("vehicles", strings.Join(vehicles, ","))
	}
	var target []types.VehicleStatsFrame
	return target, c.sendRequest(ctx, realm, accountGetVehiclesEndpointFMT.Fmt(id), &target, opts)
}

func (c *Client) AccountAchievements(ctx context.Context, realm types.Realm, id string, options ...types.Option) (types.AchievementsFrame, error) {
	opts := newDefaultRequestOptions(options)
	var target types.AchievementsFrame
	return target, c.sendRequest(ctx, realm, accountGetAchievementsEndpointFMT.Fmt(id), &target, opts)
}

func (c *Client) AccountVehicleAchievements(ctx context.Context, realm types.Realm, id string, options ...types.Option) (map[string]types.AchievementsFrame, error) {
	opts := newDefaultRequestOptions(options)
	var target map[string]types.AchievementsFrame
	return target, c.sendRequest(ctx, realm, accountGetVehicleAchievementsEndpointFMT.Fmt(id), &target, opts)
}

// bulk.Get("/accounts/achievements", query.BulkAccountsAchievementsHandler)
func (c *Client) BatchAccountAchievements(ctx context.Context, realm types.Realm, ids []string, options ...types.Option) (map[string]types.AchievementsFrame, error) {
	var target map[string]types.AchievementsFrame
	opts := newDefaultRequestOptions(options)
	opts.Query.Add("ids", strings.Join(ids, ","))
	return target, c.sendRequest(ctx, realm, bulkAccountAchievementsEndpoint, &target, opts)
}
