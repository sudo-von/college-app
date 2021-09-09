package presenter

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

var (
	ErrIntServError = "INTERNAL_SERVER_ERROR"
)

type ErrorResponse struct {
	Err     error  `json:"-"`
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *ErrorResponse) Error() string {
	return e.Err.Error()
}

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.Status)
	return nil
}

func ErrorBadRequestResponse(err error, code string) render.Renderer {
	return &ErrorResponse{
		Err:     err,
		Status:  400,
		Code:    code,
		Message: err.Error(),
	}
}

func ErrorUnauthorizedResponse(err error, code string) render.Renderer {
	return &ErrorResponse{
		Err:     err,
		Status:  401,
		Code:    code,
		Message: err.Error(),
	}
}

func ErrorForbiddenResponse(err error, code string) render.Renderer {
	return &ErrorResponse{
		Err:     err,
		Status:  403,
		Code:    code,
		Message: err.Error(),
	}
}

func ErrorNotFoundResponse(err error, code string) render.Renderer {
	return &ErrorResponse{
		Err:     err,
		Status:  404,
		Code:    code,
		Message: err.Error(),
	}
}

func ErrorConflict(err error, code string) render.Renderer {
	return &ErrorResponse{
		Err:     err,
		Status:  409,
		Code:    code,
		Message: err.Error(),
	}
}

func ErrorInternalServerResponse(err error, code string) render.Renderer {
	return &ErrorResponse{
		Err:     err,
		Status:  500,
		Code:    code,
		Message: err.Error(),
	}
}

func mergeErrors(e1, e2 error) error {
	if e1 != nil && e2 != nil {
		return errors.New(e1.Error() + "; " + e2.Error())
	}
	if e1 != nil {
		return e1
	}
	return e2
}
