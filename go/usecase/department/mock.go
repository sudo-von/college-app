package department

import (
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/pkg/search"
	"time"
)

var (
	fakeDepartmentIDS = []string{"6167698bd0187e53d4789212", "6167698bd0187e53d4789213"}
	fakeUniversityIDS = []string{"615c09f7309d7ded48c7a049", "615c0794590f4315693633a6"}
)

type DepartmentReaderMock struct{}

func (a DepartmentReaderMock) GetDepartmentByID(departmentID string) (*entity.Department, error) {

	err := search.FindInSlice(departmentID, fakeDepartmentIDS)
	if err != nil {
		return nil, err
	}

	department := &entity.Department{
		ID: departmentID,
		User: entity.BasicUser{
			ID: "615c09f7309d7ded48c7a053",
		},
		UniversityID: "615c09f7309d7ded48c7a049",
		Description:  "fake description",
		Street:       "fake street",
		Neighborhood: "fake neighborhood",
		Cost:         1000,
		Available:    true,
		CreationDate: time.Date(2021, 01, 01, 0, 0, 0, 0, time.Local),
	}
	return department, nil
}

func (a DepartmentReaderMock) GetDepartments(universityID string, departmentFilters entity.DepartmentFilters) ([]entity.Department, *int, error) {

	err := search.FindInSlice(universityID, fakeUniversityIDS)
	if err != nil {
		return nil, nil, err
	}

	departments := []entity.Department{
		{
			ID: "6167698bd0187e53d4789212",
			User: entity.BasicUser{
				ID: "615c09f7309d7ded48c7a053",
			},
			UniversityID: universityID,
			Description:  "fake description",
			Street:       "fake street",
			Neighborhood: "fake neighborhood",
			Cost:         1000,
			Available:    true,
			CreationDate: time.Date(2021, 01, 01, 0, 0, 0, 0, time.Local),
		},
	}
	total := len(departments)
	return departments, &total, nil
}

type DepartmentWriterMock struct{}

func (a DepartmentWriterMock) CreateDepartment(newDepartment entity.DepartmentPayload) error {
	return nil
}

func (a DepartmentWriterMock) UpdateDepartment(updateDepartment entity.DepartmentPayload) error {
	return nil
}

type DepartmentRepositoryMock struct {
	DepartmentReaderMock
	DepartmentWriterMock
}
