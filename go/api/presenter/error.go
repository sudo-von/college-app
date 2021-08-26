package presenter

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

type ErrorResponse struct {
	Err            error  `json:"-"`
	HTTPStatusCode int    `json:"-"`
	ErrorCode      int    `json:"error_code"`
	Message        string `json:"message"`
}

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrorBadRequestResponse(err error) render.Renderer {
	return &ErrorResponse{
		Err:            err,
		HTTPStatusCode: 400,
		ErrorCode:      400,
		Message:        err.Error(),
	}
}

func ErrorUnauthorizedResponse(err error) render.Renderer {
	return &ErrorResponse{
		Err:            err,
		HTTPStatusCode: 401,
		ErrorCode:      401,
		Message:        err.Error(),
	}
}

func ErrorForbiddenResponse(err error) render.Renderer {
	return &ErrorResponse{
		Err:            err,
		HTTPStatusCode: 403,
		ErrorCode:      403,
		Message:        err.Error(),
	}
}

func ErrorNotFoundResponse(err error) render.Renderer {
	return &ErrorResponse{
		Err:            err,
		HTTPStatusCode: 404,
		ErrorCode:      404,
		Message:        err.Error(),
	}
}

func ErrorConflict(err error) render.Renderer {
	return &ErrorResponse{
		Err:            err,
		HTTPStatusCode: 409,
		ErrorCode:      409,
		Message:        err.Error(),
	}
}

func ErrorInternalServerResponse(err error) render.Renderer {
	return &ErrorResponse{
		Err:            err,
		HTTPStatusCode: 500,
		ErrorCode:      500,
		Message:        err.Error(),
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
