package entity

import (
	"errors"
	"regexp"
	"time"
)

type Contact struct {
	ID            string
	UserID        string
	ContactName   string
	ContactNumber string
	Message       string
	CreationDate  time.Time
}

type ContactPayload struct {
	ID            string
	UserID        string
	ContactName   string
	ContactNumber string
	Message       string
	CreationDate  time.Time
}

type UpdateContactPayload struct {
	ContactName   string
	ContactNumber string
	Message       string
}

func (cp *ContactPayload) ValidateNumber() error {
	validContactNumber := false
	// Verifies if there are only numbers.
	isNumerical := regexp.MustCompile(`^[0-9]*$`)
	if isNumerical.MatchString(cp.ContactNumber) && len(cp.ContactNumber) == 12 {
		validContactNumber = true
	}
	if !validContactNumber {
		return errors.New("invalid contact number")
	}
	return nil
}

func (ucp *UpdateContactPayload) ValidateNumber() error {
	validContactNumber := false
	// Verifies if there are only numbers.
	isNumerical := regexp.MustCompile(`^[0-9]*$`)
	if isNumerical.MatchString(ucp.ContactNumber) && len(ucp.ContactNumber) == 12 {
		validContactNumber = true
	}
	if !validContactNumber {
		return errors.New("invalid contact number")
	}
	return nil
}
