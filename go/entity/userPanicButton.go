package entity

import "time"

type UserPanicButtonPayload struct {
	ID            string
	UserID        string
	ContactName   string
	ContactNumber string
	Message       string
	CreationDate  time.Time
}
