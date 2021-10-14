package presenter

import (
	"freelancer/college-app/go/entity"
	"net/http"
)

type DepartmentList struct {
	Total       int                  `json:"total" example:"1"`
	Departments []DepartmentResponse `json:"results"`
}

func (dl *DepartmentList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type DepartmentResponse struct {
	ID           string    `json:"id" example:"613aab578a6ef50007e622be"`
	User         BasicUser `json:"user"`
	Description  string    `json:"description" example:"description"`
	Street       string    `json:"street" example:"street"`
	Neighborhood string    `json:"neighborhood" example:"neighborhood"`
	Cost         float32   `json:"cost" example:"1000"`
	Available    bool      `json:"available" example:"true"`
}

func (dr *DepartmentResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ToDepartmentPresenter(department entity.Department) DepartmentResponse {

	user := BasicUser{
		ID:    department.User.ID,
		Name:  department.User.Name,
		Email: department.User.Email,
	}

	return DepartmentResponse{
		ID:           department.ID,
		User:         user,
		Description:  department.Description,
		Street:       department.Street,
		Neighborhood: department.Neighborhood,
		Cost:         department.Cost,
		Available:    department.Available,
	}
}
