package contact

import (
	"errors"
	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/entity"
)

type Service struct {
	userRepository    UserRepository
	contactRepository ContactRepository
}

func NewService(userRepository UserRepository, contactRepository ContactRepository) *Service {
	return &Service{
		userRepository,
		contactRepository,
	}
}

func (s Service) GetContactByUserID(userID, requestedUserID string) (*entity.Contact, error) {

	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		return nil, entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	err = user.ValidateRequestedUser(requestedUserID)
	if err != nil {
		return nil, entity.NewErrorUnauthorized(err, presenter.ErrInsufficientPermissions)
	}

	contact, err := s.contactRepository.GetContactByUserID(requestedUserID)
	if err != nil {
		if err.Error() == "not found" {
			return nil, entity.NewErrorNotFound(err, presenter.ErrContactNotFound)
		}
		return nil, entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	return contact, nil
}

func (s Service) CreateContact(userID, requestedUserID string, newContact entity.ContactPayload) error {

	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	err = user.ValidateRequestedUser(requestedUserID)
	if err != nil {
		return entity.NewErrorUnauthorized(err, presenter.ErrInsufficientPermissions)
	}

	// Validates contact number.
	err = newContact.ValidateNumber()
	if err != nil {
		return entity.NewErrorConflict(err, presenter.ErrInvContactNumber)
	}

	// Verifies if the given user has registered his contact yet.
	_, err = s.GetContactByUserID(userID, requestedUserID)
	if err != nil && err.Error() != "not found" {
		return err
	} else if err == nil {
		return entity.NewErrorConflict(errors.New("given user_id already has a contact registered"), presenter.ErrContactAlreadyRegistered)
	}

	// Stores new contact.
	err = s.contactRepository.CreateContact(newContact)
	if err != nil {
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}
	return nil
}

func (s Service) UpdateContactByID(userID, contactID string, newContact entity.UpdateContactPayload) error {

	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	// Gets old contact.
	oldContact, err := s.contactRepository.GetContactByID(contactID)
	if err != nil {
		if err.Error() == "not found" {
			return entity.NewErrorNotFound(err, presenter.ErrContactNotFound)
		}
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	err = user.ValidateRequestedUser(oldContact.UserID)
	if err != nil {
		return entity.NewErrorUnauthorized(err, presenter.ErrInsufficientPermissions)
	}

	// If both contact numbers are different, checks if the new contact number is valid.
	if oldContact.ContactNumber != newContact.ContactNumber {
		err := newContact.ValidateNumber()
		if err != nil {
			return entity.NewErrorConflict(err, presenter.ErrInvContactNumber)
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
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}
	return nil
}
