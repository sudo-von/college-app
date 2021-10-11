package handler

import (
	"context"
	"errors"
	"fmt"
	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/pkg/token"
	"strings"
	"time"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type TokenService interface {
	CreateToken(user *entity.User, duration time.Duration) (string, error)
	VerifyToken(token string) (*token.Payload, error)
}

type JWTController struct {
	UserService  UserService
	TokenService TokenService
}

func NewJWTController(userService UserService, tokenService TokenService) *JWTController {
	return &JWTController{
		UserService:  userService,
		TokenService: tokenService,
	}
}

func (c *JWTController) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(c.BasicAuth())
	r.Post("/login", c.Login)
	return r
}

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

func (c *JWTController) BasicAuth() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			email, password, ok := r.BasicAuth()
			if ok {
				user, err := c.UserService.AuthenticateUser(email, password)
				if err != nil {
					err := errors.New("invalid credentials")
					render.Render(w, r, presenter.ErrorUnauthorizedResponse(err, presenter.ErrInvCredentials))
					return
				}
				ctx := context.WithValue(r.Context(), ContextKeyUser, user)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
			err := errors.New("unauthorized user")
			render.Render(w, r, presenter.ErrorUnauthorizedResponse(err, presenter.ErrUnauthorizedUser))
		})
	}
}

func (c *JWTController) IsAuthorized() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authorizationHeader := r.Header.Get("Authorization")
			if len(authorizationHeader) == 0 {
				err := errors.New("authorization header is not provided")
				render.Render(w, r, presenter.ErrorUnauthorizedResponse(err, presenter.ErrAuthHeaderNotProvided))
				return
			}

			fields := strings.Fields(authorizationHeader)
			if len(fields) != 2 {
				err := errors.New("invalid authorization header format")
				render.Render(w, r, presenter.ErrorUnauthorizedResponse(err, presenter.ErrInvAuthHeaderFormat))
				return
			}

			authorizationType := fields[0]
			if authorizationType != "Bearer" {
				err := fmt.Errorf("unsupported authorization type %s", authorizationType)
				render.Render(w, r, presenter.ErrorUnauthorizedResponse(err, presenter.ErrInvUnsHeaderFormat))
				return
			}

			accessToken := fields[1]
			payload, err := c.TokenService.VerifyToken(accessToken)
			if err != nil {
				render.Render(w, r, presenter.ErrorUnauthorizedResponse(err, presenter.ErrInvToken))
				return
			}

			ctx := context.WithValue(r.Context(), ContextKeyUserID, payload.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// @tags authentication
// @summary JWT.
// @description Gets Bearer token.
// @id login
// @security BasicAuth
// @success 200
// @header 200 {string} Authorization "Bearer jwt that must be used as Api Key in the Authorize section."
// @header 200 {string} Access-Control-Allow-Headers "Authorization."
// @router /auth/login [post]
func (c *JWTController) Login(w http.ResponseWriter, r *http.Request) {

	user, ok := r.Context().Value(ContextKeyUser).(*entity.User)
	if !ok {
		err := errors.New("user not in context")
		CheckError(err, w, r)
		return
	}

	signedToken, err := c.TokenService.CreateToken(user, 15)
	if err != nil {
		CheckError(err, w, r)
		return
	}

	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", signedToken))
	w.Header().Set("Access-Control-Allow-Headers", "Authorization")
	render.Status(r, http.StatusOK)
}
