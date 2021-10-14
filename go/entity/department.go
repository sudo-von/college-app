package entity

import (
	"errors"
	"time"
)

type Department struct {
	ID           string
	User         BasicUser
	UniversityID string
	Description  string
	Street       string
	Neighborhood string
	Status       string
	Cost         float32
	Available    bool
	CreationDate time.Time
}

func (d *Department) ValidateRequestedDepartment(userID string) error {
	validRequestedDepartment := false
	if userID == d.User.ID {
		validRequestedDepartment = true
	}
	if !validRequestedDepartment {
		return errors.New("user has not authorization to request another user's department")
	}
	return nil
}

func (d *Department) ValidateUniversity(universityID string) error {
	validRequestedUniversity := false
	if d.UniversityID == universityID {
		validRequestedUniversity = true
	}
	if !validRequestedUniversity {
		return errors.New("user has not authorization to request another university resource")
	}
	return nil
}

type DepartmentPayload struct {
	ID           string
	UserID       string
	UniversityID string
	Description  string
	Street       string
	Neighborhood string
	Cost         float32
	Status       string
	Available    bool
	CreationDate time.Time
}

func (dp *DepartmentPayload) ValidateCost() error {
	validCost := false
	if dp.Cost >= 1 {
		validCost = true
	}
	if !validCost {
		return errors.New("cost can't be lower than 1")
	}
	return nil
}

type UpdateDepartmentPayload struct {
	Description  string
	Street       string
	Neighborhood string
	Cost         float32
	Available    bool
}

func (udp *UpdateDepartmentPayload) ValidateCost() error {
	validCost := false
	if udp.Cost >= 1 {
		validCost = true
	}
	if !validCost {
		return errors.New("cost can't be lower than 1")
	}
	return nil
}

type DepartmentFilters struct {
	Cost   *float32
	UserID string
}
