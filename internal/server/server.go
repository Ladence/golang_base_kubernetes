package server

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	log *logrus.Logger
	srv http.Server
}

func NewServer(log *logrus.Logger, port string) *Server {
	return &Server{
		log: log,
		srv: http.Server{
			Addr:    ":" + port,
			Handler: NewRouter(),
		},
	}
}

func (s *Server) Run() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.log.Info("Service is shutting down...")
	return s.srv.Shutdown(ctx)
}
