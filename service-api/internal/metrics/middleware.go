package metrics

import (
	"net/http"
	"time"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/metrics" {
			next.ServeHTTP(w, r)
			return
		}
		startTime := time.Now()
		HTTPRequests.Inc()
		next.ServeHTTP(w, r)
		totalTime := time.Since(startTime).Seconds()
		HTTPRequestDuration.Observe(totalTime)
	})
}
