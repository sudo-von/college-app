package handler

import (
	"freelancer/college-app/go/api/presenter"
	"net/http"

	"github.com/go-chi/render"
)

func CheckError(err error, w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, presenter.ErrInternalServer(err))
}
