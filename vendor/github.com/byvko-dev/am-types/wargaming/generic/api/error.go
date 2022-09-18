package api

type Error struct {
	Message string `json:"message" bson:"message"`
	Field   string `json:"field" bson:"field"`
	Value   string `json:"value" bson:"value"`
	Code    int    `json:"code" bson:"code"`
}
