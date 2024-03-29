package metric

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	pathCountersMutex sync.Mutex
	pathCounters      = make(map[string]*prometheus.CounterVec)
)

func Count(path string) *prometheus.CounterVec {
	pathCountersMutex.Lock()
	defer pathCountersMutex.Unlock()

	if counter, exists := pathCounters[path]; exists {
		return counter
	}

	requestCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "go_authn_requests_total",
			Help: "Total number of requests to " + path,
		},
		[]string{"path"},
	)

	prometheus.MustRegister(requestCounter)
	pathCounters[path] = requestCounter

	return requestCounter
}
