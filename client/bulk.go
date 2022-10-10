package client

import (
	e "github.com/byvko-dev/am-types/errors/v2"
	"github.com/byvko-dev/am-types/wargaming/v1/accounts"
	"github.com/byvko-dev/am-types/wargaming/v1/clans"
)

// bulk.Get("/clans/info", query.BulkAccountsInfoHandler)
func (c *Client) BulkGetClansByID(ids []int, realm string) ([]clans.CompleteProfile, *e.Error) {
	var target []clans.CompleteProfile
	return target, c.sendRequest(realm, bulkAccountInfoEndpoint, &target, defaultRequestOptions)
}

// bulk.Get("/accounts/info", query.BulkAccountsInfoHandler)
func (c *Client) BulkGetAccountsByID(ids []int, realm string) ([]accounts.CompleteProfile, *e.Error) {
	var target []accounts.CompleteProfile
	return target, c.sendRequest(realm, bulkAccountInfoEndpoint, &target, defaultRequestOptions)
}

// bulk.Get("/accounts/vehicles", query.BulkAccountsVehiclesHandler)
// bulk.Get("/accounts/achievements", query.BulkAccountsAchievementsHandler)
