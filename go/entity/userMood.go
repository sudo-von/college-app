package entity

import "time"

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
