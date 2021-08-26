package university

import (
	"errors"
	"fmt"
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
		return nil, nil, entity.NewErrorInternalServer(fmt.Errorf("GetTinyUniversities: %w", err))
	}
	return universities, total, nil
}

func (s Service) GetUniversityByID(universityID string) (*entity.University, error) {
	university, err := s.universityRepository.GetUniversityByID(universityID)
	if err != nil {
		if err.Error() == "not found" {
			return nil, entity.NewErrorNotFound(fmt.Errorf("GetUniversityByID: %w", errors.New("university not found")))
		}
		return nil, entity.NewErrorInternalServer(fmt.Errorf("GetUniversityByID: %w", err))
	}
	return university, nil
}
