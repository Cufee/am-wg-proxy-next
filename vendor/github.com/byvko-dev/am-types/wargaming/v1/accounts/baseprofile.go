package accounts

type BaseProfile struct {
	Nickname  string   `json:"nickname" bson:"nickname"`
	AccountID PlayerID `json:"account_id" bson:"account_id"`
}
