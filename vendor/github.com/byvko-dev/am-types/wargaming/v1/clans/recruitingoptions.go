package clans

type RecruitingOptions struct {
	VehiclesLevel        int `json:"vehicles_level" bson:"vehicles_level"`
	WINSRatio            int `json:"wins_ratio" bson:"wins_ratio"`
	AverageBattlesPerDay int `json:"average_battles_per_day" bson:"average_battles_per_day"`
	Battles              int `json:"battles" bson:"battles"`
	AverageDamage        int `json:"average_damage" bson:"average_damage"`
}
