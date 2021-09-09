package entity

import (
	"errors"
	"strings"
	"time"
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

type UpdateUserPayload struct {
	Name               string
	BirthDate          time.Time
	Email              string
	RegistrationNumber string
}

func (up *UserPayload) ValidateRegistrationNumber() error {
	validRegistrationNumber := false
	if len(strings.Replace(up.RegistrationNumber, " ", "", -1)) == 8 {
		validRegistrationNumber = true
	}
	if !validRegistrationNumber {
		return errors.New("invalid registration_number")
	}
	return nil
}

func (uup *UpdateUserPayload) ValidateRegistrationNumber() error {
	validRegistrationNumber := false
	if len(strings.Replace(uup.RegistrationNumber, " ", "", -1)) == 8 {
		validRegistrationNumber = true
	}
	if !validRegistrationNumber {
		return errors.New("invalid registration_number")
	}
	return nil
}
