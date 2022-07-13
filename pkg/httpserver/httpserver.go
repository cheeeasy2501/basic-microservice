package httpserver

import (
	"context"
	"net/http"
	"time"
)

const (
	_defaultShutdownTimeout = 3 * time.Second
)

type IHttpServer interface {
	start()
	Notify() <-chan error
	Shutdown() error
}

type HttpServer struct {
	srv             *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

func NewHttpServer(handler http.Handler, cfg *HttpServerConfig) *HttpServer {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
		Addr:         cfg.GetAddr(),
	}

	s := &HttpServer{
		srv:             httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
	}

	s.start()

	return s
}

func (s *HttpServer) start() {
	go func() {
		s.notify <- s.srv.ListenAndServe()
		close(s.notify)
	}()
}

func (s *HttpServer) Notify() <-chan error {
	return s.notify
}

func (s *HttpServer) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.srv.Shutdown(ctx)
}

type HttpResponse struct {
	Message  string      `json:"message"`
	AppError string      `json:"appError,omitempty"`
	Data     interface{} `json:"data"`
}

func NewResponse(message string, data interface{}) *HttpResponse {
	return &HttpResponse{
		Message: message,
		Data:    data,
	}
}

func NewErrorResponse(message string, appErr error) *HttpResponse {
	r := &HttpResponse{
		Message: message,
		Data:    nil,
	}

	if appErr != nil {
		r.AppError = appErr.Error()
	}

	return r
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
