package api

import (
	"freelancer/college-app/go/api/handler"
	"freelancer/college-app/go/api/middleware"
	"freelancer/college-app/go/token"
	"freelancer/college-app/go/usecase/university"
	"freelancer/college-app/go/usecase/user"
	"log"
	"net/http"

	chimiddleware "github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Services struct {
	UserService       user.Service
	UniversityService university.Service
	TokenService      token.Service
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
	// Token.

	// Http hanlders.
	r.Mount("/users", handler.NewUserController(services.UserService, services.TokenService).Routes())

	// Start http server.
	if err := http.ListenAndServe(":4000", r); err != nil {
		log.Printf("[routes] error: %s", err.Error())
	}
}
