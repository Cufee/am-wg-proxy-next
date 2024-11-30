package client

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/pkg/errors"
)

func (c *Client) SearchAccounts(ctx context.Context, realm, search string, limit int, fields ...string) ([]types.Account, error) {
	var response types.WgResponse[[]types.Account]
	query := url.Values{}
	if len(fields) > 0 {
		query.Set("fields", strings.Join(fields, ","))
	}
	query.Set("search", search)
	query.Set("limit", fmt.Sprint(limit))

	_, err := c.Request(ctx, realm, fmt.Sprintf("account/list/?%s", query.Encode()), "GET", nil, &response)
	if err != nil {
		return nil, err
	}
	if response.Error.Code != 0 {
		c.logger.Error().Str("realm", realm).Str("query", search).Msg("Error while searching accounts")
		return nil, errors.New(response.Error.Message)
	}

	return response.Data, nil
}

func (c *Client) AccountByID(ctx context.Context, realm string, id string, fields ...string) (types.ExtendedAccount, error) {
	accountsMap, err := c.BatchAccountByID(ctx, realm, []string{id}, fields...)
	if err != nil {
		return types.ExtendedAccount{}, err
	}

	info, ok := accountsMap[id]
	if !ok || info.ID == 0 {
		return types.ExtendedAccount{}, errors.New("account not found")
	}
	return info, nil
}

func (c *Client) BatchAccountByID(ctx context.Context, realm string, ids []string, fields ...string) (map[string]types.ExtendedAccount, error) {
	var response types.WgResponse[map[string]types.ExtendedAccount]
	query := url.Values{}
	if len(fields) > 0 {
		query.Set("fields", strings.Join(fields, ","))
	}
	query.Set("extra", "statistics.rating")
	query.Set("account_id", strings.Join(ids, ","))

	_, err := c.Request(ctx, realm, fmt.Sprintf("account/info/?%s", query.Encode()), "GET", nil, &response)
	if err != nil {
		return nil, err
	}
	if response.Error.Code != 0 {
		return nil, errors.New(response.Error.Message)
	}
	return response.Data, nil
}

func (c *Client) AccountClan(ctx context.Context, realm string, playerId string, fields ...string) (types.ClanMember, error) {
	data, err := c.BatchAccountClan(ctx, realm, []string{playerId}, fields...)
	if err != nil {
		return types.ClanMember{}, err
	}

	info, ok := data[fmt.Sprint(playerId)]
	if !ok || info.ClanID == 0 {
		return types.ClanMember{}, errors.New("clan not found")
	}
	return info, nil
}

func (c *Client) BatchAccountClan(ctx context.Context, realm string, ids []string, fields ...string) (map[string]types.ClanMember, error) {
	var response types.WgResponse[map[string]types.ClanMember]
	query := url.Values{}
	if len(fields) > 0 {
		query.Set("fields", strings.Join(fields, ","))
	}
	query.Set("extra", "clan")
	query.Set("account_id", strings.Join(ids, ","))

	_, err := c.Request(ctx, realm, fmt.Sprintf("clans/accountinfo/?%s", query.Encode()), "GET", nil, &response)
	if err != nil {
		return nil, err
	}

	if response.Error.Code != 0 {
		return nil, errors.New(response.Error.Message)
	}
	return response.Data, nil
}

func (c *Client) AccountVehicles(ctx context.Context, realm string, id string, vehicles []string, fields ...string) ([]types.VehicleStatsFrame, error) {
	var response types.WgResponse[map[string][]types.VehicleStatsFrame]
	query := url.Values{}
	if len(fields) > 0 {
		query.Set("fields", strings.Join(fields, ","))
	}
	query.Set("account_id", id)
	if len(vehicles) > 0 {
		query.Set("tank_id", strings.Join(vehicles, ","))
	}

	_, err := c.Request(ctx, realm, fmt.Sprintf("tanks/stats/?%s", query.Encode()), "GET", nil, &response)
	if err != nil {
		return nil, err
	}
	if response.Error.Code != 0 {
		return nil, errors.New(response.Error.Message)
	}

	info, ok := response.Data[id]
	if !ok {
		return nil, errors.New("account not found")
	}
	return info, nil
}

func (c *Client) AccountAchievements(ctx context.Context, realm string, id string, fields ...string) (types.AchievementsFrame, error) {
	achievementsMap, err := c.BatchAccountAchievements(ctx, realm, []string{id}, fields...)
	if err != nil {
		return types.AchievementsFrame{}, errors.Wrap(err, "GetAccountAchievements > GetBulkAccountsAchievements")
	}

	info, ok := achievementsMap[id]
	if !ok {
		return types.AchievementsFrame{}, errors.Wrap(errors.New("account not found"), "GetAccountAchievements > GetBulkAccountsAchievements")
	}
	return info, nil
}

func (c *Client) AccountVehicleAchievements(ctx context.Context, realm string, id string, fields ...string) (map[string]types.AchievementsFrame, error) {
	var response types.WgResponse[map[string][]struct {
		Achievements types.AchievementsFrame `json:"achievements"`
		AccountID    int                     `json:"account_id"`
		TankID       int                     `json:"tank_id"`
	}]
	query := url.Values{}
	if len(fields) > 0 {
		query.Set("fields", strings.Join(fields, ","))
	}
	query.Set("account_id", id)

	_, err := c.Request(ctx, realm, fmt.Sprintf("tanks/achievements/?%s", query.Encode()), "GET", nil, &response)
	if err != nil {
		return nil, err
	}
	if response.Error.Code != 0 {
		return nil, errors.New(response.Error.Message)
	}

	vehicles, ok := response.Data[id]
	if !ok {
		return nil, errors.New("account not found")
	}

	achievements := make(map[string]types.AchievementsFrame)
	for _, vehicle := range vehicles {
		achievements[strconv.Itoa(vehicle.TankID)] = vehicle.Achievements
	}
	return achievements, nil
}

func (c *Client) BatchAccountAchievements(ctx context.Context, realm string, ids []string, fields ...string) (map[string]types.AchievementsFrame, error) {
	var response types.WgResponse[map[string]struct {
		Achievements types.AchievementsFrame `json:"achievements"`
	}]
	query := url.Values{}
	query.Set("fields", "achievements")
	if len(fields) > 0 {
		query.Set("fields", strings.Join(fields, ","))
	}
	query.Set("account_id", strings.Join(ids, ","))

	_, err := c.Request(ctx, realm, fmt.Sprintf("account/achievements/?%s", query.Encode()), "GET", nil, &response)
	if err != nil {
		return nil, errors.Wrap(err, "GetBulkAccountsAchievements > client.WargamingRequest")
	}
	if response.Error.Code != 0 {
		return nil, errors.Wrap(errors.New(response.Error.Message), "GetBulkAccountsAchievements > WargamingRequest")
	}

	// Get the right data
	achievementsMap := make(map[string]types.AchievementsFrame)
	for id, data := range response.Data {
		achievementsMap[id] = data.Achievements
	}
	return achievementsMap, nil
}
