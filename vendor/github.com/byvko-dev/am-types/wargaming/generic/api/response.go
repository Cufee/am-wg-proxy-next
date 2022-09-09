package api

type Response struct {
	Status string `json:"status" bson:"status"`
	Error  Error  `json:"error" bson:"error"`
	Meta   Meta   `json:"meta" bson:"meta"`
}
