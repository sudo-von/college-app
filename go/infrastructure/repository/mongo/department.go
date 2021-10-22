package mongo

import (
	"time"

	"freelancer/college-app/go/entity"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DepartmentModel struct {
	ID           bson.ObjectId   `bson:"_id"`
	User         BasicUser       `bson:"user"`
	University   UniversityModel `bson:"university"`
	Description  string          `bson:"description"`
	Street       string          `bson:"street"`
	Neighborhood string          `bson:"neighborhood"`
	Cost         float32         `bson:"cost"`
	Status       string          `bson:"status"`
	Available    bool            `bson:"available"`
	CreationDate time.Time       `bson:"creation_date"`
}

func toEntityDepartment(department DepartmentModel) entity.Department {

	user := entity.BasicUser{
		ID:    department.User.ID.Hex(),
		Name:  department.User.Name,
		Email: department.User.Email,
	}

	return entity.Department{
		ID:           department.ID.Hex(),
		User:         user,
		UniversityID: department.University.ID.Hex(),
		Description:  department.Description,
		Street:       department.Street,
		Neighborhood: department.Neighborhood,
		Cost:         department.Cost,
		Available:    department.Available,
		Status:       department.Status,
		CreationDate: department.CreationDate,
	}
}

type DepartmentPayloadModel struct {
	ID           bson.ObjectId `bson:"_id"`
	UserID       bson.ObjectId `bson:"user_id"`
	UniversityID bson.ObjectId `bson:"university_id"`
	Description  string        `bson:"description"`
	Street       string        `bson:"street"`
	Neighborhood string        `bson:"neighborhood"`
	Cost         float32       `bson:"cost"`
	Available    bool          `bson:"available"`
	Status       string        `bson:"status"`
	CreationDate time.Time     `bson:"creation_date"`
}

func toDepartmentPayloadModel(department entity.DepartmentPayload) DepartmentPayloadModel {

	var departmentID bson.ObjectId
	if department.ID != "" {
		departmentID = bson.ObjectIdHex(department.ID)
	} else {
		departmentID = bson.NewObjectId()
	}

	var userID bson.ObjectId
	if department.UserID != "" {
		userID = bson.ObjectIdHex(department.UserID)
	} else {
		userID = bson.NewObjectId()
	}

	var universityID bson.ObjectId
	if department.UniversityID != "" {
		universityID = bson.ObjectIdHex(department.UniversityID)
	} else {
		universityID = bson.NewObjectId()
	}

	return DepartmentPayloadModel{
		ID:           departmentID,
		UserID:       userID,
		UniversityID: universityID,
		Description:  department.Description,
		Street:       department.Street,
		Neighborhood: department.Neighborhood,
		Cost:         department.Cost,
		Available:    department.Available,
		Status:       department.Status,
		CreationDate: department.CreationDate,
	}
}

type DepartmentRepository struct {
	Session      *mgo.Session
	DatabaseName string
}

func NewDepartmentRepository(repository *Repository) *DepartmentRepository {
	return &DepartmentRepository{
		Session:      repository.Session,
		DatabaseName: repository.DatabaseName,
	}
}

func (r *DepartmentRepository) GetDepartmentByID(departmentID string) (*entity.Department, error) {

	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("departments")

	searchQuery := bson.M{
		"_id":         bson.ObjectIdHex(departmentID),
		"available":   true,
		"user.status": entity.ActiveStatus,
		"status":      entity.ActiveStatus,
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
				"localField":   "university_id",
				"foreignField": "_id",
				"as":           "university",
			},
		},
		{"$unwind": "$university"},
		{"$match": searchQuery},
	}

	var departmentM DepartmentModel
	pipe := com.Pipe(pipes)
	err := pipe.One(&departmentM)
	if err != nil {
		return nil, err
	}

	department := toEntityDepartment(departmentM)
	return &department, nil
}

func (r *DepartmentRepository) GetDepartments(universityID string, departmentFilters entity.DepartmentFilters) ([]entity.Department, *int, error) {

	session := r.Session.Copy()
	defer session.Close()
	con := session.DB(r.DatabaseName).C("departments")

	searchQuery := bson.M{
		"available":     true,
		"user.status":   entity.ActiveStatus,
		"university_id": bson.ObjectIdHex(universityID),
		"status":        entity.ActiveStatus,
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
				"localField":   "university_id",
				"foreignField": "_id",
				"as":           "university",
			},
		},
		{"$unwind": "$university"},
		{"$match": searchQuery},
	}

	var departmentM []DepartmentModel
	pipe := con.Pipe(pipes)
	err := pipe.All(&departmentM)
	if err != nil {
		return nil, nil, err
	}
	total := len(departmentM)

	departments := make([]entity.Department, 0)
	for _, d := range departmentM {
		departments = append(departments, toEntityDepartment(d))
	}

	return departments, &total, nil
}

func (r *DepartmentRepository) CreateDepartment(newDepartment entity.DepartmentPayload) error {

	session := r.Session.Copy()
	defer session.Close()
	con := session.DB(r.DatabaseName).C("departments")

	departmentM := toDepartmentPayloadModel(newDepartment)
	err := con.Insert(&departmentM)
	if err != nil {
		return err
	}

	return nil
}

func (r *DepartmentRepository) UpdateDepartment(updatedDepartment entity.DepartmentPayload) error {

	session := r.Session.Copy()
	defer session.Close()
	con := session.DB(r.DatabaseName).C("departments")

	searchQuery := bson.M{
		"_id": bson.ObjectIdHex(updatedDepartment.ID),
	}

	departmentM := toDepartmentPayloadModel(updatedDepartment)
	err := con.Update(searchQuery, &departmentM)
	if err != nil {
		return err
	}

	return nil
}
