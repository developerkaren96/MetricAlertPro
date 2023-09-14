package handlers

import "github.com/developerkaren96/MetricAlertPro/internal/metrics"

type UpdateGaugeReq struct {
	name  string
	value metrics.Gauge
}

type UpdateCounterReq struct {
	name  string
	value metrics.Counter
}
