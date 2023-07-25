package server

import (
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

// API is a structure that holds dependencies and provides
// methods for orchestrating http server interactions.
type API struct {
	baseURL    string
	Mux        *chi.Mux
	httpClient HTTPClient
}

// HTTPClient is the interface that must be implemented by an API's httpClient.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

const mlbStatsBaseURL = "https://statsapi.mlb.com"

// New creates a router, sets up middleware, and initalizes routes and handlers.
func New(client HTTPClient) *API {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.RequestID,
		middleware.StripSlashes,            // strip slashes to no slash URL versions
		middleware.Recoverer,               // recover from panics without crashing server
		middleware.Timeout(30*time.Second), // start with a pretty standard timeout
	)

	baseURL := "http://localhost:4000"
	if os.Getenv("STATS_API_ENVIRONMENT") == "production" {
		baseURL = "todo_production_endpoint"
	}

	api := &API{
		baseURL:    baseURL,
		Mux:        router,
		httpClient: client,
	}

	api.initializeRoutes()

	return api
}

func (a *API) initializeRoutes() {
	a.Mux.Get("/ping", a.ping)
	a.Mux.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/schedule", func(r chi.Router) {
				r.Get("/", a.getSchedule)
			})
		})
	})
}
