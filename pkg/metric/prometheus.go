package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

type Service struct {
	pHistogram           *prometheus.HistogramVec
	httpRequestHistogram *prometheus.HistogramVec
}

func NewPrometheusService() (*Service, error) {
	cli := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "pushgateway",
		Name:      "cmd_duration_seconds",
		Help:      "CLI application execution in seconds",
		Buckets:   prometheus.DefBuckets,
	}, []string{"name"})
	http := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "http",
		Name:      "request_duration_seconds",
		Help:      "The latency of the HTTP requests.",
		Buckets:   prometheus.DefBuckets,
	}, []string{"handler", "method", "code"})

	s := &Service{
		pHistogram:           cli,
		httpRequestHistogram: http,
	}
	err := prometheus.Register(s.pHistogram)
	if err != nil && err.Error() != "duplicate metrics collector registration attempted" {
		return nil, err
	}
	err = prometheus.Register(s.httpRequestHistogram)
	if err != nil && err.Error() != "duplicate metrics collector registration attempted" {
		return nil, err
	}
	return s, nil
}

func (s *Service) SaveCLI(c *CLI) error {
	gatewayURL := "http://prometheus-pushgateway:9091"
	s.pHistogram.WithLabelValues(c.Name).Observe(c.Duration)
	return push.New(gatewayURL, "cmd_job").Collector(s.pHistogram).Push()
}

func (s *Service) SaveHTTP(h *HTTP) {
	s.httpRequestHistogram.WithLabelValues(h.Handler, h.Method, h.StatusCode).Observe(h.Duration)
}
