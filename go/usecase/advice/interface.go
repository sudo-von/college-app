package advice

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
	GetUniversityByID(universityID string) (*entity.University, error)
}

type UniversityRepository interface {
	UniversityReader
}

type AdviceReader interface {
	GetAdviceByID(adviceID string) (*entity.Advice, error)
	GetAdvices(universityID string, adviceFilters entity.AdviceFilters) ([]entity.Advice, *int, error)
}

type AdviceWriter interface {
	CreateAdvice(newAdvice entity.AdvicePayload) error
	UpdateAdvice(updatedAdvice entity.AdvicePayload) error
}

type AdviceRepository interface {
	AdviceReader
	AdviceWriter
}
