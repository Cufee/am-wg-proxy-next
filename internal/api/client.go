package api

import (
	"bytes"
	"context"

	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/cufee/am-wg-proxy-next/v2/client/common"
	"github.com/cufee/am-wg-proxy-next/v2/internal/json"
	"github.com/cufee/am-wg-proxy-next/v2/types"
	_ "github.com/joho/godotenv/autoload"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type Client struct {
	httpClient *http.Client
	host       string
	logger     zerolog.Logger
}

func NewClient(logger zerolog.Logger, host string, timeout time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: timeout},
		host:       host,
		logger:     logger,
	}
}

func (c *Client) Close() {
	c.httpClient.CloseIdleConnections()
}

type requestOptions struct {
	Query   url.Values
	Payload []types.Option
}

func newDefaultRequestOptions(opts []types.Option) requestOptions {
	defaultRequestOptions := requestOptions{
		Query:   url.Values{},
		Payload: opts,
	}
	return defaultRequestOptions
}

type endpoint string

func (e endpoint) Fmt(args ...interface{}) endpoint {
	return endpoint(fmt.Sprintf(string(e), args...))
}

const (
	accountsSearchEndpoint                   endpoint = "/accounts/search"
	accountsGetEndpointFMT                   endpoint = "/accounts/%v"
	accountClanGetEndpointFMT                endpoint = "/accounts/%v/clan"
	accountGetVehiclesEndpointFMT            endpoint = "/accounts/%v/vehicles"
	accountGetAchievementsEndpointFMT        endpoint = "/accounts/%v/achievements"
	accountGetVehicleAchievementsEndpointFMT endpoint = "/accounts/%v/vehicles/achievements"

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

func (c *Client) sendRequest(ctx context.Context, realm types.Realm, path endpoint, target interface{}, opts requestOptions) error {
	// Build URL
	urlData, err := url.Parse(fmt.Sprintf("%s/query/%s%s", c.host, realm.String(), path))
	if err != nil {
		return errors.New("failed to parse URL")
	}
	urlData.RawQuery = opts.Query.Encode()

	payload, err := json.Marshal(opts.Payload)
	if err != nil {
		return errors.Wrap(err, "failed to encode payload")
	}

	request, err := http.NewRequest("POST", urlData.String(), bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	request.Header.Set("content-type", "application/json")

	// Send request
	resp, err := c.httpClient.Do(request.WithContext(ctx))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) || os.IsTimeout(err) {
			return common.ErrRequestTimeOut
		}
		return err
	}
	defer resp.Body.Close()

	c.logger.Debug().Str("url", urlData.String()).Msgf("Got response with status %v", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return common.ErrFailedToDecodeResponse
	}

	// Header and status checks
	if resp.Header.Get("Content-Type") != "application/json" {
		c.logger.Debug().Str("url", urlData.String()).Msgf("Response is not JSON. Response body: %s", string(body))
		return common.ErrUnexpectedContentType
	}

	// Decode response
	var responseDecoded types.WgResponse[any]
	err = json.Unmarshal(body, &responseDecoded)
	if err != nil {
		return common.ErrFailedToDecodeResponse
	}
	if responseDecoded.Error.Message != "" {
		if responseDecoded.Error.Message == "SOURCE_NOT_AVAILABLE" {
			return common.ErrSourceNotAvailable
		}
		return errors.New(strings.ToLower(strings.ReplaceAll(responseDecoded.Error.Message, "_", " ")))
	}
	if resp.StatusCode > 299 {
		return common.ErrBadResponseCode
	}

	// Decode response data to target
	// there is probably a cleaner way to unmarshal a generic interface
	responseData, err := json.Marshal(responseDecoded.Data)
	if err != nil {
		return common.ErrFailedToDecodeResponse
	}
	err = json.Unmarshal(responseData, target)
	if err != nil {
		return common.ErrFailedToDecodeResponse
	}
	return nil
}
