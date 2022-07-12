package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const (
	_defaultReadTimeout  = 5 * time.Second
	_defaultWriteTimeout = 5 * time.Second
	_defaultAddr         = ":80"
	//_defaultShutdownTimeout = 3 * time.Second
)

type IHttpServer interface {
	Start()
	ConfiguringServer(config *HttpServerConfig)
	Stop()
}

type HttpServer struct {
	srv *http.Server
}

func NewHttpServer(handler http.Handler) *HttpServer {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         _defaultAddr,
	}

	s := &HttpServer{
		srv: httpServer,
	}

	s.start()

	return s
}

func (s *HttpServer) start() {
	go func() {
		s.srv.ListenAndServe()
	}()
}

// todo: check it
func (s *HttpServer) configuringServer(config *HttpServerConfig) {
	s.srv.Addr = fmt.Sprintf("%s:%s", config.Host, config.Port)
}

// todo: need more info about shutdown
func (s *HttpServer) Stop() {
	s.srv.Shutdown(context.Background())
}

type HttpErrorResponse struct {
	Message          string
	DeveloperMessage string
	Code             uint8
}

type HttpResponse struct {
	Data interface{}
}

/// a instance response
//{
//	data: {
//		entity_field_1: "123",
//		entity_field_2: "345",
//	}
//	errors: nil,
//}

/// a lot of instance response
//{
//	data: [
//		{
//			entity_field_1: "123",
//			entity_field_2: "345",
//		},
//		{
//			entity_field_1: "123",
//			entity_field_2: "345",
//		},
//	]
//	errors: nil,
//}
