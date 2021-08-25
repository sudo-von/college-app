package user

import (
	"freelancer/college-app/go/entity"
)

type Reader interface {
	GetTinyUserByID(userID string) (*entity.TinyUser, error)
	GetTinyUserByEmail(email string) (*entity.TinyUser, error)
	GetTinyUserByRegistrationNumber(registrationNumber string) (*entity.TinyUser, error)
	GetUserByID(userID string) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
}

type Writer interface {
	CreateUser(newUser entity.UserPayload) error
	UpdateUser(newUser entity.UserPayload) error
}

type UserRepository interface {
	Reader
	Writer
}

type UseCase interface {
	CreateUser(newUser entity.UserPayload) error
	GetTinyUserByID(userID, requestedUserID string) (*entity.TinyUser, error)
	UpdateTinyUser(userID, requestedUserID string, newUser entity.UpdateUserPayload) error
}
