package api

import (
	"context"
	"errors"
	"strings"

	"github.com/cufee/am-wg-proxy-next/v2/types"
)

func (c *Client) ClanByID(ctx context.Context, realm types.Realm, id string, options ...types.Option) (types.ExtendedClan, error) {
	opts := newDefaultRequestOptions(options)
	var target types.ExtendedClan
	return target, c.sendRequest(ctx, realm, clansGetEndpointFMT.Fmt(id), &target, opts)
}

// bulk.Get("/clans/info", query.BulkAccountsInfoHandler)
func (c *Client) BatchClanByID(ctx context.Context, realm types.Realm, ids []string, options ...types.Option) (map[string]types.ExtendedClan, error) {
	var target map[string]types.ExtendedClan

	opts := newDefaultRequestOptions(options)
	opts.Query.Add("ids", strings.Join(ids, ","))
	return target, c.sendRequest(ctx, realm, bulkAccountInfoEndpoint, &target, opts)
}

func (c *Client) SearchClans(ctx context.Context, realm types.Realm, query string, options ...types.Option) ([]types.Clan, error) {
	opts := newDefaultRequestOptions(options)
	opts.Query.Add("query", query)

	var target []types.Clan
	err := c.sendRequest(ctx, realm, clansSearchEndpoint, &target, opts)
	if err != nil {
		return nil, err
	}
	if len(target) == 0 {
		return nil, errors.New("no results found")
	}
	return target, nil
}
