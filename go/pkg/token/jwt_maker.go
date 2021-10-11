package token

import (
	"errors"
	"fmt"
	"freelancer/college-app/go/entity"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Service struct {
	secretKey string
}

func NewService(secretKey string) (*Service, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &Service{secretKey}, nil
}

const minSecretKeySize = 32

func (s *Service) CreateToken(user *entity.User, duration time.Duration) (string, error) {

	payload, err := NewPayload(user, duration)
	if err != nil {
		return "", err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	signedToken, err := jwtToken.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *Service) VerifyToken(token string) (*Payload, error) {

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New(ErrInvalidToken)
		}
		return []byte(s.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, errors.New(ErrExpiredToken)) {
			return nil, errors.New(ErrExpiredToken)
		}
		return nil, errors.New(ErrInvalidToken)
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, errors.New(ErrInvalidToken)
	}

	return payload, nil
}
