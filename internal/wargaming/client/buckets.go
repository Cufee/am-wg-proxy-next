package client

import (
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/cufee/am-wg-proxy-next/internal/utils"
	"github.com/rs/zerolog/log"

	_ "github.com/joho/godotenv/autoload"
)

type proxyBucket struct {
	rps      int
	host     string
	port     string
	username string
	password string

	realm   string
	wgAppId string

	mu             sync.Mutex
	limiter        chan int
	activeRequests int

	proxyUrl   *url.URL
	authHeader string
}

func (b *proxyBucket) waitForTick() {
	log.Debug().Str("realm", b.realm).Msg("Waiting for tick")

	b.mu.Lock()
	b.activeRequests++
	b.mu.Unlock()
	b.limiter <- 1
}

func (b *proxyBucket) onComplete() {
	<-b.limiter

	b.mu.Lock()
	b.activeRequests--
	b.mu.Unlock()

	log.Debug().Str("realm", b.realm).Msg("Completed request")
}

var proxyBuckets map[string][]*proxyBucket = make(map[string][]*proxyBucket)

func init() {
	// Get fallback settings
	fallbackWgAppId := utils.MustGetEnv("FALLBACK_WG_APP_ID")
	fallbackRps, err := strconv.Atoi(utils.MustGetEnv("FALLBACK_MAX_RPS"))
	if err != nil || fallbackRps == 0 {
		panic("FALLBACK_MAX_RPS is not a number")
	}

	fallbackLimiter := make(chan int, fallbackRps)
	fallbackBucket := &proxyBucket{
		mu:             sync.Mutex{},
		rps:            fallbackRps,
		wgAppId:        fallbackWgAppId,
		limiter:        fallbackLimiter,
		activeRequests: fallbackRps * 100, // Make sure this bucket is never picked
		proxyUrl:       nil,
	}
	proxyBuckets["*"] = append(proxyBuckets["*"], fallbackBucket)

	proxyHostList := os.Getenv("PROXY_HOST_LIST")
	if proxyHostList == "" {
		log.Warn().Msg("PROXY_HOST_LIST is empty, using fallback bucket")
		return
	}

	// Parse proxy settings
	for _, proxyString := range strings.Split(proxyHostList, ",") {
		bucketSettings, err := parseProxySettings(proxyString, fallbackWgAppId, fallbackRps)
		if err != nil {
			panic(err)
		}
		if bucketSettings.realm == "" {
			proxyBuckets["*"] = append(proxyBuckets["*"], bucketSettings)
		}
		proxyBuckets[bucketSettings.realm] = append(proxyBuckets[bucketSettings.realm], bucketSettings)
	}
}
