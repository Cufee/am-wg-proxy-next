package client

import (
	e "github.com/byvko-dev/am-types/errors/v2"
	"github.com/byvko-dev/am-types/wargaming/v1/clans"
)

func (c *client) GetClanByID(realm string, id int, useSlowProxy bool) (clans.CompleteProfile, *e.Error) {
	var target clans.CompleteProfile
	return target, c.sendRequest(realm, clansGetEndpointFMT.Fmt(id), &target, defaultRequestOptions)
}

func (c *client) SearchClans(realm, query string, useSlowProxy bool) (clans.BasicProfile, *e.Error) {
	opts := defaultRequestOptions
	opts.Query.Add("query", query)

	var target []clans.BasicProfile
	err := c.sendRequest(realm, clansSearchEndpoint, &target, opts)
	if err != nil {
		return clans.BasicProfile{}, err
	}
	if len(target) == 0 {
		return clans.BasicProfile{}, e.Input(nil, "No results found")
	}
	return target[0], nil
}
