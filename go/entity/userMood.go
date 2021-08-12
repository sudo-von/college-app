package entity

import "time"

type UserMood struct {
	ID       string
	UserID   string
	Mood     int
	MoodDate time.Time
}

type UserMoodPayload struct {
	ID       string
	UserID   string
	Mood     int
	MoodDate time.Time
}

type UserMoodFilters struct {
	MoodDate *time.Time
}
