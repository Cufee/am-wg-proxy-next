package api

import (
	"github.com/cufee/am-wg-proxy-next/v2/client/common"
	"github.com/cufee/am-wg-proxy-next/v2/types"
)

func (c *Client) RealmFromID(id string) (types.Realm, bool) {
	return common.RealmFromID(id)
}

func (c *Client) ParseRealm(value string) (types.Realm, bool) {
	return common.ParseRealm(value)
}
