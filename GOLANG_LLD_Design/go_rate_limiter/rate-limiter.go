package main

import (
	"fmt"
	"time"
)


// A rate limiter controls how frequently a function can be executed. Below is a simple token bucket rate limiter:

type RateLimiter struct {
	rate      int
	bucket    chan struct{}
	ticker    *time.Ticker
	stop      chan struct{}
}

// NewRateLimiter creates a new RateLimiter with the specified rate and bucket size
func NewRateLimiter(rate int, bucketSize int) *RateLimiter {
	rl := &RateLimiter{
		rate:   rate,
		bucket: make(chan struct{}, bucketSize),
		ticker: time.NewTicker(time.Second / time.Duration(rate)),
		stop:   make(chan struct{}),
	}

	go rl.start()
	return rl
}

// start refills the bucket at the specified rate
func (rl *RateLimiter) start() {
	for {
		select {
		case <-rl.ticker.C:
			select {
			case rl.bucket <- struct{}{}:
			default:
				// Bucket is full, do nothing
			}
		case <-rl.stop:
			return
		}
	}
}

// Allow checks if a token is available in the bucket
func (rl *RateLimiter) Allow() bool {
	select {
	case <-rl.bucket:
		return true
	default:
		return false
	}
}

// Stop stops the rate limiter
func (rl *RateLimiter) Stop() {
	close(rl.stop)
	rl.ticker.Stop()
}

func main() {
	rl := NewRateLimiter(5, 10) // 5 requests per second, bucket size 10

	for i := 0; i < 20; i++ {
		if rl.Allow() {
			fmt.Println("Request", i+1, "allowed")
		} else {
			fmt.Println("Request", i+1, "rate limited")
		}
		time.Sleep(100 * time.Millisecond)
	}

	rl.Stop()
}