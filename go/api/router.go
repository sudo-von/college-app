package api

import (
	"freelancer/college-app/go/api/middleware"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type Services struct {
}

func ListenAndServe(services Services) {
	// Router.
	r := chi.NewRouter()
	// Middleware configuration.
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)
	r.Use(chimiddleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	// CORS configuration.
	cors := middleware.Cors()
	r.Use(cors.Handler)

	// Http hanlders.

	// Start http server.
	if err := http.ListenAndServe(":4000", r); err != nil {
		log.Printf("[routes] error: %s", err.Error())
	}
}
