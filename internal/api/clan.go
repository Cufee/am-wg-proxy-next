package api

import (
	"errors"
	"strings"

	"github.com/cufee/am-wg-proxy-next/v2/types"
)

func (c *Client) ClanByID(realm string, id string, fields ...string) (types.ExtendedClan, error) {
	opts := newDefaultRequestOptions()
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}

	var target types.ExtendedClan
	return target, c.sendRequest(realm, clansGetEndpointFMT.Fmt(id), &target, opts)
}

// bulk.Get("/clans/info", query.BulkAccountsInfoHandler)
func (c *Client) BatchClanByID(realm string, ids []string, fields ...string) (map[string]types.ExtendedClan, error) {
	var target map[string]types.ExtendedClan

	opts := newDefaultRequestOptions()
	opts.Query.Add("ids", strings.Join(ids, ","))
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}

	return target, c.sendRequest(realm, bulkAccountInfoEndpoint, &target, opts)
}

func (c *Client) SearchClans(realm, query string, fields ...string) ([]types.Clan, error) {
	opts := newDefaultRequestOptions()
	opts.Query.Add("query", query)
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}

	var target []types.Clan
	err := c.sendRequest(realm, clansSearchEndpoint, &target, opts)
	if err != nil {
		return nil, err
	}
	if len(target) == 0 {
		return nil, errors.New("no results found")
	}
	return target, nil
}
