package clans

import (
	"github.com/byvko-dev/am-types/wargaming/generic/api"
	"github.com/byvko-dev/am-types/wargaming/v1/accounts"
)

type CompleteProfile struct {
	BasicProfile

	RecruitingOptions RecruitingOptions   `json:"recruiting_options" bson:"recruiting_options"`
	CreatorName       string              `json:"creator_name" bson:"creator_name"`
	UpdatedAt         api.UnixTimestamp   `json:"updated_at" bson:"updated_at"`
	LeaderName        string              `json:"leader_name" bson:"leader_name"`
	MembersIDS        []accounts.PlayerID `json:"members_ids" bson:"members_ids"`
	RecruitingPolicy  string              `json:"recruiting_policy" bson:"recruiting_policy"`
	IsClanDisbanded   bool                `json:"is_clan_disbanded" bson:"is_clan_disbanded"`
	OldName           string              `json:"old_name" bson:"old_name"`
	EmblemSetID       int                 `json:"emblem_set_id" bson:"emblem_set_id"`
	CreatorID         accounts.PlayerID   `json:"creator_id" bson:"creator_id"`
	Motto             string              `json:"motto" bson:"motto"`
	RenamedAt         api.UnixTimestamp   `json:"renamed_at" bson:"renamed_at"`
	OldTag            string              `json:"old_tag" bson:"old_tag"`
	LeaderID          accounts.PlayerID   `json:"leader_id" bson:"leader_id"`
	Description       string              `json:"description" bson:"description"`
}
