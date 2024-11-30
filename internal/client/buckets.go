package client

import (
	"errors"
	"net/url"
	"sync/atomic"

	"github.com/rs/zerolog"

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

	limiter        chan int
	activeRequests *atomic.Int32

	proxyUrl   *url.URL
	authHeader string
}

func (b *proxyBucket) waitForTick(logger zerolog.Logger) {
	logger.Debug().Str("realm", b.realm).Msg("Waiting for tick")

	b.activeRequests.Add(1)
	b.limiter <- 1
}

func (b *proxyBucket) onComplete(logger zerolog.Logger) {
	<-b.limiter
	b.activeRequests.Add(-1)

	logger.Debug().Str("realm", b.realm).Msg("Completed request")
}
