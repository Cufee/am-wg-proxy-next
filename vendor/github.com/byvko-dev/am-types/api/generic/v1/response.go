package api

type ResponseWithError struct {
	Data  interface{}   `json:"data" firestore:"data" bson:"data"`
	Error ResponseError `json:"error" firestore:"error" bson:"error"`
}

type ResponseError struct {
	Message string `json:"message" firestore:"message" bson:"message"`
	Context string `json:"context" firestore:"context" bson:"context"`
}
