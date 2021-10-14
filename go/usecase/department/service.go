package department

import (
	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/usecase/user"
)

type Service struct {
	departmentRepository DepartmentRepository
	userRepository       UserRepository
}

func NewService(departmentRepository DepartmentRepository, userRepository user.UserRepository) *Service {
	return &Service{
		departmentRepository,
		userRepository,
	}
}

func (s Service) GetDepartments(userID string, departmentFilters entity.DepartmentFilters) ([]entity.Department, *int, error) {

	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		if err.Error() == "not found" {
			return nil, nil, entity.NewErrorNotFound(err, presenter.ErrUserNotFound)
		}
		return nil, nil, entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	departments, total, err := s.departmentRepository.GetDepartments(user.University.ID, departmentFilters)
	if err != nil {
		return nil, nil, entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	return departments, total, nil
}

func (s Service) CreateDepartment(userID string, newDepartment entity.DepartmentPayload) error {

	// Gets university id from the user.
	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		if err.Error() == "not found" {
			return entity.NewErrorNotFound(err, presenter.ErrUserNotFound)
		}
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}
	newDepartment.UniversityID = user.University.ID

	err = s.departmentRepository.CreateDepartment(newDepartment)
	if err != nil {
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}
	return nil
}
