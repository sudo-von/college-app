package department

import (
	"freelancer/college-app/go/entity"
)

type UserReader interface {
	GetUserByID(userID string) (*entity.User, error)
}

type UserRepository interface {
	UserReader
}

type UniversityReader interface {
	GetUniversityByID(universityID string) (*entity.University, error)
}

type UniversityRepository interface {
	UniversityReader
}

type DepartmentReader interface {
	GetDepartments(universityID string, departmentFilters entity.DepartmentFilters) ([]entity.Department, *int, error)
}

type DepartmentRepository interface {
	DepartmentReader
}
