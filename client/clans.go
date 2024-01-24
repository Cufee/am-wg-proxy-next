package client

import (
	"errors"
	"strings"

	"github.com/cufee/am-wg-proxy-next/types"
)

func (c *Client) GetClanByID(realm string, id int, fields ...string) (types.ExtendedClan, error) {
	opts := newDefaultRequestOptions()
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}

	var target types.ExtendedClan
	return target, c.sendRequest(realm, clansGetEndpointFMT.Fmt(id), &target, opts)
}

func (c *Client) SearchClans(realm, query string, fields ...string) (types.Clan, error) {
	opts := newDefaultRequestOptions()
	opts.Query.Add("query", query)
	if len(fields) > 0 {
		opts.Query.Add("fields", strings.Join(fields, ","))
	}

	var target []types.Clan
	err := c.sendRequest(realm, clansSearchEndpoint, &target, opts)
	if err != nil {
		return types.Clan{}, err
	}
	if len(target) == 0 {
		return types.Clan{}, errors.New("no results found")
	}
	return target[0], nil
}
