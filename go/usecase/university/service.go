package university

import (
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

func (s Service) GetUniversityByID(universityID string) (*entity.University, error) {
	university, err := s.universityRepository.GetUniversityByID(universityID)
	if err != nil {
		return nil, err
	}
	return university, nil
}
