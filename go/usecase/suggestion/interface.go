package suggestion

import (
	"freelancer/college-app/go/entity"
)

type Writer interface {
	CreateSuggestion(newSuggestion entity.SuggestionPayload) error
}

type SuggestionRepository interface {
	Writer
}

type UseCase interface {
	CreateSuggestion(newSuggestion entity.SuggestionPayload) error
}
