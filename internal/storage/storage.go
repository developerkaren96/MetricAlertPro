package storage

import "github.com/developerkaren96/MetricAlertPro/internal/metrics"

type Storage interface {
	PushCounter(name string, value metrics.Counter)
	PushGauge(name string, value metrics.Gauge)
}
