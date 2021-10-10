package contact

import (
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/pkg/search"
	"time"
)

var (
	contactIDS  = []string{"615c09f7309d7ded48c7a0c7", "615c09f7309d7ded48c7a0c8"}
	fakeUserIDS = []string{"615c09f7309d7ded48c7a053", "615c09f7309d7ded48c7a055"}
)

type ContactReaderMock struct{}

func (c ContactReaderMock) GetContactByID(contactID string) (*entity.Contact, error) {

	err := search.FindInSlice(contactID, contactIDS)
	if err != nil {
		return nil, err
	}

	contact := &entity.Contact{
		ID:            contactID,
		UserID:        "615c09f7309d7ded48c7a053",
		ContactName:   "Fake contact name",
		ContactNumber: "Fake contact number",
		Message:       "Fake message",
		CreationDate:  time.Date(2021, 01, 01, 0, 0, 0, 0, time.Local),
	}
	return contact, nil
}

func (c ContactReaderMock) GetContactByUserID(userID string) (*entity.Contact, error) {

	err := search.FindInSlice(userID, fakeUserIDS)
	if err != nil {
		return nil, err
	}

	contact := &entity.Contact{
		ID:            "615c09f7309d7ded48c7a0c7",
		UserID:        userID,
		ContactName:   "Fake contact name",
		ContactNumber: "Fake contact number",
		Message:       "Fake message",
		CreationDate:  time.Date(2021, 01, 01, 0, 0, 0, 0, time.Local),
	}
	return contact, nil
}

type ContactWriterMock struct{}

func (c ContactWriterMock) CreateContact(newContact entity.ContactPayload) error {
	return nil
}
func (c ContactWriterMock) UpdateContactByUserID(newContact entity.ContactPayload) error {
	return nil
}

type ContactRepositoryMock struct {
	ContactReaderMock
	ContactWriterMock
}
