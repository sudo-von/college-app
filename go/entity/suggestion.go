package entity

import "time"

type SuggestionPayload struct {
	ID           string
	UserID       string
	Suggestion   string
	CreationDate time.Time
}
