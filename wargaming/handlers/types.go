package handlers

type WargamingBaseResponse struct {
	Status string                 `json:"status"`
	Error  WargamingResponseError `json:"error"`
	Meta   WargamingResponseMeta  `json:"meta"`
}

type WargamingResponseMeta struct {
	Count int `json:"count"`
}

type WargamingResponseError struct {
	Message string `json:"message"`
	Field   string `json:"field"`
	Value   string `json:"value"`
	Code    int    `json:"code"`
}
