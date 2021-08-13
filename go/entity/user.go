package entity

import "time"

const (
	ActiveStatus   = "active"
	InactiveStatus = "inactive"
	BannedStatus   = "banned"
)

type TinyUser struct {
	ID                 string
	Name               string
	BirthDate          time.Time
	Email              string
	RegistrationNumber string
}

type User struct {
	ID                 string
	Name               string
	BirthDate          time.Time
	Email              string
	Password           string
	RegistrationNumber string
	Status             string
	University         University
	CreationDate       time.Time
}

type UserPayload struct {
	ID                 string
	Name               string
	BirthDate          time.Time
	Email              string
	RegistrationNumber string
	Password           string
	UniversityID       string
	Status             string
	CreationDate       time.Time
}
