package presenter

import (
	"fmt"
	"freelancer/college-app/go/entity"
	"net/http"
)

var (
	ErrUniversityNotFound = "UNIVERSITY_NOT_FOUND"
	ErrInvClassroom       = "INVALID_CLASSROOM"
)

type TinyUniversityList struct {
	Total            int                      `json:"total" example:"10"`
	TinyUniversities []TinyUniversityResponse `json:"results"`
}

func (tul *TinyUniversityList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type TinyUniversityResponse struct {
	ID             string `json:"id" example:"613aab578a6ef50007e622be"`
	Name           string `json:"name" example:"VoN"`
	ProfilePicture string `json:"profile_picture" example:"Url"`
}

func (ur *TinyUniversityResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ToTinyUniversityPresenter(university entity.TinyUniversity) TinyUniversityResponse {
	return TinyUniversityResponse{
		ID:             university.ID,
		Name:           university.Name,
		ProfilePicture: university.ProfilePicture,
	}
}

type UniversityResponse struct {
	ID             string              `json:"id" example:"613aab578a6ef50007e622be"`
	Name           string              `json:"name" example:"VoN"`
	ProfilePicture string              `json:"profile_picture" example:"Url"`
	Classrooms     []ClassroomResponse `json:"classrooms"`
}

type ClassroomResponse struct {
	ID   string `json:"id" example:"613aab578a6ef50007e622be"`
	Name string `json:"name" example:"VoN"`
}

func (ur *UniversityResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ToUniversityPresenter(university entity.University) UniversityResponse {

	classrooms := make([]ClassroomResponse, 0)
	for _, c := range university.Classrooms {
		classrooms = append(classrooms, ClassroomResponse{
			ID:   c.ID,
			Name: c.Name,
		})
	}

	fmt.Println(classrooms)
	return UniversityResponse{
		ID:             university.ID,
		Name:           university.Name,
		ProfilePicture: university.ProfilePicture,
		Classrooms:     classrooms,
	}
}
