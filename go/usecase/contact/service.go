package contact

import (
	"errors"
	"fmt"
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

func (s *Service) GetContactByID(userID, contactID string) (*entity.Contact, error) {
	contact, err := s.contactRepository.GetContactByID(contactID)
	if err != nil {
		if err.Error() == "not found" {
			return nil, entity.NewErrorNotFound(fmt.Errorf("GetContactByID: %w", errors.New("contact not found")))
		}
		return nil, entity.NewErrorInternalServer(fmt.Errorf("GetContactByID: %w", err))
	}
	return contact, nil
}

func (s *Service) GetContactByUserID(userID, requestedUserID string) (*entity.Contact, error) {
	// Checks permissions.
	hasPermission := false
	if userID == requestedUserID {
		hasPermission = true
	}
	if !hasPermission {
		return nil, entity.NewErrorUnauthorized(errors.New("user has no permission to see this user"))
	}
	contact, err := s.contactRepository.GetContactByUserID(requestedUserID)
	if err != nil {
		if err.Error() == "not found" {
			return nil, entity.NewErrorNotFound(fmt.Errorf("GetContactByUserID: %w", errors.New("contact not found")))
		}
		return nil, entity.NewErrorInternalServer(fmt.Errorf("GetContactByUserID: %w", err))
	}
	return contact, nil
}

func (s *Service) CreateContact(userID, requestedUserID string, newContact entity.ContactPayload) error {

	// Checks permissions.
	hasPermission := false
	if userID == requestedUserID {
		hasPermission = true
	}
	if !hasPermission {
		return entity.NewErrorUnauthorized(errors.New("user has no permission to see this user"))
	}

	// Regex to verify if there are only numbers.
	isNumerical := regexp.MustCompile(`^[0-9]*$`)
	// Verifies if the contact number is valid.
	validContactNumber := false
	if isNumerical.MatchString(newContact.ContactNumber) && len(newContact.ContactNumber) == 12 {
		validContactNumber = true
	}
	if !validContactNumber {
		return entity.NewErrorConflict(errors.New("invalid contact number"))
	}
	// Verifies if the given user has registered his contact yet.
	_, err := s.GetContactByUserID(userID, requestedUserID)
	if err == nil {
		return entity.NewErrorConflict(fmt.Errorf("GetContactByUserID: %w", errors.New("given user_id already has a contact registered")))
	}
	// Stores new contact.
	err = s.contactRepository.CreateContact(newContact)
	if err != nil {
		return entity.NewErrorInternalServer(fmt.Errorf("CreateContact: %w", err))
	}
	return nil
}

func (s *Service) UpdateContactByID(userID, contactID string, newContact entity.UpdateContactPayload) error {

	// Gets old contact.
	oldContact, err := s.GetContactByID(userID, contactID)
	if err != nil {
		return err
	}

	// If both phone numbers are different, checks if the new contact number is valid.
	if oldContact.ContactNumber != newContact.ContactNumber {
		// Regex to verify if there are only numbers.
		isNumerical := regexp.MustCompile(`^[0-9]*$`)
		// Verifies if the contact number is valid.
		validContactNumber := false
		if isNumerical.MatchString(newContact.ContactNumber) && len(newContact.ContactNumber) == 12 {
			validContactNumber = true
		}
		if !validContactNumber {
			return entity.NewErrorConflict(errors.New("invalid contact number"))
		}
	}
	// Creates a new user payload and then replaces the old one.
	updatedContact := entity.ContactPayload{
		ID:            oldContact.ID,
		UserID:        oldContact.UserID,
		ContactName:   newContact.ContactName,
		ContactNumber: newContact.ContactNumber,
		Message:       newContact.Message,
		CreationDate:  oldContact.CreationDate,
	}
	err = s.contactRepository.UpdateContactByUserID(updatedContact)
	if err != nil {
		return entity.NewErrorInternalServer(fmt.Errorf("UpdateContactByUserID: %w", err))
	}

	return nil
}
