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
	ID             string           `json:"id"`
	User           TinyUserResponse `json:"user"`
	Subject        string           `json:"subject"`
	AdviceDate     string           `json:"advice_date"`
	Classroom      int              `json:"classroom"`
	StudentsNumber int              `json:"students_number"`
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

	return AdviceResponse{
		ID:             advice.ID,
		User:           user,
		Subject:        advice.Subject,
		AdviceDate:     advice.AdviceDate.Local().Format("2006-01-02 15:04"),
		Classroom:      advice.Classroom,
		StudentsNumber: advice.StudentsNumber,
	}
}

type AdvicePayload struct {
	Subject    string `json:"subject"`
	AdviceDate string `json:"advice_date"`
	Classroom  int    `json:"classroom"`
}

func (ap *AdvicePayload) validate() (err error) {
	if len(strings.TrimSpace(ap.Subject)) == 0 {
		err = mergeErrors(err, errors.New("missing field subject"))
	}
	if len(strings.TrimSpace(ap.AdviceDate)) == 0 {
		err = mergeErrors(err, errors.New("missing field advice_date"))
	}
	return
}

func (ap *AdvicePayload) Bind(r *http.Request) error {
	if err := ap.validate(); err != nil {
		return err
	}
	return nil
}
