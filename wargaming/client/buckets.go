package client

import (
	"os"
	"strconv"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

type bucket struct {
	channel  chan int
	host     string
	port     string
	username string
	password string
	wgAppId  string
}

var rpsBuckets = make(map[string]bucket)

func init() {
	// Setup fast buckets
	proxyHostList := strings.Split(os.Getenv("PROXY_HOST_LIST"), ",")
	if len(proxyHostList) == 0 {
		panic("PROXY_HOST_LIST not set")
	}
	if !strings.Contains(os.Getenv("PROXY_HOST_LIST"), ":") {
		panic("PROXY_PORT not set and FAST_PROXY_HOST_LIST does not contain a port")
	}

	maxRps := os.Getenv("PROXY_HOST_MAX_RPS")
	if maxRps == "" {
		panic("PROXY_HOST_MAX_RPS not set")
	}
	rpsPerHost, err := strconv.Atoi(maxRps)
	if err != nil {
		panic("PROXY_HOST_MAX_RPS is not a number")
	}

	pass := os.Getenv("PROXY_PASSWORD")
	if pass == "" {
		panic("PROXY_PASSWORD not set")
	}
	user := os.Getenv("PROXY_USERNAME")
	if user == "" {
		panic("PROXY_USERNAME not set")
	}

	wgAppID := os.Getenv("PROXY_WG_APP_ID")
	if wgAppID == "" {
		panic("PROXY_WG_APP_ID not set")
	}
	for _, host := range proxyHostList {
		port := strings.Split(host, ":")[1]
		host = strings.Split(host, ":")[0]

		rpsBuckets[host] = bucket{
			channel:  make(chan int, rpsPerHost),
			host:     host,
			port:     port,
			username: user,
			password: pass,
			wgAppId:  wgAppID,
		}
	}
}

func pickBucket() *bucket {
	var bkt bucket
	for _, c := range rpsBuckets {
		if len(c.channel) <= len(bkt.channel) {
			bkt = c
		}
	}
	return &bkt
}
