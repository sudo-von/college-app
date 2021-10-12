package handler

import "freelancer/college-app/go/entity"

type UserMoodServiceMock struct{}

func (s UserMoodServiceMock) GetUserMoodByUserID(userID, requestedUserID string, userMoodFilters entity.UserMoodFilters) (*entity.UserMood, error) {
	return &entity.UserMood{}, nil
}

func (s UserMoodServiceMock) CreateUserMood(userID, requestedUserID string, newUserMood entity.UserMoodPayload) error {
	return nil
}
