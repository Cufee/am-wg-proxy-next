package client

import (
	"os"
	"strconv"
	"strings"

	"github.com/byvko-dev/am-core/helpers/env"
	"github.com/cufee/am-wg-proxy-next/internal/logs"
	_ "github.com/joho/godotenv/autoload"
)

type bucket struct {
	channel          chan int
	host             string
	port             string
	username         string
	password         string
	wgAppId          string
	responseTimes    chan int
	useDirectInstead bool
	isDirect         bool
}

var rpsBuckets = make(map[string]bucket)
var useDirectBucket = false
var responseThreshold = 2000
var disableProxy = false

func init() {
	maxRps := env.MustGetString("PROXY_HOST_MAX_RPS")
	rpsPerHost, err := strconv.Atoi(maxRps)
	if err != nil {
		panic("PROXY_HOST_MAX_RPS is not a number")
	}

	// Setup fast buckets
	proxyHostList := strings.Split(env.MustGetString("PROXY_HOST_LIST"), ",")
	pass := env.MustGetString("PROXY_PASSWORD")
	user := env.MustGetString("PROXY_USERNAME")
	wgAppID := env.MustGetString("PROXY_WG_APP_ID")
	for _, host := range proxyHostList {
		port := strings.Split(host, ":")[1]
		host = strings.Split(host, ":")[0]

		rpsBuckets[host] = bucket{
			channel:       make(chan int, rpsPerHost),
			host:          host,
			port:          port,
			username:      user,
			password:      pass,
			wgAppId:       wgAppID,
			responseTimes: make(chan int, 20),
		}
	}

	// Direct bucket
	directWgAppID := os.Getenv("DIRECT_WG_APP_ID")
	if directWgAppID != "" {
		disableProxy, _ = strconv.ParseBool(env.MustGetString("DISABLE_PROXY"))
		ms, err := strconv.Atoi(env.MustGetString("DIRECT_MIN_RESPONSE_MS"))
		if err == nil {
			responseThreshold = ms
		}
		useDirectBucket = true
		rpsBuckets["direct"] = bucket{
			channel:       make(chan int, 10),
			wgAppId:       directWgAppID,
			isDirect:      true,
			responseTimes: nil,
		}
	} else {
		useDirectBucket = false
	}
}

func pickBucket() *bucket {
	var bkt bucket
	for _, c := range rpsBuckets {
		if len(c.channel) <= len(bkt.channel) && !c.isDirect {
			bkt = c
		}
	}
	if useDirectBucket && bkt.useDirectInstead {
		c := rpsBuckets["direct"]
		return &c
	}
	defer func() {
		go checkBucketResponseTimes(&bkt)
	}()
	return &bkt
}

func checkBucketResponseTimes(bkt *bucket) {
	if len(bkt.responseTimes) > 5 {
		var sum int
		var count int
		for len(bkt.responseTimes) > 0 {
			rt := <-bkt.responseTimes
			sum += rt
			count++
		}
		avg := sum / count
		if avg > responseThreshold {
			logs.Info("Switching to direct connection for bucket %v due to high (%v) avg response", bkt.host, avg)
			bkt.useDirectInstead = true
		} else {
			logs.Info("Switching to proxy connection for bucket %v due to low (%v) avg response", bkt.host, avg)
			bkt.useDirectInstead = false
		}
	}
}
