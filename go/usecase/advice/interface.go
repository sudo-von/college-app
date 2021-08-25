package advice

import (
	"freelancer/college-app/go/entity"
)

type Reader interface {
	GetAdvices(universityID string, adviceFilters entity.AdviceFilters) ([]entity.Advice, *int, error)
}

type Writer interface {
	CreateAdvice(newAdvice entity.AdvicePayload) error
}

type AdviceRepository interface {
	Reader
	Writer
}

type UseCase interface {
	GetAdvices(userID string, adviceFilters entity.AdviceFilters) ([]entity.Advice, *int, error)
	CreateAdvice(newAdvice entity.AdvicePayload) error
}
