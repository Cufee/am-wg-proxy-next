package accounts

import (
	"github.com/byvko-dev/am-types/wargaming/generic/api"
	"github.com/byvko-dev/am-types/wargaming/v1/statistics"
)

type PlayerID int

type CompleteProfile struct {
	BaseProfile // Includes ID and Nickname
	Statistics  struct {
		All    statistics.StatsFrame `json:"all" bson:"all"`
		Rating statistics.StatsFrame `json:"rating" bson:"rating"`
	} `json:"statistics" bson:"statistics"`
	CreatedAt      api.UnixTimestamp `json:"created_at" bson:"created_at"`
	UpdatedAt      api.UnixTimestamp `json:"updated_at" bson:"updated_at"`
	LastBattleTime api.UnixTimestamp `json:"last_battle_time" bson:"last_battle_time"`
}
