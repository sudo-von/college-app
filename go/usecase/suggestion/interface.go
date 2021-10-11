package suggestion

import (
	"freelancer/college-app/go/entity"
)

type SuggestionWriter interface {
	CreateSuggestion(newSuggestion entity.SuggestionPayload) error
}

type SuggestionRepository interface {
	SuggestionWriter
}
