package university

import (
	"freelancer/college-app/go/entity"
)

type UserReader interface {
	GetUserByID(userID string) (*entity.User, error)
}

type UserRepository interface {
	UserReader
}

type UniversityReader interface {
	GetTinyUniversities() ([]entity.TinyUniversity, *int, error)
	GetUniversityByID(universityID string) (*entity.University, error)
}

type UniversityRepository interface {
	UniversityReader
}
