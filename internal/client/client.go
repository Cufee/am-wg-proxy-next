package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	_ "github.com/joho/godotenv/autoload"
)

func HttpRequest(url, method string, proxy *url.URL, headers map[string]string, payload []byte, target interface{}) (int, error) {
	var err error
	var bodyBytes []byte
	var resp *http.Response
	defer func() {
		// Logging
		if err != nil || resp == nil || resp.StatusCode != http.StatusOK {
			log.Warn().Str("url", url).Str("method", method).Str("proxy", proxy.String()).Str("payload", string(payload)).Str("response", string(bodyBytes)).Err(err).Msg("HttpRequest failed")
		}
	}()

	// Prep request
	req, err := http.NewRequest(strings.ToUpper(method), url, bytes.NewBuffer(payload))
	if err != nil {
		return 0, errors.Wrap(err, "http.NewRequest")
	}

	// Set headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// Set payload headers
	if payload != nil {
		req.Header.Set("content-type", "application/json")
	}

	// Send request
	transport := &http.Transport{}
	if proxy != nil {
		transport.Proxy = http.ProxyURL(proxy)
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: transport,
	}
	defer client.CloseIdleConnections()
	resp, err = client.Do(req)
	if err != nil {
		log.Warn().Str("url", url).Str("method", method).Str("proxy", proxy.String()).Str("payload", string(payload)).Err(err).Msg("client.Do failed")
		return resp.StatusCode, err
	}
	defer resp.Body.Close()

	if target != nil {
		// Read body
		bodyBytes, err = io.ReadAll(resp.Body)
		if err != nil {
			return resp.StatusCode, err
		}

		// Decode
		err = json.Unmarshal(bodyBytes, target)
		if err != nil {
			return resp.StatusCode, err
		}
	}
	return resp.StatusCode, nil
}
