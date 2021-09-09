package token

import (
	"errors"
	"freelancer/college-app/go/entity"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("expired token")
)

type Payload struct {
	ID         uuid.UUID         `json:"id"`
	User       UserPayload       `json:"user"`
	University UniversityPayload `json:"university"`
	UserName   string            `json:"user_name"`
	IssuedAt   time.Time         `json:"issued_at"`
	ExpiredAt  time.Time         `json:"expired_at"`
}

type UserPayload struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UniversityPayload struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	ProfilePicture string `json:"profile_picture"`
}

func NewPayload(user *entity.User, duration time.Duration) (*Payload, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	userPayload := UserPayload{
		ID:   user.ID,
		Name: user.Name,
	}
	universityPayload := UniversityPayload{
		ID:             user.University.ID,
		Name:           user.University.Name,
		ProfilePicture: user.University.ProfilePicture,
	}
	payload := &Payload{
		ID:         id,
		User:       userPayload,
		University: universityPayload,
		IssuedAt:   time.Now(),
		ExpiredAt:  time.Now().Add(time.Minute * duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
