package mongo

import (
	"time"

	"freelancer/college-app/go/entity"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type SuggestionPayloadModel struct {
	ID           bson.ObjectId `bson:"_id"`
	UserID       bson.ObjectId `bson:"user_id"`
	Suggestion   string        `bson:"suggestion"`
	CreationDate time.Time     `bson:"creation_date"`
}

func toSuggestionPayloadModel(suggestionPayload entity.SuggestionPayload) SuggestionPayloadModel {

	var suggestionID bson.ObjectId
	if suggestionPayload.ID != "" {
		suggestionID = bson.ObjectIdHex(suggestionPayload.ID)
	} else {
		suggestionID = bson.NewObjectId()
	}

	var userID bson.ObjectId
	if suggestionPayload.UserID != "" {
		userID = bson.ObjectIdHex(suggestionPayload.UserID)
	} else {
		userID = bson.NewObjectId()
	}

	return SuggestionPayloadModel{
		ID:           suggestionID,
		UserID:       userID,
		Suggestion:   suggestionPayload.Suggestion,
		CreationDate: suggestionPayload.CreationDate,
	}
}

type SuggestionRepository struct {
	Session      *mgo.Session
	DatabaseName string
}

func NewSuggestionRepository(repository *Repository) *SuggestionRepository {
	return &SuggestionRepository{
		Session:      repository.Session,
		DatabaseName: repository.DatabaseName,
	}
}

func (r *SuggestionRepository) CreateSuggestion(newSuggestion entity.SuggestionPayload) error {

	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("suggestions")

	suggestionM := toSuggestionPayloadModel(newSuggestion)
	err := com.Insert(&suggestionM)
	if err != nil {
		return err
	}

	return nil
}
