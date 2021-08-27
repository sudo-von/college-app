package entity

import "time"

type Advice struct {
	ID             string
	User           TinyUser
	Classroom      Classroom
	Subject        string
	AdviceDate     time.Time
	StudentsNumber int
	Status         string
	CreationDate   time.Time
}

type AdvicePayload struct {
	ID             string
	UserID         string
	ClassroomID    string
	UniversityID   string
	Subject        string
	AdviceDate     time.Time
	StudentsNumber []string
	Status         string
	CreationDate   time.Time
}

type AdviceFilters struct {
	AdviceDate *time.Time
}
