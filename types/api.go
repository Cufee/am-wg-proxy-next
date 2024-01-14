package types

type ResponseWithError[T any] struct {
	Data  T             `json:"data" bson:"data"`
	Error ResponseError `json:"error" bson:"error"`
}

type ResponseError struct {
	Message string `json:"message,omitempty" bson:"message"`
	Context string `json:"context,omitempty" bson:"context"`
	Code    string `json:"code,omitempty" bson:"code"`
}
