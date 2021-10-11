package contact

import (
	"freelancer/college-app/go/entity"
)

type UserReader interface {
	GetUserByID(userID string) (*entity.User, error)
}

type UserRepository interface {
	UserReader
}

type ContactReader interface {
	GetContactByID(contactID string) (*entity.Contact, error)
	GetContactByUserID(userID string) (*entity.Contact, error)
}

type ContactWriter interface {
	CreateContact(newContact entity.ContactPayload) error
	UpdateContactByUserID(newContact entity.ContactPayload) error
}

type ContactRepository interface {
	ContactReader
	ContactWriter
}
