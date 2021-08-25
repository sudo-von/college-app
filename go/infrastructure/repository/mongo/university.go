package mongo

import (
	"errors"

	"freelancer/college-app/go/entity"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TinyUniversityModel struct {
	ID             bson.ObjectId `bson:"_id"`
	Name           string        `bson:"name"`
	ProfilePicture string        `bson:"profile_picture"`
}

func toTinyUniversityEntity(university TinyUniversityModel) entity.TinyUniversity {
	return entity.TinyUniversity{
		ID:             university.ID.Hex(),
		Name:           university.Name,
		ProfilePicture: university.ProfilePicture,
	}
}

type UniversityModel struct {
	ID             bson.ObjectId   `bson:"_id"`
	Name           string          `bson:"name"`
	ProfilePicture string          `bson:"profile_picture"`
	Classrooms     []ClasroomModel `bson:"classrooms"`
}

type UniversityPayloadModel struct {
	ID             bson.ObjectId `bson:"_id"`
	Name           string        `bson:"name"`
	ProfilePicture string        `bson:"profile_picture"`
}

func toUniversityEntity(university UniversityModel) entity.University {

	classrooms := make([]entity.Classroom, 0)
	for _, c := range university.Classrooms {
		classrooms = append(classrooms, entity.Classroom{
			ID:   c.ID.Hex(),
			Name: c.Name,
		})
	}

	return entity.University{
		ID:             university.ID.Hex(),
		Name:           university.Name,
		ProfilePicture: university.ProfilePicture,
		Classrooms:     classrooms,
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

func (r *UniversityRepository) GetTinyUniversities() ([]entity.TinyUniversity, *int, error) {

	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("universities")

	var universitiesM []TinyUniversityModel
	err := com.Find(bson.M{}).All(&universitiesM)
	if err != nil {
		return nil, nil, err
	}

	total, err := com.Find(bson.M{}).Count()
	if err != nil {
		return nil, nil, err
	}

	universities := make([]entity.TinyUniversity, 0)
	for _, university := range universitiesM {
		universities = append(universities, toTinyUniversityEntity(university))
	}

	return universities, &total, nil
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

	userApi := toUniversityEntity(universityM)
	return &userApi, nil
}
