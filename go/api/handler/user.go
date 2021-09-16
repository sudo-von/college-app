package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"freelancer/college-app/go/api/middleware"
	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/pkg/token"
	"freelancer/college-app/go/usecase/user"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type UserController struct {
	UserService  user.Service
	TokenService token.Service
}

func NewUserController(user user.Service, token token.Service) *UserController {
	return &UserController{
		UserService:  user,
		TokenService: token,
	}
}

func (c *UserController) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", c.Create)
	r.Group(func(r chi.Router) {
		r.Use(middleware.BasicAuth(&c.UserService))
		r.Post("/login", c.Login)
	})
	r.Group(func(r chi.Router) {
		r.Use(middleware.IsAuthorized(c.TokenService))
		r.Get("/{id}", c.GetTinyUser)
		r.Patch("/{id}", c.UpdateTinyUser)
	})
	return r
}

// Show returns a user given its id from the context.
func (c *UserController) GetTinyUser(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(middleware.ContextKeyUserID).(string)
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

// Create stores a new user.
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

	render.Status(r, http.StatusCreated)
}

//Login authenticates a user if credentials are right.
func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {

	user, ok := r.Context().Value(middleware.ContextKeyUser).(*entity.User)
	if !ok {
		err := errors.New("user not in context")
		CheckError(err, w, r)
		return
	}

	signedToken, err := c.TokenService.CreateToken(user, 15)
	if err != nil {
		CheckError(err, w, r)
		return
	}

	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", signedToken))
	w.Header().Set("Access-Control-Allow-Headers", "Authorization")
	render.Status(r, http.StatusOK)
}

// UpdateTinyUser updates a user given its id from the context.
func (c *UserController) UpdateTinyUser(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(middleware.ContextKeyUserID).(string)
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

	render.Status(r, http.StatusOK)
}
