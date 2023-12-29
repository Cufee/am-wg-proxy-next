package client

import (
	"strings"

	"github.com/cufee/am-wg-proxy-next/types"
)

// bulk.Get("/clans/info", query.BulkAccountsInfoHandler)
func (c *Client) BulkGetClansByID(ids []string, realm string) (map[string]types.ExtendedClan, error) {
	var target map[string]types.ExtendedClan

	opts := newDefaultRequestOptions()
	opts.Query.Add("ids", strings.Join(ids, ","))

	return target, c.sendRequest(realm, bulkAccountInfoEndpoint, &target, opts)
}

// bulk.Get("/accounts/info", query.BulkAccountsInfoHandler)
func (c *Client) BulkGetAccountsByID(ids []string, realm string) (map[string]types.ExtendedAccount, error) {
	var target map[string]types.ExtendedAccount

	opts := newDefaultRequestOptions()
	opts.Query.Add("ids", strings.Join(ids, ","))

	return target, c.sendRequest(realm, bulkAccountInfoEndpoint, &target, opts)
}

// bulk.Get("/accounts/achievements", query.BulkAccountsAchievementsHandler)
func (c *Client) BulkGetAccountsAchievements(ids []string, realm string) (map[string]types.AchievementsFrame, error) {
	var target map[string]types.AchievementsFrame

	opts := newDefaultRequestOptions()
	opts.Query.Add("ids", strings.Join(ids, ","))

	return target, c.sendRequest(realm, bulkAccountAchievementsEndpoint, &target, opts)
}

// bulk.Get("/accounts/clan", query.BulkAccountClanInfoHandler)
func (c *Client) BulkGetAccountsClans(ids []string, realm string) (map[string]types.ClanMember, error) {
	var target map[string]types.ClanMember

	opts := newDefaultRequestOptions()
	opts.Query.Add("ids", strings.Join(ids, ","))

	return target, c.sendRequest(realm, bulkAccountClanInfoEndpoint, &target, opts)
}
