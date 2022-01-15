package main

import (
	"github.com/Ladence/golang_base_kubernetes/internal/server"
	log "github.com/sirupsen/logrus"
	"os"
)

func logInit() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}

func main() {
	logInit()

	logger := log.New()
	logger.Info("Service started")
	srv := server.NewServer(logger)
	logger.Info("The service is ready to listen and serve.")
	logger.Fatal(srv.Run("8080"))
}
