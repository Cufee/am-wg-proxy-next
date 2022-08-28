package client

import (
	"fmt"
	"net/url"

	_ "github.com/joho/godotenv/autoload"
)

func getProxyBucketAndUrl(realm string) (*bucket, *url.URL, string, error) {
	bkt := pickBucket()
	return bkt, buildProxyURL(bkt.host, bkt.port, bkt.username, bkt.password), fmt.Sprintf("%s:%s", bkt.username, bkt.password), nil
}

func buildProxyURL(host, port, username, password string) *url.URL {
	proxyUrl := &url.URL{
		Scheme: "http",
		User:   url.UserPassword(username, password),
		Host:   fmt.Sprintf("%s:%s", host, port),
	}
	return proxyUrl
}
