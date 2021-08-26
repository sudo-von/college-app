package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"freelancer/college-app/go/api/middleware"
	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/token"
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

// Show returns a user mood given the user id from the context.
func (c *UserMoodController) Show(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(middleware.ContextKeyUserID).(string)
	if !ok {
		err := errors.New("user not in context")
		CheckError(entity.NewErrorInternalServer(fmt.Errorf("UserMoodController > Show: %w", err)), w, r)
		return
	}
	requestedUserID := chi.URLParam(r, "id")

	userMoodFilters := entity.UserMoodFilters{
		CreationDate: ParamToDate("creation_date", r.URL.Query()),
	}

	userMood, err := c.UserMoodService.GetUserMoodByUserID(userID, requestedUserID, userMoodFilters)
	if err != nil {
		CheckError(fmt.Errorf("UserMoodController > Show > GetUserMoodByUserID: %w", err), w, r)
		return
	}

	response := presenter.ToUserMoodPresenter(*userMood)
	render.Status(r, http.StatusOK)
	render.Render(w, r, &response)
}

// Create stores a new user mood.
func (c *UserMoodController) Create(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(middleware.ContextKeyUserID).(string)
	if !ok {
		err := errors.New("user not in context")
		CheckError(entity.NewErrorInternalServer(fmt.Errorf("UserMoodController > Create: %w", err)), w, r)
		return
	}

	var data presenter.UserMoodPayload
	if err := render.Bind(r, &data); err != nil {
		CheckError(entity.NewErrorBadRequest(fmt.Errorf("UserMoodController > Create: %w", err)), w, r)
		return
	}

	newUserMood := entity.UserMoodPayload{
		UserID:       userID,
		Mood:         data.Mood,
		CreationDate: time.Now().In(time.Local),
	}

	err := c.UserMoodService.CreateUserMood(newUserMood)
	if err != nil {
		CheckError(fmt.Errorf("UserMoodController > Create > CreateUserMood: %w", err), w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	render.Status(r, http.StatusOK)
}
