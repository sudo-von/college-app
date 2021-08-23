package user_mood

import (
	"freelancer/college-app/go/entity"
)

type Reader interface {
	GetUserMoodByUserID(userID string, userMoodFilters entity.UserMoodFilters) (*entity.UserMood, error)
}

type UserMoodRepository interface {
	Reader
}

type UseCase interface {
	GetUserMoodByUserID(userID string, userMoodFilters entity.UserMoodFilters) (*entity.UserMood, error)
}
