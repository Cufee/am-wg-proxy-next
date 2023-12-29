package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/cufee/am-wg-proxy-next/internal/logs"
	"github.com/cufee/am-wg-proxy-next/types"
	_ "github.com/joho/godotenv/autoload"
)

type Client struct {
	debug      bool
	httpClient *http.Client
	host       string
}

type ClientOptons struct {
	Debug bool
}

func NewClient(host string, timeout time.Duration, opts ...ClientOptons) *Client {
	var debug bool
	if len(opts) > 0 {
		debug = opts[0].Debug
	}

	return &Client{
		httpClient: &http.Client{Timeout: timeout},
		host:       host,
		debug:      debug,
	}
}

func (c *Client) Close() {
	c.httpClient.CloseIdleConnections()
}

type requestOptions struct {
	Query url.Values
}

func newDefaultRequestOptions() requestOptions {
	defaultRequestOptions := requestOptions{
		Query: url.Values{},
	}
	return defaultRequestOptions
}

type endpoint string

func (e endpoint) Fmt(args ...interface{}) endpoint {
	return endpoint(fmt.Sprintf(string(e), args...))
}

const (
	accountsSearchEndpoint            endpoint = "/accounts/search"
	accountsGetEndpointFMT            endpoint = "/accounts/%v"
	accountClanGetEndpointFMT         endpoint = "/accounts/%v/clan"
	accountGetVehiclesEndpointFMT     endpoint = "/accounts/%v/vehicles"
	accountGetAchievementsEndpointFMT endpoint = "/accounts/%v/achievements"

	clansSearchEndpoint endpoint = "/clans/search"
	clansGetEndpointFMT endpoint = "/clans/%v"

	glossaryOneVehicleEndpointFMT     endpoint = "/glossary/vehicles/%v"
	glossaryManyVehiclesEndpoint      endpoint = "/glossary/vehicles"
	glossaryOneAchievementEndpointFMT endpoint = "/glossary/achievements/%v"
	glossaryManyAchievementsEndpoint  endpoint = "/glossary/achievements"

	bulkClanInfoEndpoint            endpoint = "/bulk/clans/info"
	bulkAccountInfoEndpoint         endpoint = "/bulk/accounts/info"
	bulkAccountClanInfoEndpoint     endpoint = "/bulk/accounts/clan"
	bulkAccountAchievementsEndpoint endpoint = "/bulk/accounts/achievements"
)

func (c *Client) sendRequest(realm string, path endpoint, target interface{}, optsInput ...requestOptions) error {
	opts := newDefaultRequestOptions()
	if len(optsInput) > 0 {
		opts = optsInput[0]
	}

	// Build URL
	urlData, err := url.Parse(fmt.Sprintf("http://%s/query/%s%s", c.host, strings.ToUpper(realm), path))
	if err != nil {
		return errors.New("failed to parse URL")
	}
	urlData.RawQuery = opts.Query.Encode()

	if c.debug {
		logs.Debug("Sending request to %s", urlData.String())
	}

	// Send request
	resp, err := c.httpClient.Get(urlData.String())
	// Error checks
	if resp == nil {
		return errors.New("client.Do returned nil response")
	}
	defer resp.Body.Close()
	if err != nil {
		return errors.Join(err, errors.New("client.Do failed"))
	}

	if c.debug {
		logs.Debug("Got response with status %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Join(err, errors.New("ioutil.ReadAll failed"))
	}

	// Header and status checks
	if resp.Header.Get("Content-Type") != "application/json" {
		if c.debug {
			logs.Debug("Response is not JSON. Response body: %s", string(body))
		}
		return errors.New("response is not JSON")
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("response status is not 200")
	}

	// Decode response
	var responseDecoded struct {
		types.WgResponse
		Data interface{} `json:"data"`
	}
	err = json.Unmarshal(body, &responseDecoded)
	if err != nil {
		return errors.Join(err, errors.New("failed to decode response"))
	}
	if responseDecoded.Error.Message != "" {
		return errors.New(responseDecoded.Error.Message)
	}

	// Decode response data to target
	// there is probably a cleaner way to unmarshal a generic interface
	responseData, err := json.Marshal(responseDecoded.Data)
	if err != nil {
		return errors.Join(err, errors.New("failed to marshal response data"))
	}
	err = json.Unmarshal(responseData, target)
	if err != nil {
		return errors.Join(err, errors.New("failed to unmarshal response data"))
	}
	return nil
}
