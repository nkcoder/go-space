package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"time"
)

func (app *application) routes() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.Timeout(60 * time.Second))

	// set custom handler (JSON response) for 404 and 405
	r.NotFound(app.notFoundResponse)
	r.MethodNotAllowed(app.methodNotAllowedResponse)

	r.Get("/v1/healthcheck", app.healthcheckHandler)
	r.Post("/v1/movies", app.createMovieHandler)
	r.Get("/v1/movies/{id}", app.showMovieHandler)
	r.Put("/v1/movies/{id}", app.updateMovieHandler)
	r.Delete("/v1/movies/{id}", app.deleteMovieHandler)

	return r
}
