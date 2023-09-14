package main

import (
	"net/http"

	"github.com/developerkaren96/MetricAlertPro/internal/handlers"
	"github.com/developerkaren96/MetricAlertPro/internal/logging"
	"github.com/developerkaren96/MetricAlertPro/internal/middleware"
	"github.com/developerkaren96/MetricAlertPro/internal/server"
)

type Middleware func(http.Handler) http.Handler

func attachMiddleware(h http.Handler) http.Handler {
	common := [...]Middleware{middleware.RequestsLogger}

	for _, middleware := range common {
		h = middleware(h)
	}
	return h
}

func router(server *server.Server) http.Handler {
	r := http.NewServeMux()

	r.Handle(
		"/",
		attachMiddleware(http.NotFoundHandler()),
	)
	r.Handle(
		"/update/",
		attachMiddleware(http.StripPrefix("/update/", handlers.UpdateMetricHandler{Server: server})),
	)

	return r
}

func main() {
	app := server.NewServer()
	logging.Log.Info("Listening on " + app.Config.ListenAddress)

	if err := http.ListenAndServe(app.Config.ListenAddress, router(app)); err != nil {
		logging.Log.Fatal(err)
	}
}
