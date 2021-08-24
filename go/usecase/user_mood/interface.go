package user_mood

import (
	"freelancer/college-app/go/entity"
)

type Reader interface {
	GetUserMoodByUserID(userID string, userMoodFilters entity.UserMoodFilters) (*entity.UserMood, error)
}

type Writer interface {
	CreateUserMood(newUserMood entity.UserMoodPayload) error
}

type UserMoodRepository interface {
	Reader
	Writer
}

type UseCase interface {
	GetUserMoodByUserID(userID string, userMoodFilters entity.UserMoodFilters) (*entity.UserMood, error)
	CreateUserMood(newUserMood entity.UserMoodPayload) error
}
