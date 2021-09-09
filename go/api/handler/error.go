package handler

import (
	"fmt"
	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/entity"
	"net/http"
	"time"

	"github.com/go-chi/render"
)

func CheckError(err error, w http.ResponseWriter, r *http.Request) {
	// Prints message for debugging purposes.
	errorDate := time.Now().In(time.Local).Format("2006-01-02 15:04:05")
	fmt.Println("[error]:", errorDate, err)
	// Renders the response depending on the error type.
	switch errorType := err.(type) {
	case *entity.ErrorBadRequest:
		render.Render(w, r, presenter.ErrorBadRequestResponse(errorType.Message, errorType.Code))
	case *entity.ErrorUnauthorized:
		render.Render(w, r, presenter.ErrorUnauthorizedResponse(errorType.Message, errorType.Code))
	case *entity.ErrorForbidden:
		render.Render(w, r, presenter.ErrorForbiddenResponse(errorType.Message, errorType.Code))
	case *entity.ErrorNotFound:
		render.Render(w, r, presenter.ErrorNotFoundResponse(errorType.Message, errorType.Code))
	case *entity.ErrorConflict:
		render.Render(w, r, presenter.ErrorConflict(errorType.Message, errorType.Code))
	case *entity.ErrorInternalServer:
		render.Render(w, r, presenter.ErrorInternalServerResponse(errorType.Message, errorType.Code))
	default:
		render.Render(w, r, presenter.ErrorInternalServerResponse(err.(*entity.ErrorInternalServer).Message, err.(*entity.ErrorInternalServer).Code))
	}
}
