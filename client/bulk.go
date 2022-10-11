package client

import (
	"strings"

	e "github.com/byvko-dev/am-types/errors/v2"
	"github.com/byvko-dev/am-types/wargaming/v1/accounts"
	"github.com/byvko-dev/am-types/wargaming/v1/clans"
	"github.com/byvko-dev/am-types/wargaming/v2/statistics"
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

// bulk.Get("/accounts/achievements", query.BulkAccountsAchievementsHandler)
func (c *Client) BulkGetAccountsAchievements(ids []string, realm string) (map[string]statistics.AchievementsFrame, *e.Error) {
	var target map[string]statistics.AchievementsFrame

	opts := defaultRequestOptions
	opts.Query.Add("ids", strings.Join(ids, ","))

	return target, c.sendRequest(realm, bulkAccountAchievementsEndpoint, &target, opts)
}

// bulk.Get("/accounts/clan", query.BulkAccountClanInfoHandler)
func (c *Client) BulkGetAccountsClans(ids []string, realm string) (map[string]clans.MemberProfile, *e.Error) {
	var target map[string]clans.MemberProfile

	opts := defaultRequestOptions
	opts.Query.Add("ids", strings.Join(ids, ","))

	return target, c.sendRequest(realm, bulkAccountClanInfoEndpoint, &target, opts)
}
