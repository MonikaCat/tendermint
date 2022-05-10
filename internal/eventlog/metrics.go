package eventlog

import "github.com/go-kit/kit/metrics"

const MetricsSubsystem = "eventlog"

//go:generate go run ../../scripts/metricsgen -struct=Metrics

// Metrics define the metrics exported by the eventlog package.
type Metrics struct {
	numItems metrics.Gauge
}
