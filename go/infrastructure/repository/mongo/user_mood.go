package mongo

import (
	"errors"
	"time"

	"freelancer/college-app/go/entity"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserMoodModel struct {
	ID           bson.ObjectId `bson:"_id"`
	UserID       bson.ObjectId `bson:"user_id"`
	Mood         float64       `bson:"mood"`
	CreationDate time.Time     `bson:"creation_date"`
}

type UserMoodPayload struct {
	ID           bson.ObjectId `bson:"_id"`
	UserID       bson.ObjectId `bson:"user_id"`
	Mood         float64       `bson:"mood"`
	CreationDate time.Time     `bson:"creation_date"`
}

func toEntityUserMood(userMood UserMoodModel) entity.UserMood {
	return entity.UserMood{
		ID:           userMood.ID.Hex(),
		UserID:       userMood.UserID.Hex(),
		Mood:         userMood.Mood,
		CreationDate: userMood.CreationDate,
	}
}

func toUserMoodPayloadModel(userMoodPayload entity.UserMoodPayload) UserMoodPayload {

	var userMoodPayloadID bson.ObjectId
	if userMoodPayload.ID != "" {
		userMoodPayloadID = bson.ObjectIdHex(userMoodPayload.ID)
	} else {
		userMoodPayloadID = bson.NewObjectId()
	}

	var userID bson.ObjectId
	if userMoodPayload.UserID != "" {
		userID = bson.ObjectIdHex(userMoodPayload.UserID)
	} else {
		userID = bson.NewObjectId()
	}

	return UserMoodPayload{
		ID:           userMoodPayloadID,
		UserID:       userID,
		Mood:         userMoodPayload.Mood,
		CreationDate: userMoodPayload.CreationDate,
	}
}

type UserMoodRepository struct {
	Session      *mgo.Session
	DatabaseName string
}

func NewUserMoodRepository(repository *Repository) *UserMoodRepository {
	return &UserMoodRepository{
		Session:      repository.Session,
		DatabaseName: repository.DatabaseName,
	}
}

func (r *UserMoodRepository) GetUserMoodByUserID(userID string, userMoodFilters entity.UserMoodFilters) (*entity.UserMood, error) {

	if !bson.IsObjectIdHex(userID) {
		return nil, errors.New("given user_id is not a valid hex")
	}

	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("users_mood")

	searchQuery := bson.M{
		"user_id": bson.ObjectIdHex(userID),
	}

	currentDate := time.Now()
	initialDate := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 0, 0, 0, 0, time.Local)
	if userMoodFilters.CreationDate != nil {
		initialDate = time.Date(userMoodFilters.CreationDate.Year(), userMoodFilters.CreationDate.Month(), userMoodFilters.CreationDate.Day(), 0, 0, 0, 0, time.Local)
	}
	searchQuery["creation_date"] = bson.M{
		"$gte": initialDate,
		"$lt":  initialDate.AddDate(0, 0, 1),
	}

	var userMoodM UserMoodModel
	err := com.Find(searchQuery).One(&userMoodM)
	if err != nil {
		return nil, err
	}

	userMood := toEntityUserMood(userMoodM)
	return &userMood, nil
}

func (r *UserMoodRepository) CreateUserMood(newUserMood entity.UserMoodPayload) error {

	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("users_mood")

	userMoodM := toUserMoodPayloadModel(newUserMood)
	err := com.Insert(&userMoodM)
	if err != nil {
		return err
	}

	return nil
}
