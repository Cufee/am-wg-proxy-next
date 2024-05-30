package client

import (
	"errors"
	"net/url"
	"sync"

	"github.com/rs/zerolog/log"

	_ "github.com/joho/godotenv/autoload"
)

var bucketKeyWildcard = "*"
var ErrNoProxyBuckets = errors.New("no proxy buckets configured")

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
