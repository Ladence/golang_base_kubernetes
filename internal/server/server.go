package server

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	log *logrus.Logger
}

func NewServer(log *logrus.Logger) *Server {
	return &Server{
		log: log,
	}
}

func (s *Server) Run(port string) error {
	r := NewRouter()
	return http.ListenAndServe(":"+port, r)
}
