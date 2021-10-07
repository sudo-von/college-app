package user_mood

import (
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/pkg/search"
	"time"
)

var (
	fakeUserIDS = []string{"615c09f7309d7ded48c7a053", "615c09f7309d7ded48c7a055"}
)

type UserMoodReaderMock struct{}

func (u UserMoodReaderMock) GetUserMoodByUserID(userID string, userMoodFilters entity.UserMoodFilters) (*entity.UserMood, error) {

	err := search.FindInSlice(userID, fakeUserIDS)
	if err != nil {
		return nil, err
	}
	userMood := &entity.UserMood{
		ID:           "615cb8b7cc577570b57b48ea",
		UserID:       userID,
		Mood:         5,
		CreationDate: time.Date(2021, 01, 01, 0, 0, 0, 0, time.Local),
	}
	return userMood, nil
}

type UserMoodWriterMock struct{}

func (u UserMoodWriterMock) CreateUserMood(newUserMood entity.UserMoodPayload) error {
	return nil
}

type UserMoodRepositoryMock struct {
	UserMoodReaderMock
	UserMoodWriterMock
}
