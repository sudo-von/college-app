package entity

import (
	"errors"
	"strings"
	"time"
)

var (
	StudentRole = "student"
	AdminRole   = "admin"
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
	Role               string
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
	Role               string
	CreationDate       time.Time
}

type UpdateUserPayload struct {
	Name               string
	BirthDate          time.Time
	Email              string
	RegistrationNumber string
}

func (u *User) ValidateRequestedUser(requestedUserID string) error {
	validRequestedUser := false
	if requestedUserID == u.ID {
		validRequestedUser = true
	}
	if !validRequestedUser {
		return errors.New("user has not authorization to request another user's resource")
	}
	return nil
}

func (u *User) ValidateRequestedUniversity(requestedUniversityID string) error {
	validRequestedUniversity := false
	if requestedUniversityID == u.ID {
		validRequestedUniversity = true
	}
	if !validRequestedUniversity {
		return errors.New("user has not authorization to request another university resource")
	}
	return nil
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
