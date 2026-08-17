package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var HTTPRequests = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	},
)

var HTTPRequestDuration = prometheus.NewHistogram(
	prometheus.HistogramOpts{
		Name: "http_request_duration_seconds",
		Help: "HTTP request duration in seconds",
	},
)

func Init() {
	prometheus.MustRegister(HTTPRequestDuration, HTTPRequests)
}
