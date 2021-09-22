package entity

import (
	"errors"
	"time"
)

type Advice struct {
	ID                 string
	User               TinyUser
	Classroom          Classroom
	UniversityID       string
	Subject            string
	AdviceDate         time.Time
	StudentsWillAttend []string
	Status             string
	CreationDate       time.Time
}

type AdvicePayload struct {
	ID                 string
	UserID             string
	ClassroomID        string
	UniversityID       string
	Subject            string
	AdviceDate         time.Time
	StudentsWillAttend []string
	Status             string
	CreationDate       time.Time
}

type UpdateAdvicePayload struct {
	Subject     string
	ClassroomID string
	AdviceDate  time.Time
}

type AdviceFilters struct {
	AdviceDate *time.Time
}

// ValidateDate checks if current date is less than the advice date.
func (ap *AdvicePayload) ValidateDate() error {
	validDate := false
	currentDate := time.Now().In(time.Local)
	if currentDate.Before(ap.AdviceDate) {
		validDate = true
	}
	if !validDate {
		return errors.New("advice_date can not be before the current date")
	}
	return nil
}

// ValidateDate checks if current date is less than the advice date.
func (uap *UpdateAdvicePayload) ValidateDate() error {
	validDate := false
	currentDate := time.Now().In(time.Local)
	if currentDate.Before(uap.AdviceDate) {
		validDate = true
	}
	if !validDate {
		return errors.New("advice_date can not be before the current date")
	}
	return nil
}
