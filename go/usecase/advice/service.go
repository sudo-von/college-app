package advice

import (
	"errors"
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/usecase/university"
	"freelancer/college-app/go/usecase/user"
	"time"
)

type Service struct {
	adviceRepository     AdviceRepository
	userRepository       user.UserRepository
	universityRepository university.UniversityRepository
}

func NewService(adviceRepository AdviceRepository, userRepository user.UserRepository, universityRepository university.UniversityRepository) *Service {
	return &Service{
		adviceRepository,
		userRepository,
		universityRepository,
	}
}

func (s *Service) GetAdvices(userID string, adviceFilters entity.AdviceFilters) ([]entity.Advice, *int, error) {

	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		return nil, nil, err
	}

	advices, total, err := s.adviceRepository.GetAdvices(user.University.ID, adviceFilters)
	if err != nil {
		return nil, nil, err
	}

	return advices, total, nil
}

func (s *Service) CreateAdvice(newAdvice entity.AdvicePayload) error {

	// Checks if current date is before than the advice date.
	validDate := false
	currentDate := time.Now().In(time.Local)
	if currentDate.Before(newAdvice.AdviceDate) {
		validDate = true
	}
	if !validDate {
		return errors.New("advice_date can not be before the current date")
	}

	// Gets university id from the user.
	user, err := s.userRepository.GetUserByID(newAdvice.UserID)
	if err != nil {
		return err
	}
	newAdvice.UniversityID = user.University.ID

	// Checks if classroom is a valid classroom for the given university.

	// Stores new advice.
	err = s.adviceRepository.CreateAdvice(newAdvice)
	if err != nil {
		return err
	}
	return nil
}
