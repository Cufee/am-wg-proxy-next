package types

type WgError struct {
	Message string `json:"message" bson:"message"`
	Field   string `json:"field" bson:"field"`
	Value   string `json:"value" bson:"value"`
	Code    int    `json:"code" bson:"code"`
}

type WgMeta struct {
	Count int `json:"count" bson:"count"`
}

type WgResponse struct {
	Status string  `json:"status" bson:"status"`
	Error  WgError `json:"error" bson:"error"`
	Meta   WgMeta  `json:"meta" bson:"meta"`
}
