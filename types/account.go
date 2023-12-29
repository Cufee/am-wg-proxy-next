package types

type Account struct {
	Nickname string `json:"nickname" bson:"nickname"`
	ID       int    `json:"account_id" bson:"account_id"`
}

type ExtendedAccount struct {
	Account    // Includes ID and Nickname
	Statistics struct {
		All    StatsFrame `json:"all" bson:"all"`
		Rating StatsFrame `json:"rating" bson:"rating"`
	} `json:"statistics" bson:"statistics"`
	CreatedAt      int `json:"created_at" bson:"created_at"`
	UpdatedAt      int `json:"updated_at" bson:"updated_at"`
	LastBattleTime int `json:"last_battle_time" bson:"last_battle_time"`
}
