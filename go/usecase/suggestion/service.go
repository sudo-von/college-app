package suggestion

import (
	"freelancer/college-app/go/api/presenter"
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
		if err.Error() == "not found" {
			return entity.NewErrorNotFound(err, presenter.ErrSuggestionNotFound)
		}
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}
	return nil
}
