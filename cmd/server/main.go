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
	srv := &http.Server{Addr: ":4000", Handler: s.Mux}

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
				logger.Fatal("graceful shutdown timed out, forcing exit")
			}
		}()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			logger.Fatal(err.Error())
		}

		serverStopCtx()
	}()

	logger.Info("Server listening on port 4000")

	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Fatal(err.Error())
	}

	<-serverCtx.Done()
}
