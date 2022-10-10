package statistics

// VehicleStats - Player Vehicle stats struct, used to return final data
type VehicleStatsFrame struct {
	Stats          StatsFrame `json:"all" bson:"all"`
	LastBattleTime int        `json:"last_battle_time" bson:"last_battle_time"`
	MarkOfMastery  int        `json:"mark_of_mastery" bson:"mark_of_mastery"`
	TankID         int        `json:"tank_id" bson:"tank_id"`
}

// Adds vehicleStatsFrame a to vehicleStatsFrame b
func (a *VehicleStatsFrame) Add(b *VehicleStatsFrame) {
	a.Stats.Add(&b.Stats)
	if a.LastBattleTime < b.LastBattleTime {
		a.LastBattleTime = b.LastBattleTime
	}
	a.MarkOfMastery = b.MarkOfMastery
	a.TankID = b.TankID
}

// Substract vehicleStatsFrame b from vehicleStatsFrame a
func (a *VehicleStatsFrame) Substract(b *VehicleStatsFrame) {
	a.Stats.Substract(&b.Stats)
	if a.LastBattleTime > b.LastBattleTime {
		a.LastBattleTime = b.LastBattleTime
	}
	a.MarkOfMastery = b.MarkOfMastery
	a.TankID = b.TankID
}
