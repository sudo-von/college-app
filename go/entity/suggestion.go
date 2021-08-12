package entity

import "time"

type SuggestionPayload struct {
	ID             string
	Suggestion     string
	UserID         string
	SuggestionDate time.Time
}
