package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bradford-hamilton/stats-api/internal/server"
)

func main() {
	s := server.New()
	port := os.Getenv("STATS_API_PORT")
	if port == "" {
		port = "4000"
	}

	fmt.Printf("serving application on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, s.Mux))
}
