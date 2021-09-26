package university

import (
	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/entity"
)

type Service struct {
	userRepository       UserRepository
	universityRepository UniversityRepository
}

func NewService(userRepository UserRepository, universityRepository UniversityRepository) *Service {
	return &Service{
		userRepository,
		universityRepository,
	}
}

func (s Service) GetTinyUniversities() ([]entity.TinyUniversity, *int, error) {
	universities, total, err := s.universityRepository.GetTinyUniversities()
	if err != nil {
		return nil, nil, entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}
	return universities, total, nil
}

func (s Service) GetUniversityByID(userID, universityID string) (*entity.University, error) {

	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		return nil, entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	err = user.ValidateRequestedUniversity(universityID)
	if err != nil {
		return nil, entity.NewErrorConflict(err, presenter.ErrInsufficientPermissions)
	}

	university, err := s.universityRepository.GetUniversityByID(universityID)
	if err != nil {
		if err.Error() == "not found" {
			return nil, entity.NewErrorNotFound(err, presenter.ErrUniversityNotFound)
		}
		return nil, entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	return university, nil
}
