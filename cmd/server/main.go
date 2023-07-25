package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bradford-hamilton/stats-api/internal/server"
	"go.uber.org/zap"
)

func main() {
	var logger *zap.Logger
	var err error

	if os.Getenv("STATS_API_ENVIRONMENT") == "production" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		log.Fatal("failed to create logger")
	}
	defer logger.Sync()

	client := http.Client{Timeout: 30 * time.Second}
	s := server.New(&client, logger)

	port := os.Getenv("STATS_API_PORT")
	if port == "" {
		port = "4000"
	}

	srv := &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: s.Mux,
	}

	// Implement a graceful shutdown. This allows in-flight requests to complete before
	// shutting down the server, preventing potential data loss or corruption.
	serverCtx, serverStopCtx := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig

		shutdownCtx, _ := context.WithTimeout(serverCtx, 10*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out, forcing exit")
			}
		}()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Fatal(err)
		}

		serverStopCtx()
	}()

	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	<-serverCtx.Done()
}
