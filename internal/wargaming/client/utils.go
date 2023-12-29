package client

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/cufee/am-wg-proxy-next/internal/logs"
)

func parseProxySettings(input string, fallbackWgAppId string, fallbackRps int) (*proxyBucket, error) {
	var bucketSettings proxyBucket

	// some kind of valid protocol is required for url.Parse
	if !strings.Contains(input, "://") {
		input = "proxy://" + input
	}
	// user:password@host:port?wgAppId=your_app_id&maxRps=20?realm=na
	parsed, err := url.Parse(input)
	if err != nil {
		return nil, err
	}

	bucketSettings.port = parsed.Port()
	bucketSettings.host = parsed.Hostname()
	bucketSettings.username = parsed.User.Username()
	bucketSettings.password, _ = parsed.User.Password()

	bucketSettings.realm = strings.ToUpper(parsed.Query().Get("realm"))
	bucketSettings.wgAppId = parsed.Query().Get("wgAppId")
	if bucketSettings.wgAppId == "" {
		bucketSettings.wgAppId = fallbackWgAppId
	}

	if rps, err := strconv.Atoi(parsed.Query().Get("maxRps")); err == nil {
		bucketSettings.rps = rps
	} else {
		bucketSettings.rps = fallbackRps
	}

	bucketSettings.mu = sync.Mutex{}
	bucketSettings.ticker = time.NewTicker(time.Second / time.Duration(bucketSettings.rps))

	bucketSettings.proxyUrl = buildProxyURL(bucketSettings.host, bucketSettings.port, bucketSettings.username, bucketSettings.password)
	bucketSettings.authHeader = "Basic " + base64.StdEncoding.EncodeToString([]byte(bucketSettings.username+":"+bucketSettings.password))

	return &bucketSettings, nil
}

func buildProxyURL(host, port, username, password string) *url.URL {
	proxyUrl := &url.URL{
		Scheme: "http",
		User:   url.UserPassword(username, password),
		Host:   fmt.Sprintf("%s:%s", host, port),
	}
	return proxyUrl
}

func pickBucket(realm string) (*proxyBucket, error) {
	if len(proxyBuckets) == 0 {
		logs.Warning("No proxy buckets configured")
		return nil, nil
	}

	if realm == "" {
		realm = "*"
	}

	buckets, ok := proxyBuckets[realm]
	if !ok {
		logs.Warning("No proxy buckets configured for realm %v, using fallback", realm)
		buckets, ok = proxyBuckets["*"]
		if !ok {
			return nil, errors.New("no proxy buckets configured")
		}
	}

	if len(buckets) == 1 {
		return buckets[0], nil
	}

	// Pick the bucket with the lowest active requests
	var lowestRpsBucket int
	for i := range buckets {
		if buckets[i].activeRequests < buckets[lowestRpsBucket].activeRequests {
			lowestRpsBucket = i
		}
	}

	return buckets[lowestRpsBucket], nil
}
