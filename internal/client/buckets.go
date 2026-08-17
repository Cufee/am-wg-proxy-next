package client

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"sync/atomic"
	"time"

	"github.com/rs/zerolog"
	"golang.org/x/time/rate"

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

	limiter        *rate.Limiter
	activeRequests *atomic.Int32

	proxyUrl   *url.URL
	httpClient *http.Client
}

func newProxyBucket(rps int) proxyBucket {
	var activeRequests atomic.Int32

	return proxyBucket{
		rps:            rps,
		limiter:        rate.NewLimiter(rate.Limit(rps), 1),
		activeRequests: &activeRequests,
	}
}

func (b *proxyBucket) configureHTTPClient(timeout time.Duration) {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxConnsPerHost = b.rps
	transport.MaxIdleConnsPerHost = b.rps
	if b.proxyUrl != nil {
		transport.Proxy = http.ProxyURL(b.proxyUrl)
	} else {
		// Preserve the old direct-request behavior when no proxy is configured.
		transport.Proxy = nil
	}

	b.httpClient = &http.Client{Timeout: timeout, Transport: transport}
}

func (b *proxyBucket) waitForTick(ctx context.Context, logger zerolog.Logger) error {
	logger.Debug().Str("realm", b.realm).Msg("Waiting for tick")

	if err := b.limiter.Wait(ctx); err != nil {
		return err
	}
	b.activeRequests.Add(1)

	return nil
}

func (b *proxyBucket) onComplete(logger zerolog.Logger) {
	b.activeRequests.Add(-1)

	logger.Debug().Str("realm", b.realm).Msg("Completed request")
}
