package client

import "github.com/cufee/am-wg-proxy-next/types"

func (c *Client) NewAuthURL(realm, lang, redirectURL string) (string, error) {
	opts := newDefaultRequestOptions()
	opts.Query.Add("language", types.GetLocale(lang))

	var target struct {
		Location string `json:"location"`
	}
	return target.Location, c.sendRequest(realm, openIDLoginEndpoint, &target, opts)
}
