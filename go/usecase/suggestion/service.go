package suggestion

import (
	"freelancer/college-app/go/entity"
)

type Service struct {
	suggestionRepository SuggestionRepository
}

func NewService(suggestionRepository SuggestionRepository) *Service {
	return &Service{
		suggestionRepository,
	}
}

func (s *Service) CreateSuggestion(newSuggestion entity.SuggestionPayload) error {
	err := s.suggestionRepository.CreateSuggestion(newSuggestion)
	if err != nil {
		return err
	}
	return nil
}
