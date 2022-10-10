package client

import (
	"fmt"
	"net/url"

	"github.com/cufee/am-wg-proxy-next/internal/logs"
	_ "github.com/joho/godotenv/autoload"
)

func getProxyBucketAndUrl(realm string) (*bucket, *url.URL, string, error) {
	bkt := pickBucket()
	if bkt.isDirect {
		logs.Info("Using direct connection for realm %v", realm)
		return bkt, nil, "", nil
	}
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
