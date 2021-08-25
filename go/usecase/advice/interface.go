package advice

import (
	"freelancer/college-app/go/entity"
)

type Reader interface {
	GetAdvices(universityID string, adviceFilters entity.AdviceFilters) ([]entity.Advice, *int, error)
}

type AdviceRepository interface {
	Reader
}

type UseCase interface {
	GetAdvices(userID string, adviceFilters entity.AdviceFilters) ([]entity.Advice, *int, error)
}
