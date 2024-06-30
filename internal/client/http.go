package client

import (
	"bytes"
	"context"
	"encoding/json"

	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	_ "github.com/joho/godotenv/autoload"
)

func (c *Client) Request(ctx context.Context, realm, path, method string, payload []byte, target interface{}) (int, error) {
	bkt, err := c.getBucket(realm)
	if err != nil {
		return 0, err
	}

	bkt.waitForTick()
	defer bkt.onComplete()

	baseUri, err := baseUriFromRealm(realm)
	if err != nil {
		return 0, err
	}

	endpoint, err := url.Parse(baseUri + path)
	if err != nil {
		return 0, err
	}

	query := endpoint.Query()
	query.Set("application_id", bkt.wgAppId)
	endpoint.RawQuery = query.Encode()

	log.Debug().Str("realm", realm).Str("endpoint", endpoint.String()).Msg("Sending request")

	headers := make(map[string]string)
	if bkt.proxyUrl != nil {
		headers["Proxy-Authorization"] = bkt.authHeader
	}

	return httpRequest(ctx, endpoint, method, bkt.proxyUrl, nil, payload, target, c.options.Timeout)
}

func baseUriFromRealm(realm string) (string, error) {
	switch strings.ToUpper(realm) {
	case "EU":
		return "https://api.wotblitz.eu/wotb/", nil
	case "NA":
		return "https://api.wotblitz.com/wotb/", nil
	case "AS":
		return "https://api.wotblitz.asia/wotb/", nil
	default:
		return "", errors.New("unknown realm")
	}
}

func httpRequest(ctx context.Context, url *url.URL, method string, proxy *url.URL, headers map[string]string, payload []byte, target interface{}, timeout time.Duration) (int, error) {
	event := log.Debug().Str("path", url.Path).Str("method", method)
	if proxy != nil {
		event.Str("proxy", proxy.Host)
	}
	defer func() {
		event.Msg("wg api request")
	}()

	// Prep request
	req, err := http.NewRequest(strings.ToUpper(method), url.String(), bytes.NewBuffer(payload))
	if err != nil {
		return 0, err
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
		Timeout:   timeout,
		Transport: transport,
	}
	defer client.CloseIdleConnections()
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		event.Err(errors.Wrap(err, "client#Do failed"))
		return 0, err
	}
	if resp != nil {
		event.Int("status code", resp.StatusCode)
	}
	defer resp.Body.Close()

	if target != nil {
		err := json.NewDecoder(resp.Body).Decode(target)
		if err != nil {
			event.Err(errors.Wrap(err, "json#Decode failed"))
			return resp.StatusCode, err
		}
	}
	return resp.StatusCode, nil
}
