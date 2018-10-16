// Simple REST API server based on Chi.

package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/bwv988/gostatsapi/src/api/v1/about"
	"github.com/bwv988/gostatsapi/src/api/v1/avg"
	"github.com/bwv988/gostatsapi/src/api/v1/max"
	"github.com/bwv988/gostatsapi/src/api/v1/median"
	"github.com/bwv988/gostatsapi/src/api/v1/min"
	"github.com/bwv988/gostatsapi/src/api/v1/percentile"
)

/*
	FIXME
*/
const (
	svrPort = ":8080"
)

/*
	FIXME
*/
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	// Implement a versioned REST API.
	router.Route("/api", func(r chi.Router) {
		r.Mount("/v1/about", about.Routes())
		r.Mount("/v1/min", min.Routes())
		r.Mount("/v1/max", max.Routes())
		r.Mount("/v1/avg", avg.Routes())
		r.Mount("/v1/median", median.Routes())
		r.Mount("/v1/percentile", percentile.Routes())
	})

	return router
}

func main() {
	router := Routes()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route) // Walk and print out all routes
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error()) // panic if there is an error
	}

	log.Fatal(http.ListenAndServe(svrPort, router))
}
