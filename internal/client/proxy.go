package client

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"golang.org/x/time/rate"
)

func parseProxySettings(input string, fallbackWgAppId string, fallbackRps int) (*proxyBucket, error) {
	bucketSettings := newProxyBucket(fallbackRps)

	// some kind of valid protocol is required for url.Parse
	if !strings.Contains(input, "://") {
		input = "http://" + input
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
		bucketSettings.limiter = rate.NewLimiter(rate.Limit(rps), 1)
	} else {
		bucketSettings.rps = fallbackRps
	}

	bucketSettings.proxyUrl = buildProxyURL(bucketSettings.host, bucketSettings.port, bucketSettings.username, bucketSettings.password)

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

func ParseProxyString(input string, fallbackWgAppID string, fallbackRps int) map[string][]*proxyBucket {
	if input == "" {
		return nil
	}
	proxyBuckets := make(map[string][]*proxyBucket)

	for _, proxyString := range strings.Split(input, ",") {
		bucketSettings, err := parseProxySettings(proxyString, fallbackWgAppID, fallbackRps)
		if err != nil {
			panic(err)
		}
		if bucketSettings.realm == "" {
			proxyBuckets["*"] = append(proxyBuckets["*"], bucketSettings)
		} else {
			proxyBuckets[bucketSettings.realm] = append(proxyBuckets[bucketSettings.realm], bucketSettings)
		}
	}

	return proxyBuckets
}
