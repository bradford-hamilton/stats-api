package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bradford-hamilton/stats-api/internal/server"
)

func main() {
	c := http.Client{Timeout: 30 * time.Second}
	s := server.New(&c)
	port := os.Getenv("STATS_API_PORT")
	if port == "" {
		port = "4000"
	}

	fmt.Printf("serving application on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, s.Mux))
}
