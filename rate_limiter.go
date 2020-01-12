package main

import (
	"math"
	"net/http"
	"sync"
	"time"
)

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if GetRateLimiter().RateLimit(r) {
			// Continue
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusTooManyRequests)
		}
	})
}

type SimpleRateLimiter struct {
	mux                    sync.Mutex
	allowRequestsPerSecond int
	requestTimeQueue       []time.Time
}

var instance *SimpleRateLimiter

func InitRateLimiter(allowRequestPerSecond int) {
	instance = &SimpleRateLimiter{}
	instance.allowRequestsPerSecond = allowRequestPerSecond
}

func GetRateLimiter() *SimpleRateLimiter {
	if instance == nil {
		instance = &SimpleRateLimiter{}
	}
	return instance
}

func (rt *SimpleRateLimiter) cleanQueue() {
	ct := time.Now()

	for len(rt.requestTimeQueue) > 0 {
		v := rt.requestTimeQueue[0]

		d := math.Abs(v.Sub(ct).Seconds())
		if d < 1.0 {
			break
		}

		rt.requestTimeQueue = rt.requestTimeQueue[1:]
	}
}

func (rt *SimpleRateLimiter) RateLimit(r *http.Request) bool {
	rt.mux.Lock()
	defer rt.mux.Unlock()

	rt.cleanQueue()

	if len(rt.requestTimeQueue)+1 > rt.allowRequestsPerSecond {
		return false
	}

	rt.requestTimeQueue = append(rt.requestTimeQueue, time.Now())

	return true
}