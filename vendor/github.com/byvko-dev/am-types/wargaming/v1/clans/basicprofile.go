package clans

import "github.com/byvko-dev/am-types/wargaming/generic/api"

type ClanID int

type BasicProfile struct {
	MembersCount int               `json:"members_count" bson:"members_count"`
	Name         string            `json:"name" bson:"name"`
	CreatedAt    api.UnixTimestamp `json:"created_at" bson:"created_at"`
	Tag          string            `json:"tag" bson:"tag"`
	ClanID       ClanID            `json:"clan_id" bson:"clan_id"`
}
