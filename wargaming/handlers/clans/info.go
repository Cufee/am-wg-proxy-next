package clans

import (
	"errors"
	"fmt"

	"aftermath.link/repo/am-wg-proxy/wargaming/client"
	"aftermath.link/repo/am-wg-proxy/wargaming/handlers"
)

type ClanInfoResponse struct {
	handlers.WargamingBaseResponse
	Data map[string]ClanInfo `json:"data"`
}

type ClanInfo struct {
	RecruitingOptions RecruitingOptions `json:"recruiting_options"`
	MembersCount      int64             `json:"members_count"`
	Name              string            `json:"name"`
	CreatorName       string            `json:"creator_name"`
	ClanID            int64             `json:"clan_id"`
	CreatedAt         int64             `json:"created_at"`
	UpdatedAt         int64             `json:"updated_at"`
	LeaderName        string            `json:"leader_name"`
	MembersIDS        []int64           `json:"members_ids"`
	RecruitingPolicy  string            `json:"recruiting_policy"`
	Tag               string            `json:"tag"`
	IsClanDisbanded   bool              `json:"is_clan_disbanded"`
	OldName           string            `json:"old_name"`
	EmblemSetID       int64             `json:"emblem_set_id"`
	CreatorID         int64             `json:"creator_id"`
	Motto             string            `json:"motto"`
	RenamedAt         int64             `json:"renamed_at"`
	OldTag            string            `json:"old_tag"`
	LeaderID          int64             `json:"leader_id"`
	Description       string            `json:"description"`
}

type RecruitingOptions struct {
	VehiclesLevel        int64 `json:"vehicles_level"`
	WINSRatio            int64 `json:"wins_ratio"`
	AverageBattlesPerDay int64 `json:"average_battles_per_day"`
	Battles              int64 `json:"battles"`
	AverageDamage        int64 `json:"average_damage"`
}

func GetClanInfo(bucket, realm string, clanId int) (ClanInfo, error) {
	var response ClanInfoResponse
	_, err := client.WargamingRequest(bucket, realm, fmt.Sprintf("clans/info/?clan_id=%v", clanId), "GET", nil, &response)
	if err != nil {
		return ClanInfo{}, err
	}
	if response.Error.Code != 0 {
		return ClanInfo{}, errors.New(response.Error.Message)
	}

	info, ok := response.Data[fmt.Sprint(clanId)]
	if !ok {
		return info, errors.New("account not found")
	}
	return info, nil
}
