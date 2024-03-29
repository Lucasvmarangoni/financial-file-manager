package metric

import (
	errors "github.com/Lucasvmarangoni/logella/err"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

type Service struct {
	pushgatewayHistogram *prometheus.HistogramVec
	httpRequestHistogram *prometheus.HistogramVec
	exporterHistogram    *prometheus.HistogramVec
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
	app := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "exporter",
		Name:      "request_duration_seconds",
		Help:      "The latency of the HTTP requests.",
		Buckets:   prometheus.DefBuckets,
	}, []string{"name"})

	s := &Service{
		pushgatewayHistogram: cli,
		httpRequestHistogram: http,
		exporterHistogram:    app,
	}
	err := prometheus.Register(s.pushgatewayHistogram)
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
	err := s.pushgateway(c)
	if err != nil {
		return errors.ErrCtx(err, "s.pushgateway")
	}
	err = s.exporter(c)
	if err != nil {
		return errors.ErrCtx(err, "s.exporter")
	}

	return nil
}

func (s *Service) pushgateway(c *CLI) error {
	gatewayURL := "prometheus-pushgateway:9091"
	s.pushgatewayHistogram.WithLabelValues(c.Name).Observe(c.Duration)
	return push.New(gatewayURL, "cmd_job").Collector(s.pushgatewayHistogram).Push()
}

func (s *Service) exporter(c *CLI) error {
	gatewayURL := "nginx-prometheus-exporter:9113"
	s.exporterHistogram.WithLabelValues(c.Name).Observe(c.Duration)
	return push.New(gatewayURL, "cmd_job").Collector(s.exporterHistogram).Push()
}

func (s *Service) SaveHTTP(h *HTTP) {
	s.httpRequestHistogram.WithLabelValues(h.Handler, h.Method, h.StatusCode).Observe(h.Duration)
}
