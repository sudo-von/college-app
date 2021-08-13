package mongo

import (
	"errors"

	"freelancer/college-app/go/entity"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UniversityModel struct {
	ID             bson.ObjectId `bson:"_id"`
	Name           string        `bson:"name"`
	ProfilePicture string        `bson:"profile_picture"`
}

type UniversityPayloadModel struct {
	ID             bson.ObjectId `bson:"_id"`
	Name           string        `bson:"name"`
	ProfilePicture string        `bson:"profile_picture"`
}

func toEntityUniversity(university UniversityModel) entity.University {
	return entity.University{
		ID:             university.ID.Hex(),
		Name:           university.Name,
		ProfilePicture: university.ProfilePicture,
	}
}

type UniversityRepository struct {
	Session      *mgo.Session
	DatabaseName string
}

func NewUniversityRepository(repository *Repository) *UniversityRepository {
	return &UniversityRepository{
		Session:      repository.Session,
		DatabaseName: repository.DatabaseName,
	}
}

func (r *UniversityRepository) GetUniversityByID(universityID string) (*entity.University, error) {

	if !bson.IsObjectIdHex(universityID) {
		return nil, errors.New("given university_id is not a valid hex")
	}

	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("universities")

	var universityM UniversityModel
	err := com.FindId(bson.ObjectIdHex(universityID)).One(&universityM)
	if err != nil {
		return nil, err
	}

	userApi := toEntityUniversity(universityM)
	return &userApi, nil
}
