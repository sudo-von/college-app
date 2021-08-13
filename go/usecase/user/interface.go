package user

import (
	"freelancer/college-app/go/entity"
)

type Reader interface {
	GetUserByID(userID string) (*entity.TinyUser, error)
	GetUserByEmail(email string) (*entity.TinyUser, error)
	GetUserByRegistrationNumber(registrationNumber string) (*entity.TinyUser, error)
}

type Writer interface {
	CreateUser(newUser entity.UserPayload) error
}

type UserRepository interface {
	Reader
	Writer
}

type UseCase interface {
	CreateUser(newUser entity.UserPayload) error
}
