package types

type wgError struct {
	Message string `json:"message" bson:"message"`
	Field   string `json:"field" bson:"field"`
	Value   string `json:"value" bson:"value"`
	Code    int    `json:"code" bson:"code"`
}

type wgMeta struct {
	Count int `json:"count" bson:"count"`
}

type WgResponse[D any] struct {
	Status string  `json:"status" bson:"status"`
	Error  wgError `json:"error" bson:"error"`
	Meta   wgMeta  `json:"meta" bson:"meta"`
	Data   D       `json:"data" bson:"data"`
}
