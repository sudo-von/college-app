package presenter

import (
	"errors"
	"freelancer/college-app/go/entity"
	"net/http"
	"strings"
)

var (
	ErrAdvivceNotFound = "ADVICE_NOT_FOUND"
	ErrInvAdviceDate   = "INVALID_ADVICE_DATE"
)

type AdviceList struct {
	Total   int              `json:"total" example:"1"`
	Advices []AdviceResponse `json:"results"`
}

func (al *AdviceList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type AdviceResponse struct {
	ID                 string            `json:"id" example:"613aab578a6ef50007e622be"`
	Subject            string            `json:"subject" example:"Chemistry"`
	AdviceDate         string            `json:"advice_date" example:"2021-10-24 15:04"`
	StudentsWillAttend []string          `json:"students_will_attend"`
	Classroom          ClassroomResponse `json:"classroom"`
	User               TinyUserResponse  `json:"user"`
}

func (ar *AdviceResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ToAdvicePresenter(advice entity.Advice) AdviceResponse {

	user := TinyUserResponse{
		ID:                 advice.User.ID,
		Name:               advice.User.Name,
		BirthDate:          advice.User.BirthDate.UTC().Format("2006-01-02 15:04"),
		Email:              advice.User.Email,
		RegistrationNumber: advice.User.RegistrationNumber,
	}

	classroom := ClassroomResponse{
		ID:   advice.Classroom.ID,
		Name: advice.Classroom.Name,
	}

	return AdviceResponse{
		ID:                 advice.ID,
		User:               user,
		Subject:            advice.Subject,
		AdviceDate:         advice.AdviceDate.Local().Format("2006-01-02 15:04"),
		Classroom:          classroom,
		StudentsWillAttend: advice.StudentsWillAttend,
	}
}

type AdvicePayload struct {
	Subject     string `json:"subject" example:"Chemistry"`
	AdviceDate  string `json:"advice_date" example:"2021-10-24 15:04"`
	ClassroomID string `json:"classroom_id" example:"613aab578a6ef50007e622be"`
}

func (ap *AdvicePayload) validate() (err error) {
	if len(strings.TrimSpace(ap.Subject)) == 0 {
		err = mergeErrors(err, errors.New("missing field subject"))
	}
	if len(strings.TrimSpace(ap.AdviceDate)) == 0 {
		err = mergeErrors(err, errors.New("missing field advice_date"))
	}
	if len(strings.TrimSpace(ap.ClassroomID)) == 0 {
		err = mergeErrors(err, errors.New("missing field classroom_id"))
	}
	return
}

func (ap *AdvicePayload) Bind(r *http.Request) error {
	if err := ap.validate(); err != nil {
		return err
	}
	return nil
}

type UpdateAdvicePayload struct {
	Subject     string `json:"subject" example:"Chemistry"`
	AdviceDate  string `json:"advice_date" example:"1997-04-17 15:04"`
	ClassroomID string `json:"classroom_id" example:"613aab578a6ef50007e622be"`
}

func (ap *UpdateAdvicePayload) validate() (err error) {
	if len(strings.TrimSpace(ap.Subject)) == 0 {
		err = mergeErrors(err, errors.New("missing field subject"))
	}
	if len(strings.TrimSpace(ap.AdviceDate)) == 0 {
		err = mergeErrors(err, errors.New("missing field advice_date"))
	}
	if len(strings.TrimSpace(ap.ClassroomID)) == 0 {
		err = mergeErrors(err, errors.New("missing field classroom_id"))
	}
	return
}

func (ap *UpdateAdvicePayload) Bind(r *http.Request) error {
	if err := ap.validate(); err != nil {
		return err
	}
	return nil
}
