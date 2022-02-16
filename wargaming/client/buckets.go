package client

import (
	"errors"
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

const (
	BucketGlobal             = "global"
	BucketCacheUpdate        = "cacheUpdate"
	BucketAchievementsUpdate = "achievementsUpdate"
)

var slowRpsBucket bucket
var fastRpsBuckets = make(map[string]bucket)

func init() {
	// Setup fast buckets
	fastProxyHostList := strings.Split(os.Getenv("FAST_PROXY_HOST_LIST"), ",")
	if len(fastProxyHostList) == 0 {
		panic("FAST_PROXY_HOST_LIST not set")
	}
	fastRps := os.Getenv("FAST_PROXY_HOST_MAX_RPS")
	if fastRps == "" {
		panic("FAST_PROXY_HOST_MAX_RPS not set")
	}
	maxFastRps, err := strconv.Atoi(fastRps)
	if err != nil {
		panic("FAST_PROXY_HOST_MAX_RPS is not a number")
	}
	fastPass := os.Getenv("FAST_PROXY_PASSWORD")
	if fastPass == "" {
		panic("FAST_PROXY_PASSWORD not set")
	}
	fastUser := os.Getenv("FAST_PROXY_USERNAME")
	if fastUser == "" {
		panic("FAST_PROXY_USERNAME not set")
	}
	fastPort := os.Getenv("FAST_PROXY_PORT")
	if fastPort == "" {
		panic("FAST_PROXY_PORT not set")
	}
	fastWgAppId := os.Getenv("FAST_PROXY_WG_APP_ID")
	if fastWgAppId == "" {
		panic("FAST_PROXY_WG_APP_ID not set")
	}
	for _, host := range fastProxyHostList {
		fastRpsBuckets[host] = bucket{
			channel:  make(chan int, maxFastRps),
			host:     host,
			port:     fastPort,
			username: fastUser,
			password: fastPass,
			wgAppId:  fastWgAppId,
		}
	}

	// Setup slow bucket
	var slawBucketHost = os.Getenv("SLOW_PROXY_HOST")
	if slawBucketHost == "" {
		panic("SLOW_PROXY_HOST not set")
	}
	slowRps := os.Getenv("SLOW_PROXY_MAX_RPS")
	if slowRps == "" {
		panic("SLOW_PROXY_MAX_RPS not set")
	}
	maxSlowRps, err := strconv.Atoi(slowRps)
	if err != nil {
		panic("SLOW_PROXY_MAX_RPS is not a number")
	}
	slowPass := os.Getenv("SLOW_PROXY_PASSWORD")
	if slowPass == "" {
		panic("SLOW_PROXY_PASSWORD not set")
	}
	slowUser := os.Getenv("SLOW_PROXY_USERNAME")
	if slowUser == "" {
		panic("SLOW_PROXY_USERNAME not set")
	}
	slowPort := os.Getenv("SLOW_PROXY_PORT")
	if slowPort == "" {
		panic("SLOW_PROXY_PORT not set")
	}
	slowWgAppId := os.Getenv("SLOW_PROXY_WG_APP_ID")
	if slowWgAppId == "" {
		panic("SLOW_PROXY_WG_APP_ID not set")
	}
	slowRpsBucket = bucket{
		channel:  make(chan int, maxSlowRps),
		host:     slawBucketHost,
		port:     slowPort,
		username: slowUser,
		password: slowPass,
		wgAppId:  slowWgAppId,
	}
}

func getRpsBucket(name string) (*bucket, error) {
	switch name {
	case BucketGlobal:
		b := pickFastBucket()
		return &b, nil
	case BucketCacheUpdate:
		return &slowRpsBucket, nil
	case BucketAchievementsUpdate:
		return &slowRpsBucket, nil
	default:
		return nil, errors.New("unknown bucket")
	}
}

func pickFastBucket() bucket {
	var bkt bucket
	for _, c := range fastRpsBuckets {
		if len(c.channel) <= len(bkt.channel) {
			bkt = c
		}
	}
	return bkt
}
