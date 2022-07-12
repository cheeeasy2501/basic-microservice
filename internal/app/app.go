package app

import (
	"basic-microservice/internal/config"
	v1 "basic-microservice/internal/transport/http/v1"
	"basic-microservice/pkg/httpserver"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type App struct {
}

// Run func (a *App) Run(cfg *config.Config) {
func Run(cfg *config.Config) {
	l := log.New()
	l.SetLevel(log.DebugLevel) //cfg.Log.Level
	ch := make(chan struct{})
	h := gin.New()
	v1.NewRouter(h)
	//httpServer
	_ = httpserver.NewHttpServer(h)
	fmt.Println("Infinity channel")
	<-ch

	// Waiting signal
	//interrupt := make(chan os.Signal, 1)
	//signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	//
	//select {
	//case s := <-interrupt:
	//	l.Info("app - Run - signal: " + s.String())
	//case err = <-httpServer.Notify():
	//	l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	//case err = <-rmqServer.Notify():
	//	l.Error(fmt.Errorf("app - Run - rmqServer.Notify: %w", err))
	//}
	//
	//// Shutdown
	//err = httpServer.Shutdown()
	//if err != nil {
	//	l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	//}
}
