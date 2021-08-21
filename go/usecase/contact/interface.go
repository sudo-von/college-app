package contact

import (
	"freelancer/college-app/go/entity"
)

type Writer interface {
	CreateContact(newContact entity.ContactPayload) error
}

type ContactRepository interface {
	Writer
}

type UseCase interface {
	CreateContact(newContact entity.ContactPayload) error
}
