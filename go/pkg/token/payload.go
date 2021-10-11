package token

import (
	"errors"
	"freelancer/college-app/go/entity"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidToken = "invalid token"
	ErrExpiredToken = "expired token"
)

type Payload struct {
	ID                       uuid.UUID `json:"id"`
	UserID                   string    `json:"user_id"`
	UserName                 string    `json:"user_name"`
	UniversityID             string    `json:"university_id"`
	UniversityName           string    `json:"university_name"`
	UniversityProfilePicture string    `json:"university_profile_picture"`
	IssuedAt                 time.Time `json:"issued_at"`
	ExpiredAt                time.Time `json:"expired_at"`
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
	payload := &Payload{
		ID:                       id,
		UserID:                   user.ID,
		UserName:                 user.Name,
		UniversityID:             user.University.ID,
		UniversityName:           user.University.Name,
		UniversityProfilePicture: user.University.ProfilePicture,
		IssuedAt:                 time.Now(),
		ExpiredAt:                time.Now().Add(time.Minute * duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return errors.New(ErrExpiredToken)
	}
	return nil
}
