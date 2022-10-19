package client

import (
	"fmt"
	"net/url"

	_ "github.com/joho/godotenv/autoload"
)

func getProxyInfo(realm string) (*url.URL, string, string, error) {
	bkt := pickBucket()
	return buildProxyURL(bkt.host, bkt.port, bkt.username, bkt.password), fmt.Sprintf("%s:%s", bkt.username, bkt.password), bkt.wgAppId, nil
}

func buildProxyURL(host, port, username, password string) *url.URL {
	proxyUrl := &url.URL{
		Scheme: "http",
		User:   url.UserPassword(username, password),
		Host:   fmt.Sprintf("%s:%s", host, port),
	}
	return proxyUrl
}
