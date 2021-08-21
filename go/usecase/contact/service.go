package contact

import (
	"errors"
	"freelancer/college-app/go/entity"
	"regexp"
)

type Service struct {
	contactRepository ContactRepository
}

func NewService(contactRepository ContactRepository) *Service {
	return &Service{
		contactRepository,
	}
}

func (s *Service) CreateContact(newContact entity.ContactPayload) error {

	// Regex to verify if there are only numbers.
	isNumerical := regexp.MustCompile(`^[0-9]*$`)
	// Verify if the contact number is valid
	validContactNumber := false
	if isNumerical.MatchString(newContact.ContactNumber) && len(newContact.ContactNumber) == 12 {
		validContactNumber = true
	}
	if !validContactNumber {
		return errors.New("invalid contact number")
	}

	return nil
}
