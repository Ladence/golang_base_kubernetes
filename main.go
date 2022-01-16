package main

import (
	"github.com/Ladence/golang_base_kubernetes/internal/server"
	"github.com/Ladence/golang_base_kubernetes/internal/version"
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

	port := os.Getenv("PORT")
	if len(port) == 0 {
		logger.Fatal("Port is empty!")
	}

	logger.Infof(
		"Starting the service...\ncommit: %s, build time: %s, release: %s",
		version.Commit, version.BuildTime, version.Release,
	)
	srv := server.NewServer(logger)
	logger.Info("The service is ready to listen and serve.")
	logger.Fatal(srv.Run("8080"))
}
