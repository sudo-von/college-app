package handler

import (
	"errors"
	"net/http"
	"time"

	"freelancer/college-app/go/api/middleware"
	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/pkg/token"
	"freelancer/college-app/go/usecase/user_mood"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type UserMoodController struct {
	UserMoodService user_mood.Service
	TokenService    token.Service
}

func NewUserMoodController(userMood user_mood.Service, token token.Service) *UserMoodController {
	return &UserMoodController{
		UserMoodService: userMood,
		TokenService:    token,
	}
}

func (c *UserMoodController) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.IsAuthorized(c.TokenService))
	r.Get("/users/{id}", c.Show)
	r.Post("/", c.Create)
	return r
}

// @tags users-mood
// @summary Show user's mood.
// @description Get user's mood given its ID for the current day.
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
// @description Create user's mood for the current day.
// @security BearerJWT
// @id create-user-mood
// @param payload body presenter.UserMoodPayload true "User's mood for the current day."
// @success 201
// @router /users-mood [post]
func (c *UserMoodController) Create(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(middleware.ContextKeyUserID).(string)
	if !ok {
		err := errors.New("user not in context")
		CheckError(err, w, r)
		return
	}

	var data presenter.UserMoodPayload
	if err := render.Bind(r, &data); err != nil {
		CheckError(err, w, r)
		return
	}

	currentDate := time.Now().In(time.Local)
	userMoodFilters := entity.UserMoodFilters{
		CreationDate: &currentDate,
	}

	newUserMood := entity.UserMoodPayload{
		UserID:       userID,
		Mood:         data.Mood,
		CreationDate: time.Now().In(time.Local),
	}

	err := c.UserMoodService.CreateUserMood(newUserMood, userMoodFilters)
	if err != nil {
		CheckError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
