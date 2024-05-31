package client

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/cufee/am-wg-proxy-next/v2/types"
)

type clanSearchResponse struct {
	types.WgResponse
	Data []types.Clan `json:"data"`
}

func (c *Client) SearchClans(ctx context.Context, realm, search string, fields ...string) ([]types.Clan, error) {
	var response clanSearchResponse
	query := url.Values{}
	if len(fields) > 0 {
		query.Set("fields", strings.Join(fields, ","))
	}
	query.Set("search", search)
	query.Set("limit", "3")

	_, err := c.Request(ctx, realm, fmt.Sprintf("clans/list/?%v", query.Encode()), "GET", nil, &response)
	if err != nil {
		return nil, err
	}
	if response.Error.Code != 0 {
		return nil, errors.New(response.Error.Message)
	}
	return response.Data, nil
}

type clanInfoResponse struct {
	types.WgResponse
	Data map[string]types.ExtendedClan `json:"data"`
}

func (c *Client) ClanByID(ctx context.Context, realm string, clanId string, fields ...string) (types.ExtendedClan, error) {
	data, err := c.BatchClanByID(ctx, realm, []string{clanId}, fields...)
	if err != nil {
		return types.ExtendedClan{}, err
	}

	info, ok := data[clanId]
	if !ok || info.ID == 0 {
		return types.ExtendedClan{}, errors.New("clan not found")
	}
	return info, nil
}

func (c *Client) BatchClanByID(ctx context.Context, realm string, ids []string, fields ...string) (map[string]types.ExtendedClan, error) {
	var response clanInfoResponse
	query := url.Values{}
	if len(fields) > 0 {
		query.Set("fields", strings.Join(fields, ","))
	}
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
