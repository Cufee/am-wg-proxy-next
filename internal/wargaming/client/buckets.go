package client

import (
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/byvko-dev/am-core/helpers/env"
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

var limiter chan int
var rpsBuckets []bucket
var lastBucketIndex int32

func init() {
	// Setup fast buckets
	proxyHostList := strings.Split(env.MustGetString("PROXY_HOST_LIST"), ",")
	pass := env.MustGetString("PROXY_PASSWORD")
	user := env.MustGetString("PROXY_USERNAME")
	wgAppID := env.MustGetString("PROXY_WG_APP_ID")
	for _, host := range proxyHostList {
		port := strings.Split(host, ":")[1]
		host = strings.Split(host, ":")[0]
		rpsBuckets = append(rpsBuckets, bucket{
			channel:  make(chan int, 10),
			host:     host,
			port:     port,
			username: user,
			password: pass,
			wgAppId:  wgAppID,
		})
	}

	maxRps := env.MustGetString("PROXY_HOST_MAX_RPS")
	rpsPerHost, err := strconv.Atoi(maxRps)
	if err != nil {
		panic("PROXY_HOST_MAX_RPS is not a number")
	}
	limiter = make(chan int, rpsPerHost*len(rpsBuckets))
}

func pickBucket() bucket {
	nextIndex := lastBucketIndex + 1
	if int(nextIndex) >= len(rpsBuckets) {
		nextIndex = 0
	}
	atomic.StoreInt32(&lastBucketIndex, int32(nextIndex))
	return rpsBuckets[nextIndex]
}
