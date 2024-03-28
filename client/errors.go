package client

import "errors"

var (
	ErrUnexpectedContentType  = errors.New("unexpected content type of response received")
	ErrFailedToDecodeResponse = errors.New("failed to decode response")
	ErrBadResponseCode        = errors.New("bad response status code")
	ErrRequestTimeOut         = errors.New("request timed out")
)
