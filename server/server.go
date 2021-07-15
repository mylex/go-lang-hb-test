package server

import (
	"fmt"
	"log"
	"net/http"
)

var limiter = NewIPRateLimiter(1, 5)

func Init(port int) {
	portPrefix := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(portPrefix, limitMiddleware(Router()))

	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}
}

func limitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limiter := limiter.GetLimiter(r.RemoteAddr)
		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
