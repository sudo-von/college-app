package handler

import (
	"errors"
	"net/http"
	"time"

	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/entity"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type SuggestionService interface {
	CreateSuggestion(newSuggestion entity.SuggestionPayload) error
}

type SuggestionController struct {
	SuggestionService SuggestionService
	AuthService       func(http.Handler) http.Handler
}

func NewSuggestionController(suggestionService SuggestionService, authService func(http.Handler) http.Handler) *SuggestionController {
	return &SuggestionController{
		SuggestionService: suggestionService,
		AuthService:       authService,
	}
}

func (c *SuggestionController) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(c.AuthService)
	r.Post("/", c.Create)
	return r
}

// @tags suggestions
// @summary Create suggestion.
// @description Create suggestion.
// @security BearerJWT
// @param payload body presenter.SuggestionPayload true "Suggestion that wants to be stored."
// @id create-suggestion
// @produce json
// @success 201
// @router /suggestions [post]
func (c *SuggestionController) Create(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(ContextKeyUserID).(string)
	if !ok {
		err := errors.New("user not in context")
		CheckError(err, w, r)
		return
	}

	var data presenter.SuggestionPayload
	if err := render.Bind(r, &data); err != nil {
		CheckError(err, w, r)
		return
	}

	newSuggestion := entity.SuggestionPayload{
		UserID:       userID,
		Suggestion:   data.Suggestion,
		CreationDate: time.Now().In(time.Local),
	}

	err := c.SuggestionService.CreateSuggestion(newSuggestion)
	if err != nil {
		CheckError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
