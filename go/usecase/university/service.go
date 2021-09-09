package university

import (
	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/entity"
)

type Service struct {
	universityRepository UniversityRepository
}

func NewService(universityRepository UniversityRepository) *Service {
	return &Service{
		universityRepository: universityRepository,
	}
}

func (s Service) GetTinyUniversities() ([]entity.TinyUniversity, *int, error) {
	universities, total, err := s.universityRepository.GetTinyUniversities()
	if err != nil {
		return nil, nil, entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}
	return universities, total, nil
}

func (s Service) GetUniversityByID(universityID string) (*entity.University, error) {
	university, err := s.universityRepository.GetUniversityByID(universityID)
	if err != nil {
		if err.Error() == "not found" {
			return nil, entity.NewErrorNotFound(err, presenter.ErrUniversityNotFound)
		}
		return nil, entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}
	return university, nil
}
