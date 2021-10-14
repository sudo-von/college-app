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

type DepartmentReader interface {
	GetDepartments(universityID string, departmentFilters entity.DepartmentFilters) ([]entity.Department, *int, error)
}

type DepartmentWriter interface {
	CreateDepartment(newDepartment entity.DepartmentPayload) error
}

type DepartmentRepository interface {
	DepartmentReader
	DepartmentWriter
}
