package handler

import (
	"errors"
	"net/http"
	"time"

	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/entity"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type UserService interface {
	GetTinyUserByID(userID, requestedUserID string) (*entity.TinyUser, error)
	CreateUser(newUser entity.UserPayload) error
	UpdateTinyUser(userID, requestedUserID string, newUser entity.UpdateUserPayload) error
	AuthenticateUser(email, password string) (*entity.User, error)
}

type UserController struct {
	UserService UserService
	AuthService func(http.Handler) http.Handler
}

func NewUserController(user UserService, authService func(http.Handler) http.Handler) *UserController {
	return &UserController{
		UserService: user,
		AuthService: authService,
	}
}

func (c *UserController) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", c.Create)
	r.Group(func(r chi.Router) {
		r.Use(c.AuthService)
		r.Get("/{id}", c.GetTinyUser)
		r.Patch("/{id}", c.UpdateTinyUser)
	})
	return r
}

// @tags users
// @summary Show user.
// @description Show basic user information.
// @security BearerJWT
// @id get-tiny-user
// @produce json
// @success 200 {object} presenter.TinyUserResponse
// @param id path string true "User ID."
// @router /users/{id} [get]
func (c *UserController) GetTinyUser(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(ContextKeyUserID).(string)
	if !ok {
		err := errors.New("user not in context")
		CheckError(err, w, r)
		return
	}
	requestedUserID := chi.URLParam(r, "id")

	user, err := c.UserService.GetTinyUserByID(userID, requestedUserID)
	if err != nil {
		CheckError(err, w, r)
		return
	}

	response := presenter.ToTinyPresenterUser(*user)
	render.Status(r, http.StatusOK)
	render.Render(w, r, &response)
}

// @tags users
// @summary Create user.
// @description Create user.
// @id create-user
// @param payload body presenter.UserPayload true "User information that wants to be stored."
// @success 201
// @router /users [post]
func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {

	var data presenter.UserPayload
	if err := render.Bind(r, &data); err != nil {
		CheckError(err, w, r)
		return
	}

	birthDate, err := time.ParseInLocation("2006-01-02", data.BirthDate, time.Local)
	if err != nil {
		CheckError(err, w, r)
		return
	}
	newUser := entity.UserPayload{
		Name:               data.Name,
		BirthDate:          birthDate.In(time.Local),
		Email:              data.Email,
		RegistrationNumber: data.RegistrationNumber,
		Password:           data.Password,
		UniversityID:       data.UniversityID,
		Status:             entity.ActiveStatus,
		Role:               entity.StudentRole,
		CreationDate:       time.Now().In(time.Local),
	}

	err = c.UserService.CreateUser(newUser)
	if err != nil {
		CheckError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// @tags users
// @summary Update user.
// @description Update user information.
// @security BearerJWT
// @id update-user
// @param id path string true "User ID"
// @param updatePayload body presenter.UpdateUserPayload true "User information that wants to be updated."
// @success 200
// @router /users/{id} [PATCH]
func (c *UserController) UpdateTinyUser(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(ContextKeyUserID).(string)
	if !ok {
		err := errors.New("user not in context")
		CheckError(err, w, r)
		return
	}
	requestedUserID := chi.URLParam(r, "id")

	var data presenter.UpdateUserPayload
	if err := render.Bind(r, &data); err != nil {
		CheckError(err, w, r)
		return
	}

	birthDate, err := time.ParseInLocation("2006-01-02", data.BirthDate, time.Local)
	if err != nil {
		CheckError(err, w, r)
		return
	}
	newUser := entity.UpdateUserPayload{
		Name:               data.Name,
		BirthDate:          birthDate,
		Email:              data.Email,
		RegistrationNumber: data.RegistrationNumber,
	}

	err = c.UserService.UpdateTinyUser(userID, requestedUserID, newUser)
	if err != nil {
		CheckError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
}
