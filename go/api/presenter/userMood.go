package presenter

import (
	"freelancer/college-app/go/entity"
	"net/http"
)

var (
	ErrUserMoodNotFound          = "USER_MOOD_NOT_FOUND"
	ErrUserMoodAlreadyRegistered = "USER_MOOD_ALREADY_REGISTERED"
	ErrInvMood                   = "INVALID_MOOD_VALUE"
)

type UserMoodResponse struct {
	ID           string  `json:"id" example:"613aab578a6ef50007e622be"`
	UserID       string  `json:"user_id" example:"613aab4d8a6ef50007e622bd"`
	Mood         float64 `json:"mood" example:"10"`
	CreationDate string  `json:"creation_date" example:"2021-10-10"`
}

func (ur *UserMoodResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ToUserMoodPresenter(um entity.UserMood) UserMoodResponse {
	return UserMoodResponse{
		ID:           um.ID,
		UserID:       um.UserID,
		Mood:         um.Mood,
		CreationDate: um.CreationDate.Local().Format("2006-01-02"),
	}
}

type UserMoodPayload struct {
	Mood float64 `json:"mood" example:"5"`
}

func (ump *UserMoodPayload) validate() (err error) {
	return
}

func (ump *UserMoodPayload) Bind(r *http.Request) error {
	if err := ump.validate(); err != nil {
		return err
	}
	return nil
}
