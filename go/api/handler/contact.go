package handler

import (
	"errors"
	"net/http"
	"time"

	"freelancer/college-app/go/api/middleware"
	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/pkg/token"
	"freelancer/college-app/go/usecase/contact"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type ContactController struct {
	ContactService contact.Service
	TokenService   token.Service
}

func NewContactController(contact contact.Service, token token.Service) *ContactController {
	return &ContactController{
		ContactService: contact,
		TokenService:   token,
	}
}

func (c *ContactController) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.IsAuthorized(c.TokenService))
	r.Get("/users/{id}", c.Show)
	r.Post("/users/{id}", c.Create)
	r.Patch("/{id}", c.Update)
	return r
}

// Show returns a contact given the user id from the context.
func (c *ContactController) Show(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(middleware.ContextKeyUserID).(string)
	if !ok {
		err := errors.New("user not in context")
		CheckError(err, w, r)
		return
	}
	requestedUserID := chi.URLParam(r, "id")

	contact, err := c.ContactService.GetContactByUserID(userID, requestedUserID)
	if err != nil {
		CheckError(err, w, r)
		return
	}

	response := presenter.ToContactPresenter(*contact)
	render.Status(r, http.StatusOK)
	render.Render(w, r, &response)
}

// Create stores a new user.
func (c *ContactController) Create(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(middleware.ContextKeyUserID).(string)
	if !ok {
		err := errors.New("user not in context")
		CheckError(err, w, r)
		return
	}
	requestedUserID := chi.URLParam(r, "id")

	var data presenter.ContactPayload
	if err := render.Bind(r, &data); err != nil {
		CheckError(err, w, r)
		return
	}

	newContact := entity.ContactPayload{
		UserID:        requestedUserID,
		ContactName:   data.ContactName,
		ContactNumber: data.ContactNumber,
		Message:       data.Message,
		CreationDate:  time.Now().In(time.Local),
	}

	err := c.ContactService.CreateContact(userID, requestedUserID, newContact)
	if err != nil {
		CheckError(err, w, r)
		return
	}

	render.Status(r, http.StatusCreated)
}

// Update updates a contact given the user id from the context.
func (c *ContactController) Update(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(middleware.ContextKeyUserID).(string)
	if !ok {
		err := errors.New("user not in context")
		CheckError(err, w, r)
		return
	}
	contactID := chi.URLParam(r, "id")

	var data presenter.UpdateContactPayload
	if err := render.Bind(r, &data); err != nil {
		CheckError(err, w, r)
		return
	}

	newContact := entity.UpdateContactPayload{
		ContactName:   data.ContactName,
		ContactNumber: data.ContactNumber,
		Message:       data.Message,
	}

	err := c.ContactService.UpdateContactByID(userID, contactID, newContact)
	if err != nil {
		CheckError(err, w, r)
		return
	}

	render.Status(r, http.StatusOK)
}
