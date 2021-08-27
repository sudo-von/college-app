package user_mood

import (
	"errors"
	"fmt"
	"freelancer/college-app/go/entity"
	"time"
)

type Service struct {
	userMoodRepository UserMoodRepository
}

func NewService(userMoodRepository UserMoodRepository) *Service {
	return &Service{
		userMoodRepository,
	}
}

func (s *Service) GetUserMoodByUserID(userID, requestedUserID string, userMoodFilters entity.UserMoodFilters) (*entity.UserMood, error) {

	// Checks permissions.
	hasPermission := false
	if userID == requestedUserID {
		hasPermission = true
	}
	if !hasPermission {
		return nil, entity.NewErrorUnauthorized(errors.New("user has no permission to see this user"))
	}

	userMood, err := s.userMoodRepository.GetUserMoodByUserID(requestedUserID, userMoodFilters)
	if err != nil {
		if err.Error() == "not found" {
			return nil, entity.NewErrorNotFound(fmt.Errorf("GetUserMoodByUserID: %w", errors.New("mood not found")))
		}
		return nil, entity.NewErrorInternalServer(fmt.Errorf("GetUserMoodByUserID: %w", err))
	}

	return userMood, nil
}

func (s *Service) CreateUserMood(newUserMood entity.UserMoodPayload) error {

	// Checks if mood is a valid number in the specified range
	validMood := false
	if newUserMood.Mood >= 1 && newUserMood.Mood <= 5 {
		validMood = true
	}
	if !validMood {
		return entity.NewErrorConflict(errors.New("invalid mood, not in range from 1 to 5"))
	}
	// Checks if the user mood has already been registered today.
	currentDate := time.Now().In(time.Local)
	userMoodFilters := entity.UserMoodFilters{
		CreationDate: &currentDate,
	}
	_, err := s.GetUserMoodByUserID(newUserMood.UserID, newUserMood.UserID, userMoodFilters)
	if err == nil {
		return entity.NewErrorConflict(fmt.Errorf("GetUserMoodByUserID: %w", errors.New("user mood has already been registered today")))
	}
	// Stores new user mood.
	err = s.userMoodRepository.CreateUserMood(newUserMood)
	if err != nil {
		return entity.NewErrorInternalServer(fmt.Errorf("CreateUserMood: %w", err))
	}
	return nil
}
