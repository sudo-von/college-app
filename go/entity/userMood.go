package entity

import (
	"fmt"
	"time"
)

type UserMood struct {
	ID           string
	UserID       string
	Mood         float64
	CreationDate time.Time
}

type UserMoodPayload struct {
	ID           string
	UserID       string
	Mood         float64
	CreationDate time.Time
}

type UserMoodFilters struct {
	CreationDate *time.Time
}

// ValidateMood checks if given mood is a valid number in the specified range.
func (ump *UserMoodPayload) ValidateMood() error {
	validMood := false
	minValue := 0.0
	maxValue := 10.0
	if ump.Mood >= minValue && ump.Mood <= maxValue {
		validMood = true
	}
	if !validMood {
		return fmt.Errorf("invalid mood, not in range from %.2f to %.2f", minValue, maxValue)
	}
	return nil
}
