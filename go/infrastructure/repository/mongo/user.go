package mongo

import (
	"errors"
	"time"

	"freelancer/college-app/go/entity"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TinyUserModel struct {
	ID                 bson.ObjectId `bson:"_id"`
	Name               string        `bson:"name"`
	BirthDate          time.Time     `bson:"birth_date"`
	Email              string        `bson:"email"`
	RegistrationNumber string        `bson:"registration_number"`
}

type UserPayloadModel struct {
	ID                 bson.ObjectId `bson:"_id"`
	Name               string        `bson:"name"`
	BirthDate          time.Time     `bson:"birth_date"`
	Email              string        `bson:"email"`
	RegistrationNumber string        `bson:"registration_number"`
	Password           string        `bson:"password"`
	Status             string        `bson:"status"`
	UniversityID       string        `bson:"university_id"`
	CreationDate       time.Time     `bson:"creation_date"`
}

func toUserPayloadModel(userPayload entity.UserPayload) UserPayloadModel {

	var userID bson.ObjectId
	if userPayload.ID != "" {
		userID = bson.ObjectIdHex(userPayload.ID)
	} else {
		userID = bson.NewObjectId()
	}

	return UserPayloadModel{
		ID:                 userID,
		Name:               userPayload.Name,
		BirthDate:          userPayload.BirthDate,
		Email:              userPayload.Email,
		RegistrationNumber: userPayload.RegistrationNumber,
		Password:           userPayload.Password,
		Status:             userPayload.Status,
		UniversityID:       userPayload.UniversityID,
		CreationDate:       userPayload.CreationDate,
	}
}

func toEntityTinyUser(user TinyUserModel) entity.TinyUser {
	return entity.TinyUser{
		ID:                 user.ID.Hex(),
		Name:               user.Name,
		BirthDate:          user.BirthDate,
		Email:              user.Email,
		RegistrationNumber: user.RegistrationNumber,
	}
}

type UserRepository struct {
	Session      *mgo.Session
	DatabaseName string
}

func NewUserRepository(repository *Repository) *UserRepository {
	return &UserRepository{
		Session:      repository.Session,
		DatabaseName: repository.DatabaseName,
	}
}

func (r *UserRepository) GetUserByID(userID string) (*entity.TinyUser, error) {

	if !bson.IsObjectIdHex(userID) {
		return nil, errors.New("given user_id is not a valid hex")
	}

	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("users")
	searchQuery := bson.M{
		"_id":    bson.ObjectIdHex(userID),
		"status": entity.ActiveStatus,
	}

	var tinyUserM TinyUserModel
	err := com.Find(searchQuery).One(&tinyUserM)
	if err != nil {
		return nil, err
	}

	userApi := toEntityTinyUser(tinyUserM)
	return &userApi, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*entity.TinyUser, error) {

	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("users")
	searchQuery := bson.M{
		"email":  email,
		"status": entity.ActiveStatus,
	}

	var tinyUserM TinyUserModel
	err := com.Find(searchQuery).One(&tinyUserM)
	if err != nil {
		return nil, err
	}

	userApi := toEntityTinyUser(tinyUserM)
	return &userApi, nil
}

func (r *UserRepository) GetUserByRegistrationNumber(registrationNumber string) (*entity.TinyUser, error) {

	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("users")
	searchQuery := bson.M{
		"registration_number": registrationNumber,
		"status":              entity.ActiveStatus,
	}

	var tinyUserM TinyUserModel
	err := com.Find(searchQuery).One(&tinyUserM)
	if err != nil {
		return nil, err
	}

	userApi := toEntityTinyUser(tinyUserM)
	return &userApi, nil
}

func (r *UserRepository) CreateUser(newUser entity.UserPayload) error {

	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("users")

	userM := toUserPayloadModel(newUser)
	err := com.Insert(&userM)
	if err != nil {
		return err
	}

	return nil
}
