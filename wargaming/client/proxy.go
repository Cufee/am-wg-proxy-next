package client

import (
	"fmt"
	"net/url"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

func getProxyBucketAndUrl(realm, bucketName string) (*bucket, *url.URL, string, error) {
	bkt, err := getRpsBucket(bucketName)
	if err != nil {
		return nil, nil, "", err
	}

	username := bkt.username
	if bucketName != BucketGlobal {
		username += proxyCountryFromRealm(realm)
	}

	return bkt, buildProxyURL(bkt.host, bkt.port, username, bkt.password), fmt.Sprintf("%s:%s", username, bkt.password), nil
}

func buildProxyURL(host, port, username, password string) *url.URL {
	proxyUrl := &url.URL{
		Scheme: "http",
		User:   url.UserPassword(username, password),
		Host:   fmt.Sprintf("%s:%s", host, port),
	}
	return proxyUrl
}

func proxyCountryFromRealm(realm string) string {
	switch strings.ToUpper(realm) {
	case "RU":
		return "-country-de"
	case "EU":
		return "-country-de"
	case "NA":
		return "-country-us"
	case "ASIA":
		fallthrough
	case "AS":
		return "-country-sg"
	default:
		return ""
	}
}
