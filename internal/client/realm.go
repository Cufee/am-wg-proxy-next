package client

import (
	"github.com/cufee/am-wg-proxy-next/v2/client/common"
	"github.com/cufee/am-wg-proxy-next/v2/internal/utils"
	"github.com/cufee/am-wg-proxy-next/v2/types"
)

func (c *Client) RealmFromID(id string) (*types.Realm, error) {
	realm := utils.RealmFromID(id)
	if realm == nil {
		return nil, common.ErrInvalidAccountID
	}
	return realm, nil
}
