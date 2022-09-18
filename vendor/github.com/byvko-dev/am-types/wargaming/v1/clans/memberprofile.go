package clans

import (
	"github.com/byvko-dev/am-types/wargaming/generic/api"
	"github.com/byvko-dev/am-types/wargaming/v1/accounts"
)

type MemberProfile struct {
	Clan        BasicProfile      `json:"clan" bson:"clan"`
	AccountID   accounts.PlayerID `json:"account_id" bson:"account_id"`
	JoinedAt    api.UnixTimestamp `json:"joined_at" bson:"joined_at"`
	ClanID      ClanID            `json:"clan_id" bson:"clan_id"`
	Role        string            `json:"role" bson:"role"`
	AccountName string            `json:"account_name" bson:"account_name"`
}
