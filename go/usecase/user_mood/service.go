package user_mood

import (
	"errors"
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

func (s *Service) GetUserMoodByUserID(userID string, userMoodFilters entity.UserMoodFilters) (*entity.UserMood, error) {

	userMood, err := s.userMoodRepository.GetUserMoodByUserID(userID, userMoodFilters)
	if err != nil {
		return nil, err
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
		return errors.New("invalid mood, not in range from 1 to 5")
	}
	// Checks if the user mood has already been registered today.
	currentDate := time.Now().In(time.Local)
	userMoodFilters := entity.UserMoodFilters{
		CreationDate: &currentDate,
	}
	_, err := s.GetUserMoodByUserID(newUserMood.UserID, userMoodFilters)
	if err == nil {
		return errors.New("user mood has already been registered today")
	}

	// Stores new user mood.
	err = s.userMoodRepository.CreateUserMood(newUserMood)
	if err != nil {
		return err
	}
	return nil
}
