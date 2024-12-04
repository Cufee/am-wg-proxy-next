package client

import (
	"bytes"
	"context"

	"github.com/cufee/am-wg-proxy-next/v2/client/common"
	"github.com/cufee/am-wg-proxy-next/v2/internal/json"
	"github.com/cufee/am-wg-proxy-next/v2/types"

	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"

	_ "github.com/joho/godotenv/autoload"
)

func (c *Client) Request(ctx context.Context, realm types.Realm, path, method string, payload []byte, target interface{}) (int, error) {
	baseUri, err := baseUriFromRealm(realm)
	if err != nil {
		return 0, err
	}

	bkt, err := c.getBucket(realm.String())
	if err != nil {
		return 0, err
	}

	bkt.waitForTick(c.logger)
	defer bkt.onComplete(c.logger)

	endpoint, err := url.Parse(baseUri + path)
	if err != nil {
		return 0, err
	}

	query := endpoint.Query()
	query.Set("application_id", bkt.wgAppId)
	endpoint.RawQuery = query.Encode()

	c.logger.Debug().Str("realm", realm.String()).Str("endpoint", endpoint.String()).Msg("Sending request")

	headers := make(map[string]string)
	if bkt.proxyUrl != nil {
		headers["Proxy-Authorization"] = bkt.authHeader
	}

	return c.httpRequest(ctx, endpoint, method, bkt.proxyUrl, nil, payload, target, c.options.Timeout)
}

func baseUriFromRealm(realm types.Realm) (string, error) {
	switch realm {
	default:
		return "", common.ErrRealmNotSupported

	case types.RealmEurope:
		return "https://api.wotblitz.eu/wotb/", nil
	case types.RealmNorthAmerica:
		return "https://api.wotblitz.com/wotb/", nil
	case types.RealmAsia:
		return "https://api.wotblitz.asia/wotb/", nil
	}
}

func (c *Client) httpRequest(ctx context.Context, url *url.URL, method string, proxy *url.URL, headers map[string]string, payload []byte, target interface{}, timeout time.Duration) (int, error) {
	event := c.logger.Debug().Str("path", url.Path).Str("method", method)
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
	defer resp.Body.Close()
	event.Int("status code", resp.StatusCode)

	if target != nil {
		err := json.NewDecoder(resp.Body).Decode(target)
		if err != nil {
			event.Err(errors.Wrap(err, "json#Decode failed"))
			return resp.StatusCode, err
		}
	}
	return resp.StatusCode, nil
}
