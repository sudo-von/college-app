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
	// Verifies if the contact number is valid
	validContactNumber := false
	if isNumerical.MatchString(newContact.ContactNumber) && len(newContact.ContactNumber) == 12 {
		validContactNumber = true
	}
	if !validContactNumber {
		return errors.New("invalid contact number")
	}
	// Verifies if the given user_id has registered his contact yet.
	_, err := s.contactRepository.GetContactByUserID(newContact.UserID)
	if err == nil {
		return errors.New("given user_id already has a contact registered")
	}
	// Stores new contact.
	err = s.contactRepository.CreateContact(newContact)
	if err != nil {
		return err
	}
	return nil
}
