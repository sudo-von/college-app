package user_mood

import (
	"freelancer/college-app/go/entity"
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
