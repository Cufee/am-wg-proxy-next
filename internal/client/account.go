package client

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type searchResponse struct {
	types.WgResponse
	Data []types.Account `json:"data"`
}

func (c *Client) SearchAccounts(realm, search string, fields ...string) ([]types.Account, error) {
	var response searchResponse
	query := url.Values{}
	if len(fields) > 0 {
		query.Set("fields", strings.Join(fields, ","))
	}
	query.Set("search", search)
	query.Set("limit", "3")

	_, err := c.Request(realm, fmt.Sprintf("account/list/?%s", query.Encode()), "GET", nil, &response)
	if err != nil {
		return nil, err
	}
	if response.Error.Code != 0 {
		log.Error().Str("realm", realm).Str("query", search).Msg("Error while searching accounts")
		return nil, errors.New(response.Error.Message)
	}

	return response.Data, nil
}

type infoResponse struct {
	types.WgResponse
	Data map[string]types.ExtendedAccount `json:"data"`
}

func (c *Client) AccountByID(realm string, id string, fields ...string) (types.ExtendedAccount, error) {
	accountsMap, err := c.BatchAccountByID(realm, []string{id}, fields...)
	if err != nil {
		return types.ExtendedAccount{}, err
	}

	info, ok := accountsMap[id]
	if !ok || info.ID == 0 {
		return types.ExtendedAccount{}, errors.New("account not found")
	}
	return info, nil
}

func (c *Client) BatchAccountByID(realm string, ids []string, fields ...string) (map[string]types.ExtendedAccount, error) {
	var response infoResponse
	query := url.Values{}
	if len(fields) > 0 {
		query.Set("fields", strings.Join(fields, ","))
	}
	query.Set("extra", "statistics.rating")
	query.Set("account_id", strings.Join(ids, ","))

	_, err := c.Request(realm, fmt.Sprintf("account/info/?%s", query.Encode()), "GET", nil, &response)
	if err != nil {
		return nil, err
	}
	if response.Error.Code != 0 {
		return nil, errors.New(response.Error.Message)
	}
	return response.Data, nil
}

type accountClanInfoResponse struct {
	types.WgResponse
	Data map[string]types.ClanMember `json:"data"`
}

func (c *Client) AccountClan(realm string, playerId string, fields ...string) (types.ClanMember, error) {
	data, err := c.BatchAccountClan(realm, []string{playerId}, fields...)
	if err != nil {
		return types.ClanMember{}, err
	}

	info, ok := data[fmt.Sprint(playerId)]
	if !ok || info.ClanID == 0 {
		return types.ClanMember{}, errors.New("clan not found")
	}
	return types.ClanMember{}, nil
}

func (c *Client) BatchAccountClan(realm string, ids []string, fields ...string) (map[string]types.ClanMember, error) {
	var response accountClanInfoResponse
	query := url.Values{}
	if len(fields) > 0 {
		query.Set("fields", strings.Join(fields, ","))
	}
	query.Set("extra", "clan")
	query.Set("account_id", strings.Join(ids, ","))

	_, err := c.Request(realm, fmt.Sprintf("clans/accountinfo/?%s", query.Encode()), "GET", nil, &response)
	if err != nil {
		return nil, err
	}

	if response.Error.Code != 0 {
		return nil, errors.New(response.Error.Message)
	}
	return response.Data, nil
}

type vehiclesResponse struct {
	types.WgResponse
	Data map[string][]types.VehicleStatsFrame `json:"data"`
}

// type vehicleAchievementsResponse struct {
// 	types.WgResponse
// 	Data map[string]types.AchievementsFrame `json:"data"`
// }

func (c *Client) AccountVehicles(realm string, id string, fields ...string) ([]types.VehicleStatsFrame, error) {
	var response vehiclesResponse
	query := url.Values{}
	if len(fields) > 0 {
		query.Set("fields", strings.Join(fields, ","))
	}
	query.Set("account_id", id)

	_, err := c.Request(realm, fmt.Sprintf("tanks/stats/?%s", query.Encode()), "GET", nil, &response)
	if err != nil {
		return nil, err
	}
	if response.Error.Code != 0 {
		return nil, errors.New(response.Error.Message)
	}

	info, ok := response.Data[id]
	if !ok {
		return info, errors.New("account not found")
	}
	return info, nil
}

type achievementsResponse struct {
	types.WgResponse
	Data map[string]struct {
		Achievements types.AchievementsFrame `json:"achievements"`
	} `json:"data"`
}

func (c *Client) AccountAchievements(realm string, id string, fields ...string) (types.AchievementsFrame, error) {

	achievementsMap, err := c.BatchAccountAchievements(realm, []string{id}, fields...)
	if err != nil {
		return types.AchievementsFrame{}, errors.Wrap(err, "GetAccountAchievements > GetBulkAccountsAchievements")
	}

	info, ok := achievementsMap[id]
	if !ok {
		return types.AchievementsFrame{}, errors.Wrap(errors.New("account not found"), "GetAccountAchievements > GetBulkAccountsAchievements")
	}
	return info, nil
}

func (c *Client) BatchAccountAchievements(realm string, ids []string, fields ...string) (map[string]types.AchievementsFrame, error) {
	var response achievementsResponse
	query := url.Values{}
	query.Set("fields", "achievements")
	if len(fields) > 0 {
		query.Set("fields", strings.Join(fields, ","))
	}
	query.Set("account_id", strings.Join(ids, ","))

	_, err := c.Request(realm, fmt.Sprintf("account/achievements/?%s", query.Encode()), "GET", nil, &response)
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
