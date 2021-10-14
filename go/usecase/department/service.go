package department

import (
	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/usecase/university"
	"freelancer/college-app/go/usecase/user"
)

type Service struct {
	departmentRepository DepartmentRepository
	userRepository       UserRepository
	universityRepository UniversityRepository
}

func NewService(departmentRepository DepartmentRepository, userRepository user.UserRepository, universityRepository university.UniversityRepository) *Service {
	return &Service{
		departmentRepository,
		userRepository,
		universityRepository,
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
