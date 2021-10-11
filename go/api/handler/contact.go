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

type ContactService interface {
	GetContactByUserID(userID, requestedUserID string) (*entity.Contact, error)
	CreateContact(userID, requestedUserID string, newContact entity.ContactPayload) error
	UpdateContactByID(userID, contactID string, newContact entity.UpdateContactPayload) error
}

type ContactController struct {
	ContactService ContactService
	AuthService    func(http.Handler) http.Handler
}

func NewContactController(contactService ContactService, authService func(http.Handler) http.Handler) *ContactController {
	return &ContactController{
		ContactService: contactService,
		AuthService:    authService,
	}
}

func (c *ContactController) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(c.AuthService)
	r.Get("/users/{id}", c.Show)
	r.Post("/users/{id}", c.Create)
	r.Patch("/{id}", c.Update)
	return r
}

// @tags contacts
// @summary Show contact by user ID.
// @description Show contact given the user ID.
// @security BearerJWT
// @id show-contact-by-user-id
// @param id path string true "User ID."
// @produce json
// @success 200 {object} presenter.ContactResponse
// @router /contacts/users/{id} [get]
func (c *ContactController) Show(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(ContextKeyUserID).(string)
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

// @tags contacts
// @summary Create contact.
// @description Create contact for a specific user given its ID.
// @security BearerJWT
// @param id path string true "User ID."
// @param payload body presenter.ContactPayload true "Contact that wants to be stored."
// @id create-contact-by-user-id
// @success 201
// @router /contacts/users/{id} [post]
func (c *ContactController) Create(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(ContextKeyUserID).(string)
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

	w.WriteHeader(http.StatusCreated)
}

// @tags contacts
// @summary Update contact.
// @description Update contact given its ID.
// @security BearerJWT
// @id update-contact-by-id
// @param id path string true "Contact ID."
// @param updatePayload body presenter.UpdateContactPayload true "Contact information that wants to be updated."
// @success 200
// @router /contacts/{id} [patch]
func (c *ContactController) Update(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(ContextKeyUserID).(string)
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

	w.WriteHeader(http.StatusOK)
}
