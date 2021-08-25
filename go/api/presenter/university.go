package presenter

import (
	"fmt"
	"freelancer/college-app/go/entity"
	"net/http"
)

type TinyUniversityResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	ProfilePicture string `json:"profile_picture"`
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
	ID             string              `json:"id"`
	Name           string              `json:"name"`
	ProfilePicture string              `json:"profile_picture"`
	Classrooms     []ClassroomResponse `json:"classrooms"`
}

type ClassroomResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
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
