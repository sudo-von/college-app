package handler

import (
	"fmt"
	"freelancer/college-app/go/api/middleware"
	"freelancer/college-app/go/config"
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/usecase/jwt"
	"net/http"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type JWTController struct {
	JWTService jwt.Service
}

func NewJWTController(jwt jwt.Service) *JWTController {
	return &JWTController{
		JWTService: jwt,
	}
}

func (c *JWTController) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.BasicAuth(&c.JWTService))
	r.Post("/", c.Login)
	return r
}

// Login returns a jwt in case of successful authentication.
func (c *JWTController) Login(w http.ResponseWriter, r *http.Request) {

	user, ok := r.Context().Value(middleware.ContextKeyUser).(*entity.User)
	if !ok {
		render.Status(r, 500)
		return
	}

	signedToken, err := c.GenerateJWT(user)
	if err != nil {
		render.Status(r, 500)
		return
	}

	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", *signedToken))
	w.Header().Set("Access-Control-Allow-Headers", "Authorization")
	render.Status(r, http.StatusOK)
}

func (c *JWTController) GenerateJWT(user *entity.User) (*string, error) {

	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, jwtgo.MapClaims{
		"exp":  15000,
		"iat":  time.Now().Unix(),
		"id":   user.ID,
		"name": user.Name,
	})

	signedToken, err := token.SignedString([]byte(config.SECRET_KEY))
	if err != nil {
		return nil, err
	}
	return &signedToken, nil
}
