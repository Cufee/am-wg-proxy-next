package statistics

// StatsFrame - Stats frame struct to unpack json and bson
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

// Substract - Substract StatsFrame b from StatsFrame a
func (a *StatsFrame) Substract(b *StatsFrame) {
	// Set stats fields
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

// Add - Add StatsFrame b to StatsFrame a
func (a *StatsFrame) Add(b *StatsFrame) {
	// Set stats fields
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
