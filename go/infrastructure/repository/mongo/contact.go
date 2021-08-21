package mongo

import (
	"errors"
	"freelancer/college-app/go/entity"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ContactModel struct {
	ID            bson.ObjectId `bson:"_id"`
	UserID        bson.ObjectId `bson:"user_id"`
	ContactName   string        `bson:"contact_name"`
	ContactNumber string        `bson:"contact_number"`
	Message       string        `bson:"message"`
	CreationDate  time.Time     `bson:"creation_date"`
}

func toEntityContact(contact ContactModel) entity.Contact {
	return entity.Contact{
		ID:            contact.ID.Hex(),
		UserID:        contact.UserID.Hex(),
		ContactName:   contact.ContactName,
		ContactNumber: contact.ContactNumber,
		Message:       contact.Message,
		CreationDate:  contact.CreationDate,
	}
}

type ContactPayloadModel struct {
	ID            bson.ObjectId `bson:"_id"`
	UserID        bson.ObjectId `bson:"user_id"`
	ContactName   string        `bson:"contact_name"`
	ContactNumber string        `bson:"contact_number"`
	Message       string        `bson:"message"`
	CreationDate  time.Time     `bson:"creation_date"`
}

func toContactPayloadModel(contactPayload entity.ContactPayload) ContactPayloadModel {

	var contactID bson.ObjectId
	if contactPayload.ID != "" {
		contactID = bson.ObjectIdHex(contactPayload.ID)
	} else {
		contactID = bson.NewObjectId()
	}

	var userID bson.ObjectId
	if contactPayload.UserID != "" {
		userID = bson.ObjectIdHex(contactPayload.UserID)
	} else {
		userID = bson.NewObjectId()
	}

	return ContactPayloadModel{
		ID:            contactID,
		UserID:        userID,
		ContactName:   contactPayload.ContactName,
		ContactNumber: contactPayload.ContactNumber,
		Message:       contactPayload.Message,
		CreationDate:  contactPayload.CreationDate,
	}
}

type ContactRepository struct {
	Session      *mgo.Session
	DatabaseName string
}

func NewContactRepository(repository *Repository) *ContactRepository {
	return &ContactRepository{
		Session:      repository.Session,
		DatabaseName: repository.DatabaseName,
	}
}

func (r *ContactRepository) CreateContact(newContact entity.ContactPayload) error {

	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("contacts")

	contactM := toContactPayloadModel(newContact)
	err := com.Insert(&contactM)
	if err != nil {
		return err
	}

	return nil
}

func (r *ContactRepository) GetContactByUserID(userID string) (*entity.Contact, error) {

	if !bson.IsObjectIdHex(userID) {
		return nil, errors.New("given user_id is not a valid hex")
	}

	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("contacts")
	searchQuery := bson.M{
		"user_id": bson.ObjectIdHex(userID),
	}

	var contactM ContactModel
	err := com.Find(searchQuery).One(&contactM)
	if err != nil {
		return nil, err
	}

	contact := toEntityContact(contactM)
	return &contact, nil
}

func (r *ContactRepository) UpdateContactByUserID(updatedContact entity.ContactPayload) error {

	if !bson.IsObjectIdHex(updatedContact.ID) {
		return errors.New("given user_id is not a valid hex")
	}

	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("contacts")
	searchQuery := bson.M{
		"user_id": bson.ObjectIdHex(updatedContact.UserID),
	}

	contactM := toContactPayloadModel(updatedContact)
	err := com.Update(searchQuery, &contactM)
	if err != nil {
		return err
	}

	return nil
}
