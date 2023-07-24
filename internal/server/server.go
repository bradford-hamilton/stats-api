package server

import (
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

// API is a structure that holds dependencies and provides
// methods for orchestrating http server interactions.
type API struct {
	baseURL string
	Mux     *chi.Mux
}

// New creates a router, sets up middleware, and initalizes routes and handlers.
func New() *API {
	r := chi.NewRouter()
	r.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.StripSlashes,            // strip slashes to no slash URL versions
		middleware.Recoverer,               // recover from panics without crashing server
		middleware.Timeout(30*time.Second), // start with a pretty standard timeout
	)

	baseURL := "http://localhost:4000"
	if os.Getenv("STATS_API_ENVIRONMENT") == "production" {
		baseURL = "todo_production_endpoint"
	}

	api := &API{baseURL: baseURL, Mux: r}
	api.initializeRoutes()

	return api
}

func (a *API) initializeRoutes() {
	a.Mux.Get("/ping", a.ping)
}
