package presenter

import (
	"errors"
	"freelancer/college-app/go/entity"
	"net/http"
	"strings"
)

type AdviceList struct {
	Total   int              `json:"total"`
	Advices []AdviceResponse `json:"results"`
}

func (al *AdviceList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type AdviceResponse struct {
	ID             string            `json:"id"`
	Subject        string            `json:"subject"`
	AdviceDate     string            `json:"advice_date"`
	StudentsNumber int               `json:"students_number"`
	Classroom      ClassroomResponse `json:"classroom"`
	User           TinyUserResponse  `json:"user"`
}

func (ar *AdviceResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ToAdvicePresenter(advice entity.Advice) AdviceResponse {

	user := TinyUserResponse{
		ID:                 advice.User.ID,
		Name:               advice.User.Name,
		BirthDate:          advice.User.BirthDate.UTC().Format("2006-01-02"),
		Email:              advice.User.Email,
		RegistrationNumber: advice.User.RegistrationNumber,
	}

	classroom := ClassroomResponse{
		ID:   advice.Classroom.ID,
		Name: advice.Classroom.Name,
	}

	return AdviceResponse{
		ID:             advice.ID,
		User:           user,
		Subject:        advice.Subject,
		AdviceDate:     advice.AdviceDate.Local().Format("2006-01-02 15:04"),
		Classroom:      classroom,
		StudentsNumber: advice.StudentsNumber,
	}
}

type AdvicePayload struct {
	Subject     string `json:"subject"`
	AdviceDate  string `json:"advice_date"`
	ClassroomID string `json:"classroom_id"`
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
