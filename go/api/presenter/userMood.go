package presenter

import (
	"freelancer/college-app/go/entity"
	"net/http"
)

type UserMoodResponse struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	Mood         int    `json:"mood"`
	CreationDate string `json:"creation_date"`
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
	Mood int `json:"mood"`
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
