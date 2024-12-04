package api

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/cufee/am-wg-proxy-next/v2/types"
)

func (c *Client) SearchAccounts(ctx context.Context, realm types.Realm, query string, limit int, fields ...string) ([]types.Account, error) {
	opts := newDefaultRequestOptions()
	opts.Query.Add("query", query)
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}
	opts.Query.Set("limit", fmt.Sprint(limit))

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

func (c *Client) AccountByID(ctx context.Context, realm types.Realm, id string, fields ...string) (types.ExtendedAccount, error) {
	opts := newDefaultRequestOptions()
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}

	var target types.ExtendedAccount
	return target, c.sendRequest(ctx, realm, accountsGetEndpointFMT.Fmt(id), &target, opts)
}

// bulk.Get("/accounts/info", query.BulkAccountsInfoHandler)
func (c *Client) BatchAccountByID(ctx context.Context, realm types.Realm, ids []string, fields ...string) (map[string]types.ExtendedAccount, error) {
	var target map[string]types.ExtendedAccount

	opts := newDefaultRequestOptions()
	opts.Query.Add("ids", strings.Join(ids, ","))
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}

	return target, c.sendRequest(ctx, realm, bulkAccountInfoEndpoint, &target, opts)
}

func (c *Client) AccountClan(ctx context.Context, realm types.Realm, id string, fields ...string) (types.ClanMember, error) {
	opts := newDefaultRequestOptions()
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}

	var target types.ClanMember
	return target, c.sendRequest(ctx, realm, accountClanGetEndpointFMT.Fmt(id), &target, opts)
}

// bulk.Get("/accounts/clan", query.BulkAccountClanInfoHandler)
func (c *Client) BatchAccountClan(ctx context.Context, realm types.Realm, ids []string, fields ...string) (map[string]types.ClanMember, error) {
	var target map[string]types.ClanMember

	opts := newDefaultRequestOptions()
	opts.Query.Add("ids", strings.Join(ids, ","))
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}

	return target, c.sendRequest(ctx, realm, bulkAccountClanInfoEndpoint, &target, opts)
}

func (c *Client) AccountVehicles(ctx context.Context, realm types.Realm, id string, vehicles []string, fields ...string) ([]types.VehicleStatsFrame, error) {
	opts := newDefaultRequestOptions()
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}
	if len(vehicles) > 0 {
		opts.Query.Add("vehicles", strings.Join(vehicles, ","))
	}

	var target []types.VehicleStatsFrame
	return target, c.sendRequest(ctx, realm, accountGetVehiclesEndpointFMT.Fmt(id), &target, opts)
}

func (c *Client) AccountAchievements(ctx context.Context, realm types.Realm, id string, fields ...string) (types.AchievementsFrame, error) {
	opts := newDefaultRequestOptions()
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}

	var target types.AchievementsFrame
	return target, c.sendRequest(ctx, realm, accountGetAchievementsEndpointFMT.Fmt(id), &target, opts)
}

func (c *Client) AccountVehicleAchievements(ctx context.Context, realm types.Realm, id string, fields ...string) (map[string]types.AchievementsFrame, error) {
	opts := newDefaultRequestOptions()
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}

	var target map[string]types.AchievementsFrame
	return target, c.sendRequest(ctx, realm, accountGetVehicleAchievementsEndpointFMT.Fmt(id), &target, opts)
}

// bulk.Get("/accounts/achievements", query.BulkAccountsAchievementsHandler)
func (c *Client) BatchAccountAchievements(ctx context.Context, realm types.Realm, ids []string, fields ...string) (map[string]types.AchievementsFrame, error) {
	var target map[string]types.AchievementsFrame

	opts := newDefaultRequestOptions()
	opts.Query.Add("ids", strings.Join(ids, ","))
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}

	return target, c.sendRequest(ctx, realm, bulkAccountAchievementsEndpoint, &target, opts)
}
