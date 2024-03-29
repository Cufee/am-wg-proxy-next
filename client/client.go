package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/cufee/am-wg-proxy-next/types"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog/log"
)

type Client struct {
	debug      bool
	httpClient *http.Client
	host       string
}

type ClientOptions struct {
	Debug bool
}

func NewClient(host string, timeout time.Duration, opts ...ClientOptions) *Client {
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
	urlData, err := url.Parse(fmt.Sprintf("%s/query/%s%s", c.host, strings.ToUpper(realm), path))
	if err != nil {
		return errors.New("failed to parse URL")
	}
	urlData.RawQuery = opts.Query.Encode()

	// Send request
	resp, err := c.httpClient.Get(urlData.String())
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) || os.IsTimeout(err) {
			return ErrRequestTimeOut
		}
		return err
	}
	defer resp.Body.Close()

	if c.debug {
		log.Debug().Str("url", urlData.String()).Msgf("Got response with status %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ErrFailedToDecodeResponse
	}

	// Header and status checks
	if resp.Header.Get("Content-Type") != "application/json" {
		if c.debug {
			log.Debug().Str("url", urlData.String()).Msgf("Response is not JSON. Response body: %s", string(body))
		}
		return ErrUnexpectedContentType
	}

	// Decode response
	var responseDecoded struct {
		types.WgResponse
		Data interface{} `json:"data"`
	}
	err = json.Unmarshal(body, &responseDecoded)
	if err != nil {
		return ErrFailedToDecodeResponse
	}
	if responseDecoded.Error.Message != "" {
		if responseDecoded.Error.Message == "SOURCE_NOT_AVAILABLE" {
			return ErrSourceNotAvailable
		}
		return errors.New(strings.ToLower(strings.ReplaceAll(responseDecoded.Error.Message, "_", " ")))
	}
	if resp.StatusCode > 299 {
		return ErrBadResponseCode
	}

	// Decode response data to target
	// there is probably a cleaner way to unmarshal a generic interface
	responseData, err := json.Marshal(responseDecoded.Data)
	if err != nil {
		return ErrFailedToDecodeResponse
	}
	err = json.Unmarshal(responseData, target)
	if err != nil {
		return ErrFailedToDecodeResponse
	}
	return nil
}
