package main

import (
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/logging"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/router"
)

func main() {
	logging.InitLogger()

	logging.Logger.Info("SP// Backend Developer Test - RESTful Service")

	r := router.NewRouter()

	// Add prometheus middleware for metrics
	p := prometheus.NewPrometheus("stackpathrest", nil)
	p.Use(r)

	r.Logger.Fatal(r.Start(":8080"))
}
