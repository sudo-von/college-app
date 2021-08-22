package presenter

import "net/http"

type UserMoodResponse struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	Mood         int    `json:"mood"`
	CreationDate string `json:"creation_date"`
}

func (ur *UserMoodResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
