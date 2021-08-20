package token

import "time"

type Service interface {
	CreateToken(userID, userName string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
