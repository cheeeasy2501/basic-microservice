package main

import (
	"basic-microservice/internal/app"
	"basic-microservice/internal/config"
	"github.com/sirupsen/logrus"
	"log"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Logger
	l := logrus.New()
	l.SetLevel(logrus.DebugLevel) //cfg.Log.Level
	// Run
	app.Run(cfg, l)
}
