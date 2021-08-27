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
	errorDate := time.Now().In(time.Local).Format("2006-01-02 15:04:05")
	fmt.Println("[error]:", errorDate, err)
	// Gets first structure which was wrapped with multiple errors.
	err = UnwrapError(err)
	// Renders the response depending on th error type.
	switch errorType := err.(type) {
	case *entity.ErrorBadRequest:
		render.Render(w, r, presenter.ErrorBadRequestResponse(UnwrapError(errorType.Message)))
	case *entity.ErrorUnauthorized:
		render.Render(w, r, presenter.ErrorUnauthorizedResponse(UnwrapError(errorType.Message)))
	case *entity.ErrorForbidden:
		render.Render(w, r, presenter.ErrorForbiddenResponse(UnwrapError(errorType.Message)))
	case *entity.ErrorNotFound:
		render.Render(w, r, presenter.ErrorNotFoundResponse(UnwrapError(errorType.Message)))
	case *entity.ErrorConflict:
		render.Render(w, r, presenter.ErrorConflict(UnwrapError(errorType.Message)))
	case *entity.ErrorInternalServer:
		render.Render(w, r, presenter.ErrorInternalServerResponse(UnwrapError(errorType.Message)))
	default:
		render.Render(w, r, presenter.ErrorInternalServerResponse(UnwrapError(err.(*entity.ErrorInternalServer).Message)))
	}
}

func UnwrapError(err error) error {
	for errors.Unwrap(err) != nil {
		err = errors.Unwrap(err)
	}
	return err
}
