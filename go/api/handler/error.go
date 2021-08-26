package handler

import (
	"errors"
	"fmt"
	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/entity"
	"net/http"

	"github.com/go-chi/render"
)

func CheckError(err error, w http.ResponseWriter, r *http.Request) {
	// Prints the wrapped error for debugging purposes.
	fmt.Println(err)
	// Gets the first error which was wrapped.
	for errors.Unwrap(err) != nil {
		err = errors.Unwrap(err)
	}
	// Renders the response depending on th error type.
	switch err.(type) {
	case *entity.ErrorBadRequest:
		render.Render(w, r, presenter.ErrorBadRequestResponse(err))
	case *entity.ErrorUnauthorized:
		render.Render(w, r, presenter.ErrorUnauthorizedResponse(err))
	case *entity.ErrorForbidden:
		render.Render(w, r, presenter.ErrorForbiddenResponse(err))
	case *entity.ErrorNotFound:
		render.Render(w, r, presenter.ErrorNotFoundResponse(err))
	case *entity.ErrorConflict:
		render.Render(w, r, presenter.ErrorConflict(err))
	case *entity.ErrorInternalServer:
		render.Render(w, r, presenter.ErrorInternalServerResponse(err))
	default:
		render.Render(w, r, presenter.ErrorInternalServerResponse(err))
	}
}
