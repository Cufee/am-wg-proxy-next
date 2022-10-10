package client

import (
	"strings"

	e "github.com/byvko-dev/am-types/errors/v2"
	"github.com/byvko-dev/am-types/wargaming/v1/accounts"
	"github.com/byvko-dev/am-types/wargaming/v1/clans"
)

// bulk.Get("/clans/info", query.BulkAccountsInfoHandler)
func (c *Client) BulkGetClansByID(ids []string, realm string) (map[string]clans.CompleteProfile, *e.Error) {
	var target map[string]clans.CompleteProfile

	opts := defaultRequestOptions
	opts.Query.Add("ids", strings.Join(ids, ","))

	return target, c.sendRequest(realm, bulkAccountInfoEndpoint, &target, opts)
}

// bulk.Get("/accounts/info", query.BulkAccountsInfoHandler)
func (c *Client) BulkGetAccountsByID(ids []string, realm string) (map[string]accounts.CompleteProfile, *e.Error) {
	var target map[string]accounts.CompleteProfile

	opts := defaultRequestOptions
	opts.Query.Add("ids", strings.Join(ids, ","))

	return target, c.sendRequest(realm, bulkAccountInfoEndpoint, &target, opts)
}

// bulk.Get("/accounts/vehicles", query.BulkAccountsVehiclesHandler)
// bulk.Get("/accounts/achievements", query.BulkAccountsAchievementsHandler)
