package user_mood

import (
	"freelancer/college-app/go/entity"
)

type UserReader interface {
	GetUserByID(userID string) (*entity.User, error)
}

type UserRepository interface {
	UserReader
}

type UserMoodReader interface {
	GetUserMoodByUserID(userID string, userMoodFilters entity.UserMoodFilters) (*entity.UserMood, error)
}

type UserMoodWriter interface {
	CreateUserMood(newUserMood entity.UserMoodPayload) error
}

type UserMoodRepository interface {
	UserMoodReader
	UserMoodWriter
}

type UseCase interface {
	GetUserMoodByUserID(userID, requestedUserID string, userMoodFilters entity.UserMoodFilters) (*entity.UserMood, error)
	CreateUserMood(requestedUserID string, newUserMood entity.UserMoodPayload, userMoodFilters entity.UserMoodFilters) error
}
