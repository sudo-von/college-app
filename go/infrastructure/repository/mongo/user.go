package mongo

import (
	"errors"
	"time"

	"freelancer/college-app/go/entity"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserModel struct {
	ID                 bson.ObjectId   `bson:"_id"`
	Name               string          `bson:"name"`
	BirthDate          time.Time       `bson:"birth_date"`
	Email              string          `bson:"email"`
	Password           string          `bson:"password"`
	RegistrationNumber string          `bson:"registration_number"`
	Status             string          `bson:"status"`
	Role               string          `bson:"role"`
	University         UniversityModel `bson:"university"`
	CreationDate       time.Time       `bson:"creation_date"`
}

type BasicUser struct {
	ID    bson.ObjectId `bson:"_id"`
	Name  string        `bson:"name"`
	Email string        `bson:"email"`
}

type TinyUserModel struct {
	ID                 bson.ObjectId   `bson:"_id"`
	Name               string          `bson:"name"`
	BirthDate          time.Time       `bson:"birth_date"`
	Email              string          `bson:"email"`
	RegistrationNumber string          `bson:"registration_number"`
	University         UniversityModel `bson:"university"`
}

type UserPayloadModel struct {
	ID                 bson.ObjectId `bson:"_id"`
	Name               string        `bson:"name"`
	BirthDate          time.Time     `bson:"birth_date"`
	Email              string        `bson:"email"`
	RegistrationNumber string        `bson:"registration_number"`
	Password           string        `bson:"password"`
	Status             string        `bson:"status"`
	Role               string        `bson:"role"`
	UniversityID       bson.ObjectId `bson:"university_id"`
	CreationDate       time.Time     `bson:"creation_date"`
}

func toUserPayloadModel(userPayload entity.UserPayload) UserPayloadModel {

	var userID bson.ObjectId
	if userPayload.ID != "" {
		userID = bson.ObjectIdHex(userPayload.ID)
	} else {
		userID = bson.NewObjectId()
	}

	var universityID bson.ObjectId
	if userPayload.UniversityID != "" {
		universityID = bson.ObjectIdHex(userPayload.UniversityID)
	} else {
		universityID = bson.NewObjectId()
	}

	return UserPayloadModel{
		ID:                 userID,
		Name:               userPayload.Name,
		BirthDate:          userPayload.BirthDate,
		Email:              userPayload.Email,
		RegistrationNumber: userPayload.RegistrationNumber,
		Password:           userPayload.Password,
		Status:             userPayload.Status,
		UniversityID:       universityID,
		Role:               userPayload.Role,
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

func toEntityUser(user UserModel) entity.User {

	university := entity.University{
		ID:             user.University.ID.Hex(),
		Name:           user.University.Name,
		ProfilePicture: user.University.ProfilePicture,
	}

	return entity.User{
		ID:                 user.ID.Hex(),
		Name:               user.Name,
		BirthDate:          user.BirthDate,
		Email:              user.Email,
		Password:           user.Password,
		RegistrationNumber: user.RegistrationNumber,
		Status:             user.Status,
		University:         university,
		Role:               user.Role,
		CreationDate:       user.CreationDate,
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

func (r *UserRepository) GetTinyUserByID(userID string) (*entity.TinyUser, error) {

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

	pipes := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "universities",
				"localField":   "university_id",
				"foreignField": "_id",
				"as":           "university",
			},
		},
		{"$unwind": "$university"},
		{"$match": searchQuery},
	}

	var tinyUserM TinyUserModel
	err := com.Pipe(pipes).One(&tinyUserM)
	if err != nil {
		return nil, err
	}

	userApi := toEntityTinyUser(tinyUserM)
	return &userApi, nil
}

func (r *UserRepository) GetTinyUserByEmail(email string) (*entity.TinyUser, error) {

	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("users")
	searchQuery := bson.M{
		"email": email,
	}

	pipes := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "universities",
				"localField":   "university_id",
				"foreignField": "_id",
				"as":           "university",
			},
		},
		{"$unwind": "$university"},
		{"$match": searchQuery},
	}

	var tinyUserM TinyUserModel
	err := com.Pipe(pipes).One(&tinyUserM)
	if err != nil {
		return nil, err
	}

	userApi := toEntityTinyUser(tinyUserM)
	return &userApi, nil
}

func (r *UserRepository) GetTinyUserByRegistrationNumber(registrationNumber string) (*entity.TinyUser, error) {

	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("users")
	searchQuery := bson.M{
		"registration_number": registrationNumber,
	}

	pipes := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "universities",
				"localField":   "university_id",
				"foreignField": "_id",
				"as":           "university",
			},
		},
		{"$unwind": "$university"},
		{"$match": searchQuery},
	}

	var tinyUserM TinyUserModel
	err := com.Pipe(pipes).One(&tinyUserM)
	if err != nil {
		return nil, err
	}

	userApi := toEntityTinyUser(tinyUserM)
	return &userApi, nil
}

func (r *UserRepository) GetUserByID(userID string) (*entity.User, error) {

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

	pipes := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "universities",
				"localField":   "university_id",
				"foreignField": "_id",
				"as":           "university",
			},
		},
		{"$unwind": "$university"},
		{"$match": searchQuery},
	}

	var userM UserModel
	err := com.Pipe(pipes).One(&userM)
	if err != nil {
		return nil, err
	}

	user := toEntityUser(userM)
	return &user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*entity.User, error) {

	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("users")
	searchQuery := bson.M{
		"email":  email,
		"status": entity.ActiveStatus,
	}

	pipes := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "universities",
				"localField":   "university_id",
				"foreignField": "_id",
				"as":           "university",
			},
		},
		{"$unwind": "$university"},
		{"$match": searchQuery},
	}

	var userM UserModel
	err := com.Pipe(pipes).One(&userM)
	if err != nil {
		return nil, err
	}

	user := toEntityUser(userM)
	return &user, nil
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

func (r *UserRepository) UpdateUser(updatedUser entity.UserPayload) error {

	if !bson.IsObjectIdHex(updatedUser.ID) {
		return errors.New("given user_id is not a valid hex")
	}

	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("users")
	searchQuery := bson.M{
		"_id":    bson.ObjectIdHex(updatedUser.ID),
		"status": entity.ActiveStatus,
	}

	userM := toUserPayloadModel(updatedUser)
	err := com.Update(searchQuery, &userM)
	if err != nil {
		return err
	}

	return nil
}
