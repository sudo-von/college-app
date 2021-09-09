package token

import (
	"freelancer/college-app/go/entity"
	"time"
)

type Service interface {
	CreateToken(user *entity.User, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
