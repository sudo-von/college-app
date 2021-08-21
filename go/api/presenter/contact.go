package presenter

import (
	"errors"
	"net/http"
	"strings"
	"time"
)

type ContactResponse struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	ContactName   string    `json:"contact_name"`
	ContactNumber string    `json:"contact_number"`
	Message       string    `json:"message"`
	CreationDate  time.Time `json:"creation_date"`
}

type ContactPayload struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	ContactName   string    `json:"contact_name"`
	ContactNumber string    `json:"contact_number"`
	Message       string    `json:"message"`
	CreationDate  time.Time `json:"creation_date"`
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
