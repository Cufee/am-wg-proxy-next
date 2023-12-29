package client

import (
	"errors"

	"github.com/cufee/am-wg-proxy-next/types"
)

func (c *Client) GetClanByID(realm string, id int) (types.ExtendedClan, error) {
	var target types.ExtendedClan
	return target, c.sendRequest(realm, clansGetEndpointFMT.Fmt(id), &target, newDefaultRequestOptions())
}

func (c *Client) SearchClans(realm, query string) (types.Clan, error) {
	opts := newDefaultRequestOptions()
	opts.Query.Add("query", query)

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
