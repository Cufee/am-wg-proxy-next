package client

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/cufee/am-wg-proxy-next/v2/types"
)

func (c *Client) SearchClans(ctx context.Context, realm types.Realm, search string, opts ...Option) ([]types.Clan, error) {
	var response types.WgResponse[[]types.Clan]
	options := GetOptions(opts...)
	query := options.Query()
	query.Set("search", search)

	_, err := c.Request(ctx, realm, fmt.Sprintf("clans/list/?%v", query.Encode()), "GET", nil, &response)
	if err != nil {
		return nil, err
	}
	if response.Error.Code != 0 {
		return nil, errors.New(response.Error.Message)
	}
	return response.Data, nil
}

func (c *Client) ClanByID(ctx context.Context, realm types.Realm, clanId string, opts ...Option) (types.ExtendedClan, error) {
	data, err := c.BatchClanByID(ctx, realm, []string{clanId}, opts...)
	if err != nil {
		return types.ExtendedClan{}, err
	}

	info, ok := data[clanId]
	if !ok || info.ID == 0 {
		return types.ExtendedClan{}, errors.New("clan not found")
	}
	return info, nil
}

func (c *Client) BatchClanByID(ctx context.Context, realm types.Realm, ids []string, opts ...Option) (map[string]types.ExtendedClan, error) {
	var response types.WgResponse[map[string]types.ExtendedClan]
	options := GetOptions(opts...)
	query := options.Query()
	query.Set("clan_id", strings.Join(ids, ","))

	_, err := c.Request(ctx, realm, fmt.Sprintf("clans/info/?%s", query.Encode()), "GET", nil, &response)
	if err != nil {
		return nil, err
	}
	if response.Error.Code != 0 {
		return nil, errors.New(response.Error.Message)
	}

	return response.Data, nil
}
