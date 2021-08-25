package handler

import (
	"net/http"

	"freelancer/college-app/go/api/middleware"
	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/token"
	"freelancer/college-app/go/usecase/university"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type UniversityController struct {
	UniversityService university.Service
	TokenService      token.Service
}

func NewUniversityController(university university.Service, token token.Service) *UniversityController {
	return &UniversityController{
		UniversityService: university,
		TokenService:      token,
	}
}

func (c *UniversityController) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.IsAuthorized(c.TokenService))
	r.Get("/{id}", c.ShowUniversity)
	return r
}

// ShowUniversity renders a university given its id.
func (c *UniversityController) ShowUniversity(w http.ResponseWriter, r *http.Request) {

	universityID := chi.URLParam(r, "id")

	university, err := c.UniversityService.GetUniversityByID(universityID)
	if err != nil {
		CheckError(err, w, r)
		return
	}

	response := presenter.ToUniversityPresenter(*university)
	render.Status(r, http.StatusOK)
	render.Render(w, r, &response)
}
