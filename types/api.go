package types

type ResponseWithError struct {
	Data  interface{}    `json:"data" bson:"data"`
	Error *ResponseError `json:"error" bson:"error"`
}

type ResponseError struct {
	Message string `json:"message" bson:"message"`
	Context string `json:"context,omitempty" bson:"context"`
	Code    string `json:"code,omitempty" bson:"code"`
}
