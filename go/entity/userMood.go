package entity

import (
	"fmt"
	"time"
)

type UserMood struct {
	ID           string
	UserID       string
	Mood         int
	CreationDate time.Time
}

type UserMoodPayload struct {
	ID           string
	UserID       string
	Mood         int
	CreationDate time.Time
}

type UserMoodFilters struct {
	CreationDate *time.Time
}

// ValidateMood checks if given mood is a valid number in the specified range.
func (ump *UserMoodPayload) ValidateMood() error {
	validMood := false
	minValue := 0
	maxValue := 10
	if ump.Mood >= minValue && ump.Mood <= maxValue {
		validMood = true
	}
	if !validMood {
		return fmt.Errorf("invalid mood, not in range from %d to %d", minValue, maxValue)
	}
	return nil
}
