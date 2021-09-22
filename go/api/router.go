package api

import (
	"freelancer/college-app/go/api/handler"
	"freelancer/college-app/go/api/middleware"
	"freelancer/college-app/go/pkg/token"
	"freelancer/college-app/go/usecase/advice"
	"freelancer/college-app/go/usecase/contact"
	"freelancer/college-app/go/usecase/suggestion"
	"freelancer/college-app/go/usecase/university"
	"freelancer/college-app/go/usecase/user"
	"freelancer/college-app/go/usecase/user_mood"
	"log"
	"net/http"

	_ "freelancer/college-app/go/docs"

	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Services struct {
	UserService       user.Service
	UniversityService university.Service
	SuggestionService suggestion.Service
	ContactService    contact.Service
	UserMoodService   user_mood.Service
	AdviceService     advice.Service
	TokenService      token.Service
}

// @title College-app API
// @version 1.0.0
// @description Official documentation to consume the API.

// @contact.name Jesús 'VoN' Rodríguez
// @contact.url https://www.linkedin.com/in/jes%C3%BAs-%C3%A1ngel-rodr%C3%ADguez-mart%C3%ADnez-84991a1b4/
// @contact.email sudo.von.contact@gmail.com

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apiKey BearerJWT
// @in header
// @name Authorization
// @tokenUrl http://localhost.com:4000/users/login
func ListenAndServe(services Services) {

	r := chi.NewRouter()
	// Middleware configuration.
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)
	r.Use(chimiddleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// CORS configuration.
	cors := middleware.Cors()
	r.Use(cors.Handler)

	// Documentation.
	r.Get("/swagger", httpSwagger.WrapHandler)

	// Http handlers.
	r.Mount("/advices", handler.NewAdviceController(services.AdviceService, services.TokenService).Routes())
	r.Mount("/contacts", handler.NewContactController(services.ContactService, services.TokenService).Routes())
	r.Mount("/suggestions", handler.NewSuggestionController(services.SuggestionService, services.TokenService).Routes())
	r.Mount("/users", handler.NewUserController(services.UserService, services.TokenService).Routes())
	r.Mount("/users-mood", handler.NewUserMoodController(services.UserMoodService, services.TokenService).Routes())
	r.Mount("/universities", handler.NewUniversityController(services.UniversityService, services.TokenService).Routes())

	// Start http server.
	if err := http.ListenAndServe(":4000", r); err != nil {
		log.Panic("[router] ListenAndServe error: %w", err.Error())
	}
}
