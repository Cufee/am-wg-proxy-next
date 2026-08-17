package client

import (
	"context"
	"errors"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/rs/zerolog"
)

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return f(request)
}

func newTestProxyBucket(t *testing.T, proxyAddress string, rps int) *proxyBucket {
	t.Helper()

	proxyURL, err := url.Parse(proxyAddress)
	if err != nil {
		t.Fatal(err)
	}
	proxyURL.User = url.UserPassword("user", "password")

	bucket := newProxyBucket(rps)
	bucket.proxyUrl = proxyURL
	bucket.configureHTTPClient(time.Second)

	return &bucket
}

func TestProxiedRequestsReuseConnection(t *testing.T) {
	var mu sync.Mutex
	connections := map[string]struct{}{}
	var proxyAuthorization string

	proxy := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		mu.Lock()
		connections[request.RemoteAddr] = struct{}{}
		proxyAuthorization = request.Header.Get("Proxy-Authorization")
		mu.Unlock()
		w.WriteHeader(http.StatusNoContent)
	}))
	defer proxy.Close()

	bucket := newTestProxyBucket(t, proxy.URL, 10)
	client := Client{logger: zerolog.Nop()}
	endpoint, err := url.Parse("http://upstream.example/resource")
	if err != nil {
		t.Fatal(err)
	}

	for range 2 {
		status, err := client.httpRequest(context.Background(), endpoint, http.MethodGet, bucket, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		if status != http.StatusNoContent {
			t.Fatalf("unexpected status: %d", status)
		}
	}

	mu.Lock()
	defer mu.Unlock()
	if len(connections) != 1 {
		t.Fatalf("expected one proxy connection for two requests, got %d", len(connections))
	}
	if proxyAuthorization != "Basic dXNlcjpwYXNzd29yZA==" {
		t.Fatalf("unexpected proxy authorization: %q", proxyAuthorization)
	}
}

func TestProxiedRequestKeepsIdleConnection(t *testing.T) {
	idle := make(chan struct{})
	closed := make(chan struct{})
	var idleOnce sync.Once
	var closedOnce sync.Once

	proxy := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	proxy.Config.ConnState = func(_ net.Conn, state http.ConnState) {
		switch state {
		case http.StateIdle:
			idleOnce.Do(func() { close(idle) })
		case http.StateClosed:
			closedOnce.Do(func() { close(closed) })
		}
	}
	proxy.Start()
	defer proxy.Close()

	bucket := newTestProxyBucket(t, proxy.URL, 10)
	client := Client{logger: zerolog.Nop()}
	endpoint, err := url.Parse("http://upstream.example/resource")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := client.httpRequest(context.Background(), endpoint, http.MethodGet, bucket, nil, nil); err != nil {
		t.Fatal(err)
	}

	select {
	case <-idle:
	case <-time.After(time.Second):
		t.Fatal("proxy connection did not become idle")
	}

	select {
	case <-closed:
		t.Fatal("proxy connection was closed after the request")
	case <-time.After(100 * time.Millisecond):
	}
}

func TestProxyBucketLimitsRequestsPerSecond(t *testing.T) {
	bucket := newProxyBucket(10)
	started := time.Now()

	for range 11 {
		if err := bucket.waitForTick(context.Background(), zerolog.Nop()); err != nil {
			t.Fatal(err)
		}
		bucket.onComplete(zerolog.Nop())
	}

	if elapsed := time.Since(started); elapsed < 900*time.Millisecond {
		t.Fatalf("11 requests completed in %s; expected a 10 RPS limit", elapsed)
	}
}

func TestRequestReturnsContextCancellationBeforeOutboundCall(t *testing.T) {
	bucket := newProxyBucket(1)
	bucket.wgAppId = "app-id"
	bucket.limiter.Allow() // Consume the one-token burst so the next request waits.

	var requests atomic.Int32
	bucket.httpClient = &http.Client{Transport: roundTripperFunc(func(*http.Request) (*http.Response, error) {
		requests.Add(1)
		return nil, errors.New("unexpected outbound request")
	})}

	client := Client{
		proxyBuckets: map[string][]*proxyBucket{types.RealmEurope.String(): {&bucket}},
		logger:       zerolog.Nop(),
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := client.Request(ctx, types.RealmEurope, "account/list/", http.MethodGet, nil, nil)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context cancellation, got %v", err)
	}
	if requests.Load() != 0 {
		t.Fatalf("expected no outbound requests, got %d", requests.Load())
	}
	if bucket.activeRequests.Load() != 0 {
		t.Fatalf("expected no active requests, got %d", bucket.activeRequests.Load())
	}
}
