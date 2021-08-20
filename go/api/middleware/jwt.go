package middleware

import (
	"context"
	"freelancer/college-app/go/usecase/user"

	"net/http"
)

type ContextKey struct {
	Name string
}

func (k *ContextKey) String() string {
	return "studyapp context value " + k.Name
}

var (
	ContextKeyUser = &ContextKey{"User"}
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
