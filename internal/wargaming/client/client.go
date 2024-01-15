package client

import (
	"errors"
	"net/url"
	"strings"

	"github.com/cufee/am-wg-proxy-next/internal/client"
	"github.com/rs/zerolog/log"

	_ "github.com/joho/godotenv/autoload"
)

func WargamingRequest(realm, path, method string, payload []byte, target interface{}) (int, error) {
	bkt, err := pickBucket(realm)
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

	return client.HttpRequest(endpoint.String(), method, bkt.proxyUrl, nil, payload, target)
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
