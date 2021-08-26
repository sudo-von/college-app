package middleware

import (
	"context"
	"errors"
	"fmt"
	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/token"
	"freelancer/college-app/go/usecase/user"
	"strings"

	"net/http"

	"github.com/go-chi/render"
)

type ContextKey struct {
	Name string
}

func (k *ContextKey) String() string {
	return "studyapp context value " + k.Name
}

var (
	ContextKeyUser   = &ContextKey{"User"}
	ContextKeyUserID = &ContextKey{"UserID"}
)

func BasicAuth(uses *user.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			email, password, ok := r.BasicAuth()
			if ok {
				user, err := uses.AuthenticateUser(email, password)
				if err != nil {
					http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
					return
				}
				ctx := context.WithValue(r.Context(), ContextKeyUser, user)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		})
	}
}

func IsAuthorized(uses token.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authorizationHeader := r.Header.Get("Authorization")
			if len(authorizationHeader) == 0 {
				err := errors.New("authorization header is not provided")
				render.Render(w, r, presenter.ErrorUnauthorizedResponse(err))
				return
			}

			fields := strings.Fields(authorizationHeader)
			if len(fields) != 2 {
				err := errors.New("invalid authorization header format")
				render.Render(w, r, presenter.ErrorUnauthorizedResponse(err))
				return
			}

			authorizationType := fields[0]
			if authorizationType != "Bearer" {
				err := fmt.Errorf("unsupported authorization type %s", authorizationType)
				render.Render(w, r, presenter.ErrorUnauthorizedResponse(err))
				return
			}

			accessToken := fields[1]
			payload, err := uses.VerifyToken(accessToken)
			if err != nil {
				render.Render(w, r, presenter.ErrorUnauthorizedResponse(err))
				return
			}

			ctx := context.WithValue(r.Context(), ContextKeyUserID, payload.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
