package suggestion

import (
	"freelancer/college-app/go/entity"
)

type SuggestionWriterMock struct{}

func (s SuggestionWriterMock) CreateSuggestion(newSuggestion entity.SuggestionPayload) error {
	return nil
}

type SuggestionRepositoryMock struct {
	SuggestionWriterMock
}
