package client

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/pkg/errors"
)

func (c *Client) SearchAccounts(ctx context.Context, realm types.Realm, search string, opts ...types.Option) ([]types.Account, error) {
	var response types.WgResponse[[]types.Account]
	options := types.GetOptions(opts...)
	query := options.Query()
	query.Set("search", search)

	_, err := c.Request(ctx, realm, fmt.Sprintf("account/list/?%s", query.Encode()), "GET", nil, &response)
	if err != nil {
		return nil, err
	}
	if response.Error.Code != 0 {
		c.logger.Error().Str("realm", realm.String()).Str("query", search).Msg("Error while searching accounts")
		return nil, errors.New(response.Error.Message)
	}

	return response.Data, nil
}

func (c *Client) AccountByID(ctx context.Context, realm types.Realm, id string, opts ...types.Option) (types.ExtendedAccount, error) {
	accountsMap, err := c.BatchAccountByID(ctx, realm, []string{id}, opts...)
	if err != nil {
		return types.ExtendedAccount{}, err
	}

	info, ok := accountsMap[id]
	if !ok || info.ID == 0 {
		return types.ExtendedAccount{}, errors.New("account not found")
	}
	return info, nil
}

func (c *Client) BatchAccountByID(ctx context.Context, realm types.Realm, ids []string, opts ...types.Option) (map[string]types.ExtendedAccount, error) {
	var response types.WgResponse[map[string]types.ExtendedAccount]
	options := types.GetOptions(opts...)
	options.Extra = append(options.Extra, "statistics.rating")
	query := options.Query()
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

func (c *Client) AccountClan(ctx context.Context, realm types.Realm, playerId string, opts ...types.Option) (types.ClanMember, error) {
	data, err := c.BatchAccountClan(ctx, realm, []string{playerId}, opts...)
	if err != nil {
		return types.ClanMember{}, err
	}

	info, ok := data[fmt.Sprint(playerId)]
	if !ok || info.ClanID == 0 {
		return types.ClanMember{}, errors.New("clan not found")
	}
	return info, nil
}

func (c *Client) BatchAccountClan(ctx context.Context, realm types.Realm, ids []string, opts ...types.Option) (map[string]types.ClanMember, error) {
	var response types.WgResponse[map[string]types.ClanMember]
	options := types.GetOptions(opts...)
	options.Extra = append(options.Extra, "clan")
	query := options.Query()
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

func (c *Client) AccountVehicles(ctx context.Context, realm types.Realm, id string, vehicles []string, opts ...types.Option) ([]types.VehicleStatsFrame, error) {
	var response types.WgResponse[map[string][]types.VehicleStatsFrame]
	options := types.GetOptions(opts...)
	query := options.Query()
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

func (c *Client) AccountAchievements(ctx context.Context, realm types.Realm, id string, opts ...types.Option) (types.AchievementsFrame, error) {
	achievementsMap, err := c.BatchAccountAchievements(ctx, realm, []string{id}, opts...)
	if err != nil {
		return types.AchievementsFrame{}, errors.Wrap(err, "GetAccountAchievements > GetBulkAccountsAchievements")
	}

	info, ok := achievementsMap[id]
	if !ok {
		return types.AchievementsFrame{}, errors.Wrap(errors.New("account not found"), "GetAccountAchievements > GetBulkAccountsAchievements")
	}
	return info, nil
}

func (c *Client) AccountVehicleAchievements(ctx context.Context, realm types.Realm, id string, opts ...types.Option) (map[string]types.AchievementsFrame, error) {
	var response types.WgResponse[map[string][]struct {
		Achievements types.AchievementsFrame `json:"achievements"`
		AccountID    int                     `json:"account_id"`
		TankID       int                     `json:"tank_id"`
	}]
	options := types.GetOptions(opts...)
	query := options.Query()
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

func (c *Client) BatchAccountAchievements(ctx context.Context, realm types.Realm, ids []string, opts ...types.Option) (map[string]types.AchievementsFrame, error) {
	var response types.WgResponse[map[string]struct {
		Achievements types.AchievementsFrame `json:"achievements"`
	}]
	options := types.GetOptions(opts...)
	options.Fields = append(options.Fields, "achievements")
	query := options.Query()
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
