package app

import (
	"basic-microservice/internal/config"
	"basic-microservice/internal/repository"
	"basic-microservice/internal/service"
	v1 "basic-microservice/internal/transport/http/v1"
	"basic-microservice/pkg/httpserver"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
}

// Run func (a *App) Run(cfg *config.Config) {
func Run(cfg *config.Config, l *logrus.Logger) {
	l.Info("app - Run - Run app")

	// Configure gorm DB
	db := &gorm.DB{} // mock
	// Configure repos and services
	repos := repository.NewRepositories(db)
	svs := service.NewServices(repos)
	// Configure router
	h := gin.New()
	v1.NewRouter(h, svs)

	// Start http server
	httpCfg, err := httpserver.NewConfig()
	if err != nil {
		fmt.Println("app - Run - Bad httpServer config")
	}

	httpServer := httpserver.NewHttpServer(h, httpCfg)

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
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	l.Info("app - Run - is stopped")
}
