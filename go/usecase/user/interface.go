package user

import (
	"freelancer/college-app/go/entity"
)

type Reader interface {
	GetTinyUserByID(userID string) (*entity.TinyUser, error)
	GetTinyUserByEmail(email string) (*entity.TinyUser, error)
	GetTinyUserByRegistrationNumber(registrationNumber string) (*entity.TinyUser, error)
	GetUserByEmail(email string) (*entity.User, error)
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
