package entity

import "time"

type Advice struct {
	ID             string
	User           TinyUser
	University     University
	Subject        string
	AdviceDate     time.Time
	Classroom      int
	StudentsNumber int
	Status         string
	CreationDate   time.Time
}

type AdvicePayload struct {
	ID             string
	UserID         string
	UniversityID   string
	Subject        string
	AdviceDate     time.Time
	Classroom      int
	StudentsNumber int
	Status         string
	CreationDate   time.Time
}

type AdviceFilters struct {
	AdviceDate *time.Time
}
