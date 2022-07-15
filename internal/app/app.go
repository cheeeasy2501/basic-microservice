package app

import (
	"basic-microservice/internal/config"
	"basic-microservice/internal/repository"
	"basic-microservice/internal/service"
	v1 "basic-microservice/internal/transport/http/v1"
	"basic-microservice/pkg/database"
	"basic-microservice/pkg/httpserver"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config, l *logrus.Logger) {
	l.Info("app - Run - Run app")

	// Configure gorm DB
	db, err := database.NewDatabase(&cfg.DB)
	if err != nil {
		l.Errorf("app - Run - Bad database instance: %w", err)
		return
	}

	err = db.Open()
	if err != nil {
		l.Errorf("app - Run - Database connection is not open: %w", err)
		return
	}

	// Configure repos and services
	repos := repository.NewRepositories(db)
	svs := service.NewServices(repos)

	// Configure router
	h := gin.New()
	v1.NewRouter(h, svs)

	// Start http server
	httpServer := httpserver.NewHttpServer(h, &cfg.Http)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Errorf("app - Run - httpServer.Shutdown: %w", err)
	}

	err = db.Close()
	if err != nil {
		l.Errorf("app - Run - database.Close: %w", err)
	}

	l.Info("app - Run - is stopped")
}
