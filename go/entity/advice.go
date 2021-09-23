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

func (a *Advice) ValidateRequestedAdvice(userID string) error {
	validRequestedAdvice := false
	if userID == a.User.ID {
		validRequestedAdvice = true
	}
	if !validRequestedAdvice {
		return errors.New("user has not authorization to request another user's advice")
	}
	return nil
}

// ValidateUniversity checks if the requested advice belongs to the same university as the user's.
func (a *Advice) ValidateUniversity(universityID string) error {
	validRequestedAdvice := false
	if a.UniversityID == universityID {
		validRequestedAdvice = true
	}
	if !validRequestedAdvice {
		return errors.New("user has not authorization to request another university advice resource")
	}
	return nil
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
