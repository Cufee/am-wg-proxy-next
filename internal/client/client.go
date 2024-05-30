package client

import (
	"errors"
	"sync"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type Options struct {
	BucketsString string
	Timeout       time.Duration
}

type Client struct {
	proxyBuckets map[string][]*proxyBucket
	options      Options
}

func (c *Client) addBucket(key string, bucket *proxyBucket) {
	if c.proxyBuckets == nil {
		c.proxyBuckets = make(map[string][]*proxyBucket)
	}
	c.proxyBuckets[key] = append(c.proxyBuckets[key], bucket)
}

func NewClient(wargamingAppID string, requestsPerSecond int, opts Options) (*Client, error) {
	if wargamingAppID == "" {
		return nil, errors.New("wargaming application id is required")
	}
	if opts.Timeout == 0 {
		opts.Timeout = time.Second * 3
	}

	client := Client{proxyBuckets: make(map[string][]*proxyBucket), options: opts}
	client.addBucket(bucketKeyWildcard, &proxyBucket{
		mu:             sync.Mutex{},
		rps:            requestsPerSecond,
		wgAppId:        wargamingAppID,
		limiter:        make(chan int, requestsPerSecond),
		activeRequests: requestsPerSecond * 100, // Make sure this bucket is never picked
		proxyUrl:       nil,
	})

	for key, bucketSlice := range ParseProxyString(opts.BucketsString, wargamingAppID, requestsPerSecond) {
		for _, b := range bucketSlice {
			client.addBucket(key, b)
		}
	}

	return &client, nil
}

func (c *Client) getBucket(key string) (*proxyBucket, error) {
	if len(c.proxyBuckets) == 0 {
		return nil, nil
	}

	buckets, ok := c.proxyBuckets[key]
	if !ok || len(buckets) == 0 {
		wildcardBuckets, ok := c.proxyBuckets[bucketKeyWildcard]
		if !ok {
			return nil, ErrNoProxyBuckets
		}
		buckets = wildcardBuckets
	}
	if len(buckets) == 1 {
		return buckets[0], nil
	}

	// Pick the bucket with the lowest active requests
	var lowestRpsBucketIndex int
	for i := range buckets {
		if buckets[i].activeRequests < buckets[lowestRpsBucketIndex].activeRequests {
			lowestRpsBucketIndex = i
		}
	}

	return buckets[lowestRpsBucketIndex], nil
}
