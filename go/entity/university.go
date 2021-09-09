package entity

import (
	"errors"
)

type TinyUniversity struct {
	ID             string
	Name           string
	ProfilePicture string
}

type University struct {
	ID             string
	Name           string
	ProfilePicture string
	Classrooms     []Classroom
}

type UniversityPayload struct {
	ID             string
	Name           string
	ProfilePicture string
}

// ValidateClassroom checks if given classroom id is a valid classroom from that university.
func (u *University) ValidateClassroom(classroomID string) error {
	validClassroom := false
	for _, c := range u.Classrooms {
		if c.ID == classroomID {
			validClassroom = true
		}
	}
	if !validClassroom {
		return errors.New("invalid classroom id")
	}
	return nil
}
