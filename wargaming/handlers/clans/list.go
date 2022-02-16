package clans

import (
	"errors"
	"fmt"

	"aftermath.link/repo/am-wg-proxy/wargaming/client"
	"aftermath.link/repo/am-wg-proxy/wargaming/handlers"
)

type ClanSearchResponse struct {
	handlers.WargamingBaseResponse
	Data []BasicClanDetails `json:"data"`
}

type BasicClanDetails struct {
	MembersCount int64  `json:"members_count"`
	Name         string `json:"name"`
	CreatedAt    int64  `json:"created_at"`
	Tag          string `json:"tag"`
	ClanID       int64  `json:"clan_id"`
}

func SearchClans(bucket, realm, search string) ([]BasicClanDetails, error) {
	var response ClanSearchResponse
	_, err := client.WargamingRequest(bucket, realm, fmt.Sprintf("clans/list/?search=%v&limit=3", search), "GET", nil, &response)
	if err != nil {
		return nil, err
	}
	if response.Error.Code != 0 {
		return nil, errors.New(response.Error.Message)
	}
	return response.Data, nil
}
