package handler

import (
	"errors"
	"fmt"
	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/entity"
	"net/http"
	"time"

	"github.com/go-chi/render"
)

func CheckError(err error, w http.ResponseWriter, r *http.Request) {
	// Prints the wrapped error for debugging purposes.
	errorDate := time.Now().In(time.Local).Format("2006-01-02 15:04:00")
	fmt.Println("[error]:", errorDate, err)
	// Gets first structure which was wrapped with multiple errors.
	for errors.Unwrap(err) != nil {
		err = errors.Unwrap(err)
	}
	// Renders the response depending on th error type.
	switch err := err.(type) {
	case *entity.ErrorBadRequest:
		response := errors.Unwrap(err.Message)
		render.Render(w, r, presenter.ErrorBadRequestResponse(response))
	case *entity.ErrorUnauthorized:
		response := errors.Unwrap(err.Message)
		render.Render(w, r, presenter.ErrorUnauthorizedResponse(response))
	case *entity.ErrorForbidden:
		response := errors.Unwrap(err.Message)
		render.Render(w, r, presenter.ErrorForbiddenResponse(response))
	case *entity.ErrorNotFound:
		response := errors.Unwrap(err.Message)
		render.Render(w, r, presenter.ErrorNotFoundResponse(response))
	case *entity.ErrorConflict:
		response := errors.Unwrap(err.Message)
		render.Render(w, r, presenter.ErrorConflict(response))
	case *entity.ErrorInternalServer:
		response := errors.Unwrap(err.Message)
		render.Render(w, r, presenter.ErrorInternalServerResponse(response))
	default:
		response := errors.Unwrap(err.(*entity.ErrorInternalServer).Message)
		render.Render(w, r, presenter.ErrorInternalServerResponse(response))
	}
}
