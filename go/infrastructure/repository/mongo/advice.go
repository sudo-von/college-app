package mongo

import (
	"fmt"
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

	university := entity.University{
		ID:             advice.User.University.ID.Hex(),
		Name:           advice.User.University.Name,
		ProfilePicture: advice.User.University.ProfilePicture,
	}

	user := entity.TinyUser{
		ID:                 advice.User.ID.Hex(),
		Name:               advice.User.Name,
		BirthDate:          advice.User.BirthDate,
		Email:              advice.User.Email,
		RegistrationNumber: advice.User.RegistrationNumber,
		University:         university,
	}

	return entity.Advice{
		ID:         advice.ID.Hex(),
		User:       user,
		Subject:    advice.Subject,
		Classroom:  advice.Classroom,
		AdviceDate: advice.AdviceDate,
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

	fmt.Println(universityID, adviceFilters)

	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("advices")

	searchQuery := bson.M{
		"status":              entity.ActiveStatus,
		"user.status":         entity.ActiveStatus,
		"user.university._id": bson.ObjectIdHex(universityID),
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
		{
			"$lookup": bson.M{
				"from":         "universities",
				"localField":   "user.university_id",
				"foreignField": "_id",
				"as":           "user.university",
			},
		},
		{"$unwind": "$user.university"},
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
