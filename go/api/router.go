package api

import (
	"freelancer/college-app/go/api/handler"
	"log"
	"net/http"

	_ "freelancer/college-app/go/docs"

	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Services struct {
	UserService         handler.UserService
	UniversityService   handler.UniversityService
	SuggestionService   handler.SuggestionService
	ContactService      handler.ContactService
	DepartmentService   handler.DepartmentService
	UserMoodService     handler.UserMoodService
	AdviceService       handler.AdviceService
	InstalledAppService handler.InstalledAppService
	TokenService        handler.TokenService
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
	r.Use(chimiddleware.Recoverer)
	r.Use(chimiddleware.URLFormat)
	r.Use(chimiddleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// CORS configuration.
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "File-Name"},
		ExposedHeaders:   []string{"Link", "File-Name", "Authorization"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(cors.Handler)

	// Controllers.
	jwtController := handler.NewJWTController(services.UserService, services.TokenService)
	adviceController := handler.NewAdviceController(services.AdviceService, jwtController.IsAuthorized())
	contactController := handler.NewContactController(services.ContactService, jwtController.IsAuthorized())
	departmentController := handler.NewDepartmentController(services.DepartmentService, jwtController.IsAuthorized())
	suggestionController := handler.NewSuggestionController(services.SuggestionService, jwtController.IsAuthorized())
	userController := handler.NewUserController(services.UserService, jwtController.IsAuthorized())
	userMoodController := handler.NewUserMoodController(services.UserMoodService, jwtController.IsAuthorized())
	universityController := handler.NewUniversityController(services.UniversityService, jwtController.IsAuthorized())
	installedAppController := handler.NewInstalledAppController(services.InstalledAppService, jwtController.IsAuthorized())

	// Documentation.
	r.Mount("/swagger", httpSwagger.WrapHandler)

	// Http handlers.
	r.Mount("/advices", adviceController.Routes())
	r.Mount("/contacts", contactController.Routes())
	r.Mount("/departments", departmentController.Routes())
	r.Mount("/installed-apps", installedAppController.Routes())
	r.Mount("/suggestions", suggestionController.Routes())
	r.Mount("/users", userController.Routes())
	r.Mount("/users-mood", userMoodController.Routes())
	r.Mount("/universities", universityController.Routes())
	r.Mount("/auth", jwtController.Routes())

	// Start http server.
	if err := http.ListenAndServe(":4000", r); err != nil {
		log.Panic("[router] ListenAndServe error: %w", err.Error())
	}
}
