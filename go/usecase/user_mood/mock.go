package user_mood

import (
	"errors"
	"freelancer/college-app/go/entity"
	"time"
)

func NewFakeUserMood() entity.UserMood {
	return entity.UserMood{
		ID:           "615cb8b7cc577570b57b48ea",
		UserID:       "615c09f7309d7ded48c7a053",
		Mood:         5,
		CreationDate: time.Date(2021, 01, 01, 0, 0, 0, 0, time.Local),
	}
}

type UserMoodReaderMock struct{}

func (u UserMoodReaderMock) GetUserMoodByUserID(userID string, userMoodFilters entity.UserMoodFilters) (*entity.UserMood, error) {
	if userID != "615c09f7309d7ded48c7a053" {
		return nil, errors.New("not found")
	}
	userMood := NewFakeUserMood()
	return &userMood, nil
}

type UserMoodWriterMock struct{}

func (u UserMoodWriterMock) CreateUserMood(newUserMood entity.UserMoodPayload) error {
	return nil
}

type UserMoodRepositoryMock struct {
	UserMoodReaderMock
	UserMoodWriterMock
}
