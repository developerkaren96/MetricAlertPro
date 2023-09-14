package main

import (
	"context"

	"github.com/developerkaren96/MetricAlertPro/internal/agent"
	"github.com/developerkaren96/MetricAlertPro/internal/logging"
	"github.com/developerkaren96/MetricAlertPro/internal/metrics"
)

func main() {
	agent := agent.NewAgent()
	ctx := context.Background()

	stats := metrics.NewMetrics()
	go agent.Poll(ctx, stats)
	go agent.Report(ctx, stats)

	logging.Log.Info("Agent has started")
	select {}
}
