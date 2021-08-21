package entity

import "time"

type Contact struct {
	ID            string
	UserID        string
	ContactName   string
	ContactNumber string
	Message       string
	CreationDate  time.Time
}

type ContactPayload struct {
	ID            string
	UserID        string
	ContactName   string
	ContactNumber string
	Message       string
	CreationDate  time.Time
}

type UpdateContactPayload struct {
	ContactName   string
	ContactNumber string
	Message       string
}
