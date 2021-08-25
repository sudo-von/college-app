package advice

import (
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/usecase/user"
)

type Service struct {
	adviceRepository AdviceRepository
	userRepository   user.UserRepository
}

func NewService(adviceRepository AdviceRepository, userRepository user.UserRepository) *Service {
	return &Service{
		adviceRepository,
		userRepository,
	}
}

func (s *Service) GetAdvices(userID string, adviceFilters entity.AdviceFilters) ([]entity.Advice, *int, error) {

	user, err := s.userRepository.GetTinyUserByID(userID)
	if err != nil {
		return nil, nil, err
	}

	advices, total, err := s.adviceRepository.GetAdvices(user.University.ID, adviceFilters)
	if err != nil {
		return nil, nil, err
	}

	return advices, total, nil
}
