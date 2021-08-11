package entity

import "time"

const (
	ActiveStatus   = "active"
	InactiveStatus = "inactive"
	BannedStatus   = "banned"
)

type TinyUser struct {
	ID                 string
	RegistrationNumber string
	Name               string
	Email              string
}

type User struct {
	ID                 string
	RegistrationNumber string
	Name               string
	Email              string
	Password           string
	University         University
	Status             string
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
