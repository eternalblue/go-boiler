package main

import (
	"example.com/go/boiler/internal/app"
	"example.com/go/boiler/internal/app/infra/config"
	"example.com/go/boiler/internal/app/infra/rest"

	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.InitConfiguration()

	if logLevel, err := logrus.ParseLevel(cfg.LogLevel); err == nil {
		logrus.SetLevel(logLevel)
	} else {
		logrus.Warn("failed to set logging level")
	}

	application := app.NewApplication(cfg)
	defaultRouter := rest.DefaultRouter()

	rest.MapRoutes(defaultRouter, application)

	if err := defaultRouter.Run(":8080"); err != nil {
		panic(err)
	}
}
