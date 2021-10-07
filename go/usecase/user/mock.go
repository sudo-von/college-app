package user

import (
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/pkg/search"
	"time"
)

var (
	fakeUserIDS                 = []string{"615c09f7309d7ded48c7a053", "615c09f7309d7ded48c7a054"}
	fakeUserRegistrationNumbers = []string{"16190775", "16190776"}
	fakeUserEmails              = []string{"sudo.von.contact@gmail.com", "sudovonctf@gmail.com"}
)

type UserReaderMock struct{}

func (u UserReaderMock) GetTinyUserByID(userID string) (*entity.TinyUser, error) {

	err := search.FindInSlice(userID, fakeUserIDS)
	if err != nil {
		return nil, err
	}

	tinyUser := &entity.TinyUser{
		ID:                 userID,
		Name:               "Von",
		BirthDate:          time.Date(1997, 04, 17, 0, 0, 0, 0, time.Local),
		Email:              "sudo.von.contact@gmail.com",
		RegistrationNumber: "16190775",
	}

	return tinyUser, nil
}

func (u UserReaderMock) GetTinyUserByEmail(email string) (*entity.TinyUser, error) {

	err := search.FindInSlice(email, fakeUserEmails)
	if err != nil {
		return nil, err
	}

	tinyUser := &entity.TinyUser{
		ID:                 email,
		Name:               "Von",
		BirthDate:          time.Date(1997, 04, 17, 0, 0, 0, 0, time.Local),
		Email:              "sudo.von.contact@gmail.com",
		RegistrationNumber: "16190775",
	}

	return tinyUser, nil
}

func (u UserReaderMock) GetTinyUserByRegistrationNumber(registrationNumber string) (*entity.TinyUser, error) {

	err := search.FindInSlice(registrationNumber, fakeUserRegistrationNumbers)
	if err != nil {
		return nil, err
	}

	tinyUser := &entity.TinyUser{
		ID:                 registrationNumber,
		Name:               "Von",
		BirthDate:          time.Date(1997, 04, 17, 0, 0, 0, 0, time.Local),
		Email:              "sudo.von.contact@gmail.com",
		RegistrationNumber: "16190775",
	}

	return tinyUser, nil
}

func (u UserReaderMock) GetUserByID(userID string) (*entity.User, error) {

	err := search.FindInSlice(userID, fakeUserIDS)
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:                 userID,
		Name:               "Von",
		BirthDate:          time.Date(1997, 04, 17, 0, 0, 0, 0, time.Local),
		Email:              "sudo.von.contact@gmail.com",
		Password:           "12345678",
		RegistrationNumber: "16190775",
		Status:             entity.ActiveStatus,
		Role:               entity.StudentRole,
		University: entity.University{
			ID:             "615c09f7309d7ded48c7a049",
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
	}, nil
}

func (u UserReaderMock) GetUserByEmail(email string) (*entity.User, error) {

	err := search.FindInSlice(email, fakeUserEmails)
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:                 "615c09f7309d7ded48c7a053",
		Name:               "Von",
		BirthDate:          time.Date(1997, 04, 17, 0, 0, 0, 0, time.Local),
		Email:              email,
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
	}, nil
}

type UserWriterMock struct{}

func (u UserWriterMock) CreateUser(newUser entity.UserPayload) error {
	return nil
}

func (u UserWriterMock) UpdateUser(newUser entity.UserPayload) error {
	return nil
}

type UserRepositoryMock struct {
	UserReaderMock
	UserWriterMock
}
