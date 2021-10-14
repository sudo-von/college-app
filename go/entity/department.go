package entity

import (
	"time"
)

type Department struct {
	ID           string
	User         BasicUser
	UniversityID string
	Description  string
	Street       string
	Neighborhood string
	Cost         float32
	Available    bool
	CreationDate time.Time
}

type DepartmentPayload struct {
	ID           string
	UserID       string
	UniversityID string
	Description  string
	Street       string
	Neighborhood string
	Cost         float32
	Available    bool
	CreationDate time.Time
}

type DepartmentFilters struct {
	Cost   *float32
	UserID string
}
