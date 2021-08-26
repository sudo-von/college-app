package handler

import (
	"net/http"
	"time"

	"freelancer/college-app/go/api/middleware"
	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/token"
	"freelancer/college-app/go/usecase/advice"

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
	return r
}

// List returns a list of advices given a university.
func (c *AdviceController) List(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(middleware.ContextKeyUserID).(string)
	if !ok {
		CheckError(entity.NewErrorInternalServer("user not in context"), w, r)
		return
	}

	adviceDate := time.Now().In(time.Local)
	adviceFilters := entity.AdviceFilters{
		AdviceDate: &adviceDate,
	}

	list, total, err := c.AdviceService.GetAdvices(userID, adviceFilters)
	if err != nil {
		CheckError(err, w, r)
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
		CheckError(entity.NewErrorInternalServer("user not in context"), w, r)
		return
	}

	var data presenter.AdvicePayload
	if err := render.Bind(r, &data); err != nil {
		CheckError(entity.NewErrorBadRequest(err.Error()), w, r)
		return
	}

	adviceDate, err := time.ParseInLocation("2006-01-02 15:04", data.AdviceDate, time.Local)
	if err != nil {
		CheckError(err, w, r)
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
		CheckError(err, w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	render.Status(r, http.StatusOK)
}
