package types

/*
Account statistics schema for WG API
*/
type StatsFrame struct {
	Rating               float64 `json:"mm_rating" bson:"mm_rating"`
	Spotted              int     `json:"spotted" bson:"spotted"`
	Hits                 int     `json:"hits" bson:"hits"`
	Frags                int     `json:"frags" bson:"frags"`
	MaxXp                int     `json:"max_xp" bson:"max_xp"`
	Wins                 int     `json:"wins" bson:"wins"`
	Losses               int     `json:"losses" bson:"losses"`
	CapturePoints        int     `json:"capture_points" bson:"capture_points"`
	Battles              int     `json:"battles" bson:"battles"`
	DamageDealt          int     `json:"damage_dealt" bson:"damage_dealt"`
	DamageReceived       int     `json:"damage_received" bson:"damage_received"`
	MaxFrags             int     `json:"max_frags" bson:"max_frags"`
	Shots                int     `json:"shots" bson:"shots"`
	Xp                   int     `json:"xp" bson:"xp"`
	SurvivedBattles      int     `json:"survived_battles" bson:"survived_battles"`
	DroppedCapturePoints int     `json:"dropped_capture_points" bson:"dropped_capture_points"`
}

/*
Subtracts StatsFrame b from StatsFrame a in place
*/
func (a *StatsFrame) Subtract(b *StatsFrame) {
	a.Spotted = a.Spotted - b.Spotted
	a.Hits = a.Hits - b.Hits
	a.Frags = a.Frags - b.Frags
	a.MaxXp = a.MaxXp - b.MaxXp
	a.Wins = a.Wins - b.Wins
	a.Losses = a.Losses - b.Losses
	a.CapturePoints = a.CapturePoints - b.CapturePoints
	a.Battles = a.Battles - b.Battles
	a.DamageDealt = a.DamageDealt - b.DamageDealt
	a.DamageReceived = a.DamageReceived - b.DamageReceived
	a.MaxFrags = a.MaxFrags - b.MaxFrags
	a.Shots = a.Shots - b.Shots
	a.Xp = a.Xp - b.Xp
	a.SurvivedBattles = a.SurvivedBattles - b.SurvivedBattles
	a.DroppedCapturePoints = a.DroppedCapturePoints - b.DroppedCapturePoints
}

/*
Adds StatsFrame b to StatsFrame a in place
*/
func (a *StatsFrame) Add(b *StatsFrame) {
	a.Spotted = a.Spotted + b.Spotted
	a.Hits = a.Hits + b.Hits
	a.Frags = a.Frags + b.Frags
	a.MaxXp = a.MaxXp + b.MaxXp
	a.Wins = a.Wins + b.Wins
	a.Losses = a.Losses + b.Losses
	a.CapturePoints = a.CapturePoints + b.CapturePoints
	a.Battles = a.Battles + b.Battles
	a.DamageDealt = a.DamageDealt + b.DamageDealt
	a.DamageReceived = a.DamageReceived + b.DamageReceived
	a.MaxFrags = a.MaxFrags + b.MaxFrags
	a.Shots = a.Shots + b.Shots
	a.Xp = a.Xp + b.Xp
	a.SurvivedBattles = a.SurvivedBattles + b.SurvivedBattles
	a.DroppedCapturePoints = a.DroppedCapturePoints + b.DroppedCapturePoints
}

/*
Account Achievements schema for WG API
*/
type AchievementsFrame struct {
	MarkOfMastery    int `json:"markOfMastery,omitempty" bson:"markOfMastery,omitempty"`
	MarkOfMasteryI   int `json:"markOfMasteryI,omitempty" bson:"markOfMasteryI,omitempty"`
	MarkOfMasteryII  int `json:"markOfMasteryII,omitempty" bson:"markOfMasteryII,omitempty"`
	MarkOfMasteryIII int `json:"markOfMasteryIII,omitempty" bson:"markOfMasteryIII,omitempty"`
}

/*
Adds b to a in place
*/
func (a *AchievementsFrame) Add(b *AchievementsFrame) {
	// Achievements
	a.MarkOfMastery += b.MarkOfMastery
	a.MarkOfMasteryI += b.MarkOfMasteryI
	a.MarkOfMasteryII += b.MarkOfMasteryII
	a.MarkOfMasteryIII += b.MarkOfMasteryIII
}

/*
Subtracts b from a in place
*/
func (a *AchievementsFrame) Subtract(b *AchievementsFrame) {
	// Achievements
	a.MarkOfMastery -= b.MarkOfMastery
	a.MarkOfMasteryI -= b.MarkOfMasteryI
	a.MarkOfMasteryII -= b.MarkOfMasteryII
	a.MarkOfMasteryIII -= b.MarkOfMasteryIII

}

/*
Account vehicle stats schema for WG API
*/
type VehicleStatsFrame struct {
	Stats          StatsFrame `bson:",inline"`
	LastBattleTime int        `json:"last_battle_time" bson:"last_battle_time"`
	MarkOfMastery  int        `json:"mark_of_mastery"  bson:"mark_of_mastery"`
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

// Subtract vehicleStatsFrame b from vehicleStatsFrame a
func (a *VehicleStatsFrame) Subtract(b *VehicleStatsFrame) {
	a.Stats.Subtract(&b.Stats)
	if a.LastBattleTime > b.LastBattleTime {
		a.LastBattleTime = b.LastBattleTime
	}
	a.MarkOfMastery = b.MarkOfMastery
	a.TankID = b.TankID
}
