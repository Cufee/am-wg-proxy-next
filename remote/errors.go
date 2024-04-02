package remote

import "errors"

var (
	ErrUnexpectedContentType  = errors.New("unexpected content type of response received")
	ErrFailedToDecodeResponse = errors.New("failed to decode response")
	ErrSourceNotAvailable     = errors.New("source not available")
	ErrBadResponseCode        = errors.New("bad response status code")
	ErrRequestTimeOut         = errors.New("request timed out")
)
