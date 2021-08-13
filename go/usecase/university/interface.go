package university

import (
	"freelancer/college-app/go/entity"
)

type Reader interface {
	GetUniversityByID(universityID string) (*entity.University, error)
}

type UniversityRepository interface {
	Reader
}

type UseCase interface {
	GetUniversityByID(universityID string) (*entity.University, error)
}
