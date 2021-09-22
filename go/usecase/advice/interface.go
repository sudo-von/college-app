package advice

import (
	"freelancer/college-app/go/entity"
)

type Reader interface {
	GetAdviceByID(adviceID string) (*entity.Advice, error)
	GetAdvices(universityID string, adviceFilters entity.AdviceFilters) ([]entity.Advice, *int, error)
}

type Writer interface {
	CreateAdvice(newAdvice entity.AdvicePayload) error
	UpdateAdvice(updatedAdvice entity.AdvicePayload) error
}

type AdviceRepository interface {
	Reader
	Writer
}

type UseCase interface {
	GetAdvices(userID string, adviceFilters entity.AdviceFilters) ([]entity.Advice, *int, error)
	CreateAdvice(newAdvice entity.AdvicePayload) error
	GetAdviceByID(userID, adviceID string) (*entity.Advice, error)
	UpdateAdvice(updatedAdvice entity.UpdateAdvicePayload) error
	DeleteAdvice(userID, adviceID string) error
	UpdateAdviceStudentsNumber(userID, adviceID string) error
}
