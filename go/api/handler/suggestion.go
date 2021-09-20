package handler

import (
	"errors"
	"net/http"
	"time"

	"freelancer/college-app/go/api/middleware"
	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/pkg/token"
	"freelancer/college-app/go/usecase/suggestion"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type SuggestionController struct {
	SuggestionService suggestion.Service
	TokenService      token.Service
}

func NewSuggestionController(suggestion suggestion.Service, token token.Service) *SuggestionController {
	return &SuggestionController{
		SuggestionService: suggestion,
		TokenService:      token,
	}
}

func (c *SuggestionController) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.IsAuthorized(c.TokenService))
	r.Post("/", c.Create)
	return r
}

// @tags suggestions
// @summary Create suggestion.
// @description Create suggestion.
// @security BearerJWT
// @id create-suggestion
// @produce json
// @success 201
// @router /suggestions [post]
func (c *SuggestionController) Create(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(middleware.ContextKeyUserID).(string)
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
