package contact

import (
	"freelancer/college-app/go/entity"
)

type Reader interface {
	GetContactByID(contactID string) (*entity.Contact, error)
	GetContactByUserID(userID string) (*entity.Contact, error)
}

type Writer interface {
	CreateContact(newContact entity.ContactPayload) error
	UpdateContactByUserID(newContact entity.ContactPayload) error
}

type ContactRepository interface {
	Reader
	Writer
}

type UseCase interface {
	GetContactByID(contactID string) (*entity.Contact, error)
	GetContactByUserID(userID, requestedUserID string) (*entity.Contact, error)
	CreateContact(newContact entity.ContactPayload) error
	UpdateContactByUserID(userID string, newContact entity.UpdateContactPayload) error
}
