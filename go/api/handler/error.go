package handler

import (
	"errors"
	"fmt"
	"freelancer/college-app/go/api/presenter"
	"net/http"

	"github.com/go-chi/render"
)

type ErrorBadRequest struct {
	message string
}

func (e *ErrorBadRequest) Error() string {
	return e.message
}

func NewErrorBadRequest(message string) error {
	return &ErrorBadRequest{
		message: message,
	}
}

type ErrorUnauthorized struct {
	message string
}

func (e *ErrorUnauthorized) Error() string {
	return e.message
}

func NewErrorUnauthorized(message string) error {
	return &ErrorUnauthorized{
		message: message,
	}
}

type ErrorForbidden struct {
	message string
}

func (e *ErrorForbidden) Error() string {
	return e.message
}

func NewErrorForbidden(message string) error {
	return &ErrorForbidden{
		message: message,
	}
}

type ErrorNotFound struct {
	message string
}

func (e *ErrorNotFound) Error() string {
	return e.message
}

func NewErrorNotFound(message string) error {
	return &ErrorNotFound{
		message: message,
	}
}

type ErrorInternalServer struct {
	message string
}

func (e *ErrorInternalServer) Error() string {
	return e.message
}

func NewErrorInternalServer(message string) error {
	return &ErrorInternalServer{
		message: message,
	}
}

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
	case *ErrorBadRequest:
		render.Render(w, r, presenter.ErrorBadRequestResponse(unwrappedError))
	case *ErrorUnauthorized:
		render.Render(w, r, presenter.ErrorUnauthorizedResponse(unwrappedError))
	case *ErrorForbidden:
		render.Render(w, r, presenter.ErrorForbiddenResponse(unwrappedError))
	case *ErrorNotFound:
		render.Render(w, r, presenter.ErrorNotFoundResponse(unwrappedError))
	case *ErrorInternalServer:
		render.Render(w, r, presenter.ErrorInternalServerResponse(unwrappedError))
	default:
		render.Render(w, r, presenter.ErrorInternalServerResponse(unwrappedError))
	}
}
