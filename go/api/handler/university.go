package handler

import (
	"errors"
	"net/http"

	"freelancer/college-app/go/api/middleware"
	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/pkg/token"
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
	r.Get("/", c.List)
	r.Group(func(r chi.Router) {
		r.Use(middleware.IsAuthorized(c.TokenService))
		r.Get("/{id}", c.Show)
	})
	return r
}

// @tags universities
// @summary List universities.
// @description List basic universities information.
// @id list-universities
// @produce json
// @success 200 {object} presenter.TinyUniversityList
// @router /universities [get]
func (c *UniversityController) List(w http.ResponseWriter, r *http.Request) {

	list, total, err := c.UniversityService.GetTinyUniversities()
	if err != nil {
		CheckError(err, w, r)
		return
	}

	res := presenter.TinyUniversityList{
		Total:            *total,
		TinyUniversities: make([]presenter.TinyUniversityResponse, 0, len(list)),
	}

	for _, university := range list {
		res.TinyUniversities = append(res.TinyUniversities, presenter.ToTinyUniversityPresenter(university))
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, &res)
}

// @tags universities
// @summary Show university.
// @description Show basic university information.
// @id show-university
// @produce json
// @security BearerJWT
// @success 200 {object} presenter.UniversityResponse
// @param id path string true "University ID."
// @router /universities/{id} [get]
func (c *UniversityController) Show(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(middleware.ContextKeyUserID).(string)
	if !ok {
		err := errors.New("user not in context")
		CheckError(err, w, r)
		return
	}
	universityID := chi.URLParam(r, "id")

	university, err := c.UniversityService.GetUniversityByID(userID, universityID)
	if err != nil {
		CheckError(err, w, r)
		return
	}

	response := presenter.ToUniversityPresenter(*university)
	render.Status(r, http.StatusOK)
	render.Render(w, r, &response)
}
