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

	api "github.com/byvko-dev/am-types/api/generic/v1"
	e "github.com/byvko-dev/am-types/errors/v2"
	_ "github.com/joho/godotenv/autoload"
)

type Client struct {
	httpClient *http.Client
	host       string
}

func NewClient(host string, timeout time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: timeout},
		host:       host,
	}
}

type requestOptions struct {
	Query url.Values
}

var defaultRequestOptions = requestOptions{
	Query: url.Values{},
}

var globalClient = &http.Client{Timeout: time.Second * 3}

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

	glossaryOneVehicleEndpointFMT       endpoint = "/glossary/vehicles/%v"
	glossaryManyVehiclesEndpointFMT     endpoint = "/glossary/vehicles"
	glossaryOneAchievementEndpointFMT   endpoint = "/glossary/achievements/%v"
	glossaryManyAchievementsEndpointFMT endpoint = "/glossary/achievements"

	bulkClanInfoEndpoint            endpoint = "/bulk/clans/info"
	bulkAccountInfoEndpoint         endpoint = "/bulk/accounts/info"
	bulkAccountVehiclesEndpoint     endpoint = "/bulk/accounts/vehicles"
	bulkAccountAchievementsEndpoint endpoint = "/bulk/accounts/achievements"
)

func (c *Client) sendRequest(realm string, path endpoint, target interface{}, optsInput ...requestOptions) *e.Error {
	opts := defaultRequestOptions
	if len(optsInput) > 0 {
		opts = optsInput[0]
	}

	// Build URL
	urlData, err := url.Parse(fmt.Sprintf("http://%s/query/%s%s", c.host, strings.ToUpper(realm), path))
	if err != nil {
		return e.Internal(err, "Failed to parse final request URL")
	}
	urlData.RawQuery = opts.Query.Encode()

	// Send request
	resp, err := globalClient.Get(urlData.String())
	if err != nil {
		return e.Internal(err, "Failed to send request")
	}

	// Error checks
	if resp == nil {
		return e.Internal(errors.New("response is nil"), "Failed to send request")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return e.Internal(err, "Failed to read response body")
	}

	// Header and status checks
	if resp.Header.Get("Content-Type") != "application/json" {
		return e.Internal(errors.New("response is not JSON"), string(body))
	}
	if resp.StatusCode != http.StatusOK {
		return e.Internal(errors.New(resp.Status), string(body))
	}

	// Decode response
	var responseDecoded api.ResponseWithError
	err = json.Unmarshal(body, &responseDecoded)
	if err != nil {
		return e.Internal(err, "Failed to unmarshal response")
	}
	if responseDecoded.Error.Context != "" || responseDecoded.Error.Message != "" {
		return e.Input(errors.New(responseDecoded.Error.Context), responseDecoded.Error.Message)
	}

	// Decode response data to target
	// there is probably a cleaner way to unmarshal a generic interface
	responseData, err := json.Marshal(responseDecoded.Data)
	if err != nil {
		return e.Internal(err, "Failed to parse response data")
	}
	err = json.Unmarshal(responseData, target)
	if err != nil {
		return e.Internal(err, "Failed to decode response")
	}
	return nil
}
