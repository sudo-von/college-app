package presenter

import (
	"errors"
	"freelancer/college-app/go/entity"
	"net/http"
	"strings"
	"time"
)

type TinyUserResponse struct {
	ID                 string             `json:"id"`
	Name               string             `json:"name"`
	BirthDate          string             `json:"birth_date"`
	Email              string             `json:"email"`
	RegistrationNumber string             `json:"registration_number"`
	University         UniversityResponse `json:"university"`
}

func (ur *TinyUserResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ToTinyPresenterUser(user entity.TinyUser) TinyUserResponse {

	university := UniversityResponse{
		ID:             user.University.ID,
		Name:           user.University.Name,
		ProfilePicture: user.University.ProfilePicture,
	}

	return TinyUserResponse{
		ID:                 user.ID,
		Name:               user.Name,
		BirthDate:          user.BirthDate.UTC().Format("2006-01-02"),
		Email:              user.Email,
		RegistrationNumber: user.RegistrationNumber,
		University:         university,
	}
}

type UserPayload struct {
	ID                 string    `json:"id"`
	Name               string    `json:"name"`
	BirthDate          string    `json:"birth_date"`
	Email              string    `json:"email"`
	RegistrationNumber string    `json:"registration_number"`
	Password           string    `json:"password"`
	UniversityID       string    `json:"university_id"`
	Status             string    `json:"status"`
	CreationDate       time.Time `json:"creation_date"`
}

func (up *UserPayload) validate() (err error) {
	if len(strings.TrimSpace(up.Name)) == 0 {
		err = mergeErrors(err, errors.New("missing field name"))
	}
	if len(strings.TrimSpace(up.BirthDate)) == 0 {
		err = mergeErrors(err, errors.New("missing field birth_date"))
	}
	if len(strings.TrimSpace(up.Email)) == 0 {
		err = mergeErrors(err, errors.New("missing field email"))
	}
	if len(strings.TrimSpace(up.RegistrationNumber)) == 0 {
		err = mergeErrors(err, errors.New("missing field registration_number"))
	}
	if len(strings.TrimSpace(up.Password)) == 0 {
		err = mergeErrors(err, errors.New("missing field password"))
	}
	if len(strings.TrimSpace(up.UniversityID)) == 0 {
		err = mergeErrors(err, errors.New("missing field university_id"))
	}
	return
}

func (up *UserPayload) Bind(r *http.Request) error {
	if err := up.validate(); err != nil {
		return err
	}
	return nil
}

type UpdateUserPayload struct {
	Name               string `json:"name"`
	BirthDate          string `json:"birth_date"`
	Email              string `json:"email"`
	RegistrationNumber string `json:"registration_number"`
}

func (uup *UpdateUserPayload) validate() (err error) {
	if len(strings.TrimSpace(uup.Name)) == 0 {
		err = mergeErrors(err, errors.New("missing field name"))
	}
	if len(strings.TrimSpace(uup.BirthDate)) == 0 {
		err = mergeErrors(err, errors.New("missing field birth_date"))
	}
	if len(strings.TrimSpace(uup.Email)) == 0 {
		err = mergeErrors(err, errors.New("missing field email"))
	}
	if len(strings.TrimSpace(uup.RegistrationNumber)) == 0 {
		err = mergeErrors(err, errors.New("missing field registration_number"))
	}
	return
}

func (uup *UpdateUserPayload) Bind(r *http.Request) error {
	if err := uup.validate(); err != nil {
		return err
	}
	return nil
}
