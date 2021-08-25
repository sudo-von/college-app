package presenter

import (
	"freelancer/college-app/go/entity"
	"net/http"
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

	university := UniversityResponse{
		ID:             advice.User.University.ID,
		Name:           advice.User.University.Name,
		ProfilePicture: advice.User.University.ProfilePicture,
	}

	user := TinyUserResponse{
		ID:                 advice.User.ID,
		Name:               advice.User.Name,
		BirthDate:          advice.User.BirthDate.UTC().Format("2006-01-02"),
		Email:              advice.User.Email,
		RegistrationNumber: advice.User.RegistrationNumber,
		University:         university,
	}

	return AdviceResponse{
		ID:             advice.ID,
		User:           user,
		Subject:        advice.Subject,
		AdviceDate:     advice.AdviceDate.Local().Format("2006-01-02"),
		Classroom:      advice.Classroom,
		StudentsNumber: advice.StudentsNumber,
	}
}
