package presenter

type UserMoodResponse struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	Mood         int    `json:"mood"`
	CreationDate string `json:"creation_date"`
}
