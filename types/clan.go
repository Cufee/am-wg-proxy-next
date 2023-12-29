package types

type Clan struct {
	MembersCount int    `json:"members_count" bson:"members_count"`
	Name         string `json:"name" bson:"name"`
	CreatedAt    int    `json:"created_at" bson:"created_at"`
	Tag          string `json:"tag" bson:"tag"`
	ID           int    `json:"clan_id" bson:"clan_id"`
}

type ExtendedClan struct {
	Clan

	RecruitingOptions ClanRecruitingOptions `json:"recruiting_options" bson:"recruiting_options"`
	CreatorName       string                `json:"creator_name" bson:"creator_name"`
	UpdatedAt         int                   `json:"updated_at" bson:"updated_at"`
	LeaderName        string                `json:"leader_name" bson:"leader_name"`
	MembersIDS        []int                 `json:"members_ids" bson:"members_ids"`
	RecruitingPolicy  string                `json:"recruiting_policy" bson:"recruiting_policy"`
	IsClanDisbanded   bool                  `json:"is_clan_disbanded" bson:"is_clan_disbanded"`
	OldName           string                `json:"old_name" bson:"old_name"`
	EmblemSetID       int                   `json:"emblem_set_id" bson:"emblem_set_id"`
	CreatorID         int                   `json:"creator_id" bson:"creator_id"`
	Motto             string                `json:"motto" bson:"motto"`
	RenamedAt         int                   `json:"renamed_at" bson:"renamed_at"`
	OldTag            string                `json:"old_tag" bson:"old_tag"`
	LeaderID          int                   `json:"leader_id" bson:"leader_id"`
	Description       string                `json:"description" bson:"description"`
}

type ClanMember struct {
	Clan        Clan   `json:"clan" bson:"clan"`
	AccountID   int    `json:"account_id" bson:"account_id"`
	JoinedAt    int    `json:"joined_at" bson:"joined_at"`
	ClanID      int    `json:"clan_id" bson:"clan_id"`
	Role        string `json:"role" bson:"role"`
	AccountName string `json:"account_name" bson:"account_name"`
}

type ClanRecruitingOptions struct {
	VehiclesLevel        int `json:"vehicles_level" bson:"vehicles_level"`
	WINSRatio            int `json:"wins_ratio" bson:"wins_ratio"`
	AverageBattlesPerDay int `json:"average_battles_per_day" bson:"average_battles_per_day"`
	Battles              int `json:"battles" bson:"battles"`
	AverageDamage        int `json:"average_damage" bson:"average_damage"`
}
