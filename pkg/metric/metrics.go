package metric

import (


    "github.com/prometheus/client_golang/prometheus"
)

var pathCounters = make(map[string]*prometheus.CounterVec)

func Count(path string) *prometheus.CounterVec {
    if counter, exists := pathCounters[path]; exists {
        return counter
    }

    requestCounter := prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "authn_create_requests_total",
            Help: "Total number of requests to " + path,
        },
        []string{"method"},
    )

    prometheus.MustRegister(requestCounter)

    pathCounters[path] = requestCounter

    return requestCounter
}
