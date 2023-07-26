package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"go.uber.org/zap"
)

const mlbStatsBaseURL = "https://statsapi.mlb.com"

// API is a structure that holds dependencies and provides
// methods for orchestrating http server interactions.
type API struct {
	Mux        *chi.Mux
	httpClient HTTPClient
	log        *zap.Logger
}

// HTTPClient is the interface that must be implemented by an API's httpClient.
// This will allow us to mock outgoing calls during tests.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// New creates a router, sets up middleware, and initializes routes and handlers.
func New(client HTTPClient, logger *zap.Logger) *API {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.RequestID,
		middleware.StripSlashes,            // strip slashes to no slash URL versions
		middleware.Recoverer,               // recover from panics without crashing server
		middleware.Timeout(30*time.Second), // start with a pretty standard timeout
	)

	api := &API{
		Mux:        router,
		httpClient: client,
		log:        logger,
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
