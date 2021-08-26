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
	var unwrappedError error
	for errors.Unwrap(err) != nil {
		unwrappedError = errors.Unwrap(err)
	}
	// Renders the response depending on th error type.
	switch err.(type) {
	case *entity.ErrorBadRequest:
		render.Render(w, r, presenter.ErrorBadRequestResponse(unwrappedError))
	case *entity.ErrorUnauthorized:
		render.Render(w, r, presenter.ErrorUnauthorizedResponse(unwrappedError))
	case *entity.ErrorForbidden:
		render.Render(w, r, presenter.ErrorForbiddenResponse(unwrappedError))
	case *entity.ErrorNotFound:
		render.Render(w, r, presenter.ErrorNotFoundResponse(unwrappedError))
	case *entity.ErrorInternalServer:
		render.Render(w, r, presenter.ErrorInternalServerResponse(unwrappedError))
	default:
		render.Render(w, r, presenter.ErrorInternalServerResponse(unwrappedError))
	}
}
