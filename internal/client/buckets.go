package client

import (
	"context"
	"errors"
	"net/url"
	"sync"

	"github.com/rs/zerolog/log"
	"golang.org/x/sync/semaphore"

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
	limiter        *semaphore.Weighted
	activeRequests int

	proxyUrl   *url.URL
	authHeader string
}

func (b *proxyBucket) waitForTick(ctx context.Context) error {
	log.Debug().Str("realm", b.realm).Msg("Waiting for tick")

	b.mu.Lock()
	b.activeRequests++
	b.mu.Unlock()

	return b.limiter.Acquire(ctx, 1)
}

func (b *proxyBucket) onComplete() {
	b.limiter.Release(1)

	b.mu.Lock()
	b.activeRequests--
	b.mu.Unlock()

	log.Debug().Str("realm", b.realm).Msg("Completed request")
}
