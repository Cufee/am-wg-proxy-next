package accounts

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/cufee/am-wg-proxy-next/internal/wargaming/client"
	"github.com/cufee/am-wg-proxy-next/types"
	"github.com/pkg/errors"
)

type achievementsResponse struct {
	types.WgResponse
	Data map[string]struct {
		Achievements types.AchievementsFrame `json:"achievements"`
	} `json:"data"`
}

func GetAccountAchievements(realm string, id string, fields ...string) (*types.AchievementsFrame, error) {
	achievementsMap, err := GetBulkAccountsAchievements(realm, []string{id}, fields...)
	if err != nil {
		return nil, errors.Wrap(err, "GetAccountAchievements > GetBulkAccountsAchievements")
	}

	info, ok := achievementsMap[id]
	if !ok {
		return nil, errors.Wrap(errors.New("account not found"), "GetAccountAchievements > GetBulkAccountsAchievements")
	}
	return &info, nil
}

func GetBulkAccountsAchievements(realm string, ids []string, fields ...string) (map[string]types.AchievementsFrame, error) {
	var response achievementsResponse
	query := url.Values{}
	query.Set("fields", "achievements")
	if len(fields) > 0 {
		query.Set("fields", strings.Join(fields, ","))
	}
	query.Set("account_id", strings.Join(ids, ","))

	_, err := client.WargamingRequest(realm, fmt.Sprintf("account/achievements/?%s", query.Encode()), "GET", nil, &response)
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
