package university

import (
	"freelancer/college-app/go/entity"
)

type Reader interface {
	GetTinyUniversities() ([]entity.TinyUniversity, *int, error)
	GetUniversityByID(universityID string) (*entity.University, error)
}

type UniversityRepository interface {
	Reader
}

type UseCase interface {
	GetTinyUniversities() ([]entity.TinyUniversity, *int, error)
	GetUniversityByID(universityID string) (*entity.University, error)
}
