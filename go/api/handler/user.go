package handler

import (
	"net/http"
	"time"

	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/usecase/user"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type UserController struct {
	UserService user.Service
}

func NewUserController(user user.Service) *UserController {
	return &UserController{
		UserService: user,
	}
}

func (c *UserController) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", c.Create)
	return r
}

// Create stores a new user.
func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {

	var data presenter.UserPayload
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, presenter.ErrInvalidRequest(err))
		return
	}

	loc, err := time.LoadLocation("America/Mexico_City")
	if err != nil {
		CheckError(err, w, r)
		return
	}

	birthDate, err := time.Parse("2006-01-02", data.BirthDate)
	if err != nil {
		CheckError(err, w, r)
		return
	}

	newUser := entity.UserPayload{
		Name:               data.Name,
		BirthDate:          birthDate,
		Email:              data.Email,
		RegistrationNumber: data.RegistrationNumber,
		Password:           data.Password,
		UniversityID:       data.UniversityID,
		Status:             entity.ActiveStatus,
		CreationDate:       time.Now().In(loc),
	}

	err = c.UserService.CreateUser(newUser)
	if err != nil {
		CheckError(err, w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	render.Status(r, http.StatusOK)
}
