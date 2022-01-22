package main

import (
	"context"
	"github.com/Ladence/golang_base_kubernetes/internal/server"
	"github.com/Ladence/golang_base_kubernetes/internal/version"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
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

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM)

	logger.Infof(
		"Starting the service...\ncommit: %s, build time: %s, release: %s",
		version.Commit, version.BuildTime, version.Release,
	)
	logger.Info("The service is ready to listen and serve.")
	srv := server.NewServer(logger, port)
	go func() {
		logger.Fatal(srv.Run())
	}()

	sig := <-sigchan
	logger.Infof("Got signal: %v", sig)
	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Fatalf("Error on shutting down server. Error: %v", err)
	}
}
