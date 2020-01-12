package ratelimiter

import (
	"log"
	"math"
	"net/http"
	"sync"
	"time"
)

// RateLimitMiddleware is a middleware to perform rate limiting.
func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if getRateLimiter().RateLimit(r) {
			// Continue
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusTooManyRequests)
		}
	})
}

type rateLimiter struct {
	mux                    sync.Mutex
	allowRequestsPerSecond int
	requestTimeQueue       []time.Time
}

var instance *rateLimiter

// InitRateLimiter initialize singleton instance of RateLimiter
func InitRateLimiter(allowRequestPerSecond int) {
	instance = &rateLimiter{}
	instance.allowRequestsPerSecond = allowRequestPerSecond
}

func getRateLimiter() *rateLimiter {
	if instance == nil {
		log.Panicln("[FAILED] You must initialize with InitRateLimiter before use.")
	}
	return instance
}

func (rt *rateLimiter) cleanQueue() {
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

// RateLimit returns boolean value whether an incoming request is allowed.
func (rt *rateLimiter) RateLimit(r *http.Request) bool {
	rt.mux.Lock()
	defer rt.mux.Unlock()

	rt.cleanQueue()

	if len(rt.requestTimeQueue)+1 > rt.allowRequestsPerSecond {
		return false
	}

	rt.requestTimeQueue = append(rt.requestTimeQueue, time.Now())

	return true
}
