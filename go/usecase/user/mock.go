package user

import (
	"errors"
	"freelancer/college-app/go/entity"
	"time"
)

func NewFakeTinyUser() *entity.TinyUser {
	return &entity.TinyUser{
		ID:                 "615c09f7309d7ded48c7a053",
		Name:               "Von",
		BirthDate:          time.Date(1997, 04, 17, 0, 0, 0, 0, time.Local),
		Email:              "sudo.von.contact@gmail.com",
		RegistrationNumber: "16190775",
	}
}

func NewFakeUser() *entity.User {
	return &entity.User{
		ID:                 "615c09f7309d7ded48c7a053",
		Name:               "Von",
		BirthDate:          time.Date(1997, 04, 17, 0, 0, 0, 0, time.Local),
		Email:              "sudo.von.contact@gmail.com",
		Password:           "12345678",
		RegistrationNumber: "16190775",
		Status:             entity.ActiveStatus,
		Role:               entity.StudentRole,
		University: entity.University{
			ID:             "615c0794590f4315693633a6",
			Name:           "Fake university name",
			ProfilePicture: "Fake profile picture url",
			Classrooms: []entity.Classroom{
				{
					ID:   "615c07fb5df8db802627eb6c",
					Name: "Fake classroom name",
				},
			},
		},
		CreationDate: time.Date(2021, 01, 01, 0, 0, 0, 0, time.Local),
	}
}

type UserReaderMock struct{}

func (u UserReaderMock) GetTinyUserByID(userID string) (*entity.TinyUser, error) {
	if userID != "615c09f7309d7ded48c7a053" {
		return nil, errors.New("not found")
	}
	return NewFakeTinyUser(), nil
}

func (u UserReaderMock) GetTinyUserByEmail(email string) (*entity.TinyUser, error) {
	if email != "sudo.von.contact@gmail.com" {
		return nil, errors.New("not found")
	}
	return NewFakeTinyUser(), nil
}

func (u UserReaderMock) GetTinyUserByRegistrationNumber(registrationNumber string) (*entity.TinyUser, error) {
	if registrationNumber != "16190775" {
		return nil, errors.New("not found")
	}
	return NewFakeTinyUser(), nil
}

func (u UserReaderMock) GetUserByID(userID string) (*entity.User, error) {
	if userID != "615c09f7309d7ded48c7a053" {
		return nil, errors.New("not found")
	}
	return NewFakeUser(), nil
}

func (u UserReaderMock) GetUserByEmail(email string) (*entity.User, error) {
	if email != "sudo.von.contact@gmail.com" {
		return nil, errors.New("not found")
	}
	return NewFakeUser(), nil
}

type UserWriterMock struct{}

func (u UserWriterMock) CreateUser(newUser entity.UserPayload) error {
	return nil
}

func (u UserWriterMock) UpdateUser(newUser entity.UserPayload) error {
	return nil
}
