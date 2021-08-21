package contact

import (
	"freelancer/college-app/go/entity"
)

type Reader interface {
	GetContactByUserID(userID string) (*entity.Contact, error)
}

type Writer interface {
	CreateContact(newContact entity.ContactPayload) error
}

type ContactRepository interface {
	Reader
	Writer
}

type UseCase interface {
	CreateContact(newContact entity.ContactPayload) error
}
