package handler

import (
	"errors"
	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/entity"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type DepartmentService interface {
	GetDepartments(userID string, departmentFilters entity.DepartmentFilters) ([]entity.Department, *int, error)
}

type DepartmentController struct {
	DepartmentService DepartmentService
	AuthService       func(http.Handler) http.Handler
}

func NewDepartmentController(departmentService DepartmentService, authService func(http.Handler) http.Handler) *DepartmentController {
	return &DepartmentController{
		DepartmentService: departmentService,
		AuthService:       authService,
	}
}

func (c *DepartmentController) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(c.AuthService)
	r.Get("/", c.List)
	return r
}

// @tags departments
// @summary List departments.
// @description List departments.
// @security BearerJWT
// @id list-departments
// @produce json
// @success 200 {object} presenter.DepartmentList
// @router /departments [get]
func (c *DepartmentController) List(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(ContextKeyUserID).(string)
	if !ok {
		err := errors.New("user not in context")
		CheckError(err, w, r)
		return
	}

	departmentFilters := entity.DepartmentFilters{}

	list, total, err := c.DepartmentService.GetDepartments(userID, departmentFilters)
	if err != nil {
		CheckError(err, w, r)
		return
	}

	res := presenter.DepartmentList{
		Total:       *total,
		Departments: make([]presenter.DepartmentResponse, 0, len(list)),
	}

	for _, department := range list {
		res.Departments = append(res.Departments, presenter.ToDepartmentPresenter(department))
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, &res)
}
