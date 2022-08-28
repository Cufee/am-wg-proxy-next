package accounts

import (
	"fmt"
	"strings"

	"aftermath.link/repo/am-wg-proxy/wargaming/client"
	"github.com/byvko-dev/am-types/wargaming/generic/api"
	"github.com/byvko-dev/am-types/wargaming/v1/statistics"
	"github.com/pkg/errors"
)

type AchievementsResponse struct {
	api.Response
	Data map[string]struct {
		Achievements statistics.AchievementsFrame `json:"achievements"`
	} `json:"data"`
}

func GetAccountAchievements(realm string, id string) (statistics.AchievementsFrame, error) {
	achievementsMap, err := GetBulkAccountsAchievements(realm, id)
	if err != nil {
		return statistics.AchievementsFrame{}, errors.Wrap(err, "GetAccountAchievements > GetBulkAccountsAchievements")
	}

	info, ok := achievementsMap[id]
	if !ok {
		return info, errors.Wrap(errors.New("account not found"), "GetAccountAchievements > GetBulkAccountsAchievements")
	}
	return info, nil
}

func GetBulkAccountsAchievements(realm string, ids ...string) (map[string]statistics.AchievementsFrame, error) {
	var response AchievementsResponse
	_, err := client.WargamingRequest(realm, fmt.Sprintf("account/achievements/?account_id=%s&fields=achievements", strings.Join(ids, ",")), "GET", nil, &response)
	if err != nil {
		return nil, errors.Wrap(err, "GetBulkAccountsAchievements > client.WargamingRequest")
	}
	if response.Error.Code != 0 {
		return nil, errors.Wrap(errors.New(response.Error.Message), "GetBulkAccountsAchievements > WargamingRequest")
	}

	// Get the right data
	achievementsMap := make(map[string]statistics.AchievementsFrame)
	for id, data := range response.Data {
		achievementsMap[id] = data.Achievements
	}
	return achievementsMap, nil
}
