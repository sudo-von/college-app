package handler

import (
	"errors"
	"net/http"
	"time"

	"freelancer/college-app/go/api/middleware"
	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/pkg/token"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type UserMoodService interface {
	GetUserMoodByUserID(userID, requestedUserID string, userMoodFilters entity.UserMoodFilters) (*entity.UserMood, error)
	CreateUserMood(userID, requestedUserID string, newUserMood entity.UserMoodPayload) error
}

type UserMoodController struct {
	UserMoodService UserMoodService
	TokenService    token.Service
}

func NewUserMoodController(userMood UserMoodService, token token.Service) *UserMoodController {
	return &UserMoodController{
		UserMoodService: userMood,
		TokenService:    token,
	}
}

func (c *UserMoodController) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.IsAuthorized(c.TokenService))
	r.Get("/users/{id}", c.Show)
	r.Post("/users/{id}", c.Create)
	return r
}

// @tags users-mood
// @summary Show user's mood.
// @description Get user's mood for the current day given the user ID.
// @security BearerJWT
// @id get-user-mood-by-user-id
// @produce json
// @success 200 {object} presenter.UserMoodResponse
// @param id path string true "User ID."
// @router /users-mood/users/{id} [get]
func (c *UserMoodController) Show(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(middleware.ContextKeyUserID).(string)
	if !ok {
		err := errors.New("user not in context")
		CheckError(err, w, r)
		return
	}
	requestedUserID := chi.URLParam(r, "id")

	userMoodFilters := entity.UserMoodFilters{
		CreationDate: ParamToDate("creation_date", r.URL.Query()),
	}

	userMood, err := c.UserMoodService.GetUserMoodByUserID(userID, requestedUserID, userMoodFilters)
	if err != nil {
		CheckError(err, w, r)
		return
	}

	response := presenter.ToUserMoodPresenter(*userMood)
	render.Status(r, http.StatusOK)
	render.Render(w, r, &response)
}

// @tags users-mood
// @summary Create user's mood.
// @description Create user's mood for the current day given the user ID.
// @security BearerJWT
// @id create-user-mood
// @param id path string true "User ID."
// @param payload body presenter.UserMoodPayload true "User's mood for the current day."
// @success 201
// @router /users-mood/{users}/{id} [post]
func (c *UserMoodController) Create(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(middleware.ContextKeyUserID).(string)
	if !ok {
		err := errors.New("user not in context")
		CheckError(err, w, r)
		return
	}
	requestedUserID := chi.URLParam(r, "id")

	var data presenter.UserMoodPayload
	if err := render.Bind(r, &data); err != nil {
		CheckError(err, w, r)
		return
	}

	newUserMood := entity.UserMoodPayload{
		UserID:       requestedUserID,
		Mood:         data.Mood,
		CreationDate: time.Now().In(time.Local),
	}

	err := c.UserMoodService.CreateUserMood(userID, requestedUserID, newUserMood)
	if err != nil {
		CheckError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
