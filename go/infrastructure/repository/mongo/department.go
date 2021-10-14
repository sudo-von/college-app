package mongo

import (
	"time"

	"freelancer/college-app/go/entity"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DepartmentModel struct {
	ID           bson.ObjectId `bson:"_id"`
	User         BasicUser     `bson:"user"`
	Description  string        `bson:"description"`
	Street       string        `bson:"street"`
	Neighborhood string        `bson:"neighborhood"`
	Cost         float32       `bson:"cost"`
	Available    bool          `bson:"available"`
	CreationDate time.Time     `bson:"creation_date"`
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
		Description:  department.Description,
		Street:       department.Street,
		Neighborhood: department.Neighborhood,
		Cost:         department.Cost,
		Available:    department.Available,
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

func (r *DepartmentRepository) GetDepartments(universityID string, departmentFilters entity.DepartmentFilters) ([]entity.Department, *int, error) {

	session := r.Session.Copy()
	defer session.Close()
	con := session.DB(r.DatabaseName).C("departments")

	searchQuery := bson.M{
		"available":     true,
		"user.status":   entity.ActiveStatus,
		"university_id": bson.ObjectIdHex(universityID),
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
