package presenter

import (
	"errors"
	"freelancer/college-app/go/entity"
	"net/http"
	"strings"
)

type ContactResponse struct {
	ID            string `json:"id"`
	UserID        string `json:"user_id"`
	ContactName   string `json:"contact_name"`
	ContactNumber string `json:"contact_number"`
	Message       string `json:"message"`
}

func (ur *ContactResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ToContactPresenter(contact entity.Contact) ContactResponse {
	return ContactResponse{
		ID:            contact.ID,
		UserID:        contact.UserID,
		ContactName:   contact.ContactName,
		ContactNumber: contact.ContactNumber,
		Message:       contact.Message,
	}
}

type ContactPayload struct {
	ContactName   string `json:"contact_name"`
	ContactNumber string `json:"contact_number"`
	Message       string `json:"message"`
}

func (cp *ContactPayload) validate() (err error) {
	if len(strings.TrimSpace(cp.ContactName)) == 0 {
		err = mergeErrors(err, errors.New("missing field contact_name"))
	}
	if len(strings.TrimSpace(cp.ContactNumber)) == 0 {
		err = mergeErrors(err, errors.New("missing field contact_number"))
	}
	if len(strings.TrimSpace(cp.Message)) == 0 {
		err = mergeErrors(err, errors.New("missing field message"))
	}
	return
}

func (cp *ContactPayload) Bind(r *http.Request) error {
	if err := cp.validate(); err != nil {
		return err
	}
	return nil
}

type UpdateContactPayload struct {
	ContactName   string `json:"contact_name"`
	ContactNumber string `json:"contact_number"`
	Message       string `json:"message"`
}

func (ucp *UpdateContactPayload) validate() (err error) {
	if len(strings.TrimSpace(ucp.ContactName)) == 0 {
		err = mergeErrors(err, errors.New("missing field contact_name"))
	}
	if len(strings.TrimSpace(ucp.ContactNumber)) == 0 {
		err = mergeErrors(err, errors.New("missing field contact_number"))
	}
	if len(strings.TrimSpace(ucp.Message)) == 0 {
		err = mergeErrors(err, errors.New("missing field message"))
	}
	return
}

func (ucp *UpdateContactPayload) Bind(r *http.Request) error {
	if err := ucp.validate(); err != nil {
		return err
	}
	return nil
}
