package handler

import (
	"errors"
	"fmt"
	"freelancer/college-app/go/api/middleware"
	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/pkg/token"
	"freelancer/college-app/go/usecase/advice"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type AdviceController struct {
	AdviceService advice.Service
	TokenService  token.Service
}

func NewAdviceController(advice advice.Service, token token.Service) *AdviceController {
	return &AdviceController{
		AdviceService: advice,
		TokenService:  token,
	}
}

func (c *AdviceController) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.IsAuthorized(c.TokenService))
	r.Get("/", c.List)
	r.Post("/", c.Create)
	r.Patch("/{id}", c.Update)
	r.Delete("/{id}", c.Delete)
	return r
}

// List returns a list of advices given a university.
func (c *AdviceController) List(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(middleware.ContextKeyUserID).(string)
	if !ok {
		err := errors.New("user not in context")
		CheckError(entity.NewErrorInternalServer(fmt.Errorf("AdviceController > List: %w", err)), w, r)
		return
	}

	adviceDate := time.Now().In(time.Local)
	adviceFilters := entity.AdviceFilters{
		AdviceDate: &adviceDate,
	}

	list, total, err := c.AdviceService.GetAdvices(userID, adviceFilters)
	if err != nil {
		CheckError(fmt.Errorf("AdviceController > List > GetAdvices: %w", err), w, r)
		return
	}

	res := presenter.AdviceList{
		Total:   *total,
		Advices: make([]presenter.AdviceResponse, 0, len(list)),
	}

	for _, advice := range list {
		res.Advices = append(res.Advices, presenter.ToAdvicePresenter(advice))
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, &res)
}

// Create stores a new advice.
func (c *AdviceController) Create(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(middleware.ContextKeyUserID).(string)
	if !ok {
		err := errors.New("user not in context")
		CheckError(entity.NewErrorInternalServer(fmt.Errorf("AdviceController > Create: %w", err)), w, r)
		return
	}

	var data presenter.AdvicePayload
	if err := render.Bind(r, &data); err != nil {
		CheckError(entity.NewErrorBadRequest(fmt.Errorf("AdviceController > Create: %w", err)), w, r)
		return
	}

	adviceDate, err := time.ParseInLocation("2006-01-02 15:04", data.AdviceDate, time.Local)
	if err != nil {
		CheckError(entity.NewErrorBadRequest(fmt.Errorf("AdviceController > Create: %w", err)), w, r)
		return
	}

	newAdvice := entity.AdvicePayload{
		UserID:       userID,
		Subject:      data.Subject,
		AdviceDate:   adviceDate,
		ClassroomID:  data.ClassroomID,
		Status:       entity.ActiveStatus,
		CreationDate: time.Now().In(time.Local),
	}

	err = c.AdviceService.CreateAdvice(newAdvice)
	if err != nil {
		CheckError(fmt.Errorf("AdviceController > Create > CreateAdvice: %w", err), w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	render.Status(r, http.StatusOK)
}

// Update updates an advice given its id.
func (c *AdviceController) Update(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(middleware.ContextKeyUserID).(string)
	if !ok {
		err := errors.New("user not in context")
		CheckError(entity.NewErrorInternalServer(fmt.Errorf("AdviceController > UpdateAdvice: %w", err)), w, r)
		return
	}
	adviceID := chi.URLParam(r, "id")

	var data presenter.UpdateAdvicePayload
	if err := render.Bind(r, &data); err != nil {
		CheckError(entity.NewErrorBadRequest(fmt.Errorf("AdviceController > UpdateAdvice: %w", err)), w, r)
		return
	}

	adviceDate, err := time.ParseInLocation("2006-01-02 15:04", data.AdviceDate, time.Local)
	if err != nil {
		CheckError(entity.NewErrorBadRequest(fmt.Errorf("AdviceController > UpdateAdvice: %w", err)), w, r)
		return
	}

	newAdvice := entity.UpdateAdvicePayload{
		Subject:     data.Subject,
		AdviceDate:  adviceDate,
		ClassroomID: data.ClassroomID,
	}

	err = c.AdviceService.UpdateAdvice(userID, adviceID, newAdvice)
	if err != nil {
		CheckError(fmt.Errorf("AdviceController > UpdateAdvice > UpdateAdvice: %w", err), w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	render.Status(r, http.StatusOK)
}

// Delete disables an advice given its id.
func (c *AdviceController) Delete(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(middleware.ContextKeyUserID).(string)
	if !ok {
		err := errors.New("user not in context")
		CheckError(entity.NewErrorInternalServer(fmt.Errorf("AdviceController > Delete: %w", err)), w, r)
		return
	}
	adviceID := chi.URLParam(r, "id")

	err := c.AdviceService.DeleteAdvice(userID, adviceID)
	if err != nil {
		CheckError(fmt.Errorf("AdviceController > Delete > DeleteAdvice: %w", err), w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	render.Status(r, http.StatusAccepted)
}
