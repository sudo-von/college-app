package mongo

import (
	"time"

	"freelancer/college-app/go/entity"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AdviceModel struct {
	ID         bson.ObjectId `bson:"_id"`
	User       TinyUserModel `bson:"user"`
	Subject    string        `bson:"subject"`
	Classroom  int           `bson:"classroom"`
	AdviceDate time.Time     `bson:"advice_date"`
}

func toEntityAdvice(advice AdviceModel) entity.Advice {

	user := entity.TinyUser{
		ID:                 advice.User.ID.Hex(),
		Name:               advice.User.Name,
		BirthDate:          advice.User.BirthDate,
		Email:              advice.User.Email,
		RegistrationNumber: advice.User.RegistrationNumber,
	}

	return entity.Advice{
		ID:         advice.ID.Hex(),
		User:       user,
		Subject:    advice.Subject,
		Classroom:  advice.Classroom,
		AdviceDate: advice.AdviceDate,
	}
}

type AdvicePayloadModel struct {
	ID             bson.ObjectId `bson:"_id"`
	UserID         bson.ObjectId `bson:"user_id"`
	UniversityID   bson.ObjectId `bson:"university_id"`
	Subject        string        `bson:"subject"`
	AdviceDate     time.Time     `bson:"advice_date"`
	Classroom      int           `bson:"classroom"`
	StudentsNumber int           `bson:"students_number"`
	Status         string        `bson:"status"`
	CreationDate   time.Time     `bson:"creation_date"`
}

func toAdvicePayloadModel(advice entity.AdvicePayload) AdvicePayloadModel {

	var adviceID bson.ObjectId
	if advice.ID != "" {
		adviceID = bson.ObjectIdHex(advice.ID)
	} else {
		adviceID = bson.NewObjectId()
	}

	var userID bson.ObjectId
	if advice.UserID != "" {
		userID = bson.ObjectIdHex(advice.UserID)
	} else {
		userID = bson.NewObjectId()
	}

	var universityID bson.ObjectId
	if advice.UniversityID != "" {
		universityID = bson.ObjectIdHex(advice.UniversityID)
	} else {
		universityID = bson.NewObjectId()
	}

	return AdvicePayloadModel{
		ID:             adviceID,
		UserID:         userID,
		UniversityID:   universityID,
		Subject:        advice.Subject,
		AdviceDate:     advice.AdviceDate,
		Classroom:      advice.Classroom,
		StudentsNumber: advice.StudentsNumber,
		Status:         advice.Status,
		CreationDate:   advice.CreationDate,
	}
}

type AdviceRepository struct {
	Session      *mgo.Session
	DatabaseName string
}

func NewAdviceRepository(repository *Repository) *AdviceRepository {
	return &AdviceRepository{
		Session:      repository.Session,
		DatabaseName: repository.DatabaseName,
	}
}

func (r *AdviceRepository) GetAdvices(universityID string, adviceFilters entity.AdviceFilters) ([]entity.Advice, *int, error) {

	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("advices")

	searchQuery := bson.M{
		"status":        entity.ActiveStatus,
		"user.status":   entity.ActiveStatus,
		"university_id": bson.ObjectIdHex(universityID),
		"advice_date": bson.M{
			"$gte": adviceFilters.AdviceDate.In(time.Local),
		},
	}

	pipes := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "users",
				"localField":   "user_id",
				"foreignField": "_id",
				"as":           "user",
			},
		},
		{"$unwind": "$user"},
		{"$match": searchQuery},
	}

	var advicesM []AdviceModel
	pipe := com.Pipe(pipes)
	err := pipe.All(&advicesM)
	if err != nil {
		return nil, nil, err
	}

	total := len(advicesM)
	consultancies := make([]entity.Advice, 0)
	for _, advice := range advicesM {
		consultancies = append(consultancies, toEntityAdvice(advice))
	}

	return consultancies, &total, nil
}

func (r *AdviceRepository) CreateAdvice(newAdvice entity.AdvicePayload) error {

	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("advices")

	adviceM := toAdvicePayloadModel(newAdvice)
	err := com.Insert(&adviceM)
	if err != nil {
		return err
	}

	return nil
}
