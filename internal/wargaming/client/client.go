package client

import (
	"encoding/base64"
	"errors"
	"net/url"
	"strings"
	"time"

	"github.com/cufee/am-wg-proxy-next/internal/client"
	"github.com/cufee/am-wg-proxy-next/internal/logs"
	_ "github.com/joho/godotenv/autoload"
)

// Application ID will be added to query string here
func WargamingRequest(realm, path, method string, payload []byte, target interface{}) (int, error) {
	bucket, proxyUrl, proxyAuth, err := getProxyBucketAndUrl(realm)
	if err != nil {
		return 0, err
	}

	start := time.Now()
	bucket.channel <- 1
	defer func() {
		go func() {
			if time.Since(start) < time.Second {
				time.Sleep(time.Second - time.Since(start))
			}
			<-bucket.channel
		}()
	}()

	baseUri, err := baseUriFromRealm(realm)
	if err != nil {
		return 0, err
	}

	endpoint, err := url.Parse(baseUri + path)
	if err != nil {
		return 0, err
	}

	query := endpoint.Query()
	query.Set("application_id", bucket.wgAppId)
	endpoint.RawQuery = query.Encode()

	logs.Debug("WargamingRequest: %v %v", method, endpoint.String())

	headers := make(map[string]string)
	if proxyUrl != nil {
		basic := "Basic " + base64.StdEncoding.EncodeToString([]byte(proxyAuth))
		headers["Proxy-Authorization"] = basic
	}

	startTime := time.Now()
	defer func() {
		if bucket.responseTimes != nil {
			go func() {
				bucket.responseTimes <- int(time.Since(startTime) / time.Millisecond)
			}()
		}
	}()
	return client.HttpRequest(endpoint.String(), method, proxyUrl, nil, payload, target)
}

func baseUriFromRealm(realm string) (string, error) {
	switch strings.ToUpper(realm) {
	case "RU":
		return "https://api.wotblitz.ru/wotb/", nil
	case "EU":
		return "https://api.wotblitz.eu/wotb/", nil
	case "NA":
		return "https://api.wotblitz.com/wotb/", nil
	case "ASIA":
		fallthrough
	case "AS":
		return "https://api.wotblitz.asia/wotb/", nil
	default:
		return "", errors.New("unknown realm")
	}
}
