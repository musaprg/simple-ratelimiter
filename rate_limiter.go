package ratelimiter

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

type rateLimiter struct {
	mux                    sync.Mutex
	allowRequestsPerSecond int
	requestTimeQueue       []time.Time
}

var instance *rateLimiter

func InitRateLimiter(allowRequestPerSecond int) {
	instance = &rateLimiter{}
	instance.allowRequestsPerSecond = allowRequestPerSecond
}

func GetRateLimiter() *rateLimiter {
	if instance == nil {
		instance = &rateLimiter{}
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
