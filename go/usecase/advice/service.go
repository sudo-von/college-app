package advice

import (
	"errors"
	"fmt"
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
		if err.Error() == "not found" {
			return nil, nil, entity.NewErrorNotFound(fmt.Errorf("GetUserByID: %w", errors.New("user not found")))
		}
		return nil, nil, entity.NewErrorInternalServer(fmt.Errorf("GetUserByID: %w", err))
	}

	advices, total, err := s.adviceRepository.GetAdvices(user.University.ID, adviceFilters)
	if err != nil {
		return nil, nil, entity.NewErrorInternalServer(fmt.Errorf("GetAdvices: %w", err))
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
		return entity.NewErrorConflict(errors.New("advice_date can not be before the current date"))
	}

	// Gets university id from the user.
	user, err := s.userRepository.GetUserByID(newAdvice.UserID)
	if err != nil {
		if err.Error() == "not found" {
			return entity.NewErrorNotFound(fmt.Errorf("GetUserByID: %w", errors.New("user not found")))
		}
		return entity.NewErrorInternalServer(fmt.Errorf("GetUserByID: %w", err))
	}
	newAdvice.UniversityID = user.University.ID

	// Checks if given classroom is a valid classroom from that university.
	validClassroom := false
	university, err := s.universityRepository.GetUniversityByID(user.University.ID)
	if err != nil {
		if err.Error() == "not found" {
			return entity.NewErrorNotFound(fmt.Errorf("GetUniversityByID: %w", errors.New("university not found")))
		}
		return entity.NewErrorInternalServer(fmt.Errorf("GetUniversityByID: %w", err))
	}
	for _, c := range university.Classrooms {
		if c.ID == newAdvice.ClassroomID {
			validClassroom = true
		}
	}
	if !validClassroom {
		return entity.NewErrorConflict(errors.New("invalid classroom_id"))
	}

	// Stores new advice.
	err = s.adviceRepository.CreateAdvice(newAdvice)
	if err != nil {
		return entity.NewErrorInternalServer(fmt.Errorf("CreateAdvice: %w", err))
	}
	return nil
}
