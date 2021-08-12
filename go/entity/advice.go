package entity

import "time"

type Advice struct {
	ID           string
	Subject      string
	AdviceDate   time.Time
	Classroom    int
	User         TinyUser
	University   University
	Status       string
	CreationDate time.Time
}

type AdvicePayload struct {
	ID           string
	Subject      string
	AdviceDate   time.Time
	Classroom    int
	UserID       string
	UniversityID string
	Status       string
	CreationDate time.Time
}

type AdviceFilters struct {
	InitialDate  *time.Time
	FinalDate    *time.Time
	UserID       string
	UniversityID string
}
