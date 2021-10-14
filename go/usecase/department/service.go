package department

import (
	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/entity"
)

type Service struct {
	departmentRepository DepartmentRepository
	userRepository       UserRepository
}

func NewService(departmentRepository DepartmentRepository, userRepository UserRepository) *Service {
	return &Service{
		departmentRepository,
		userRepository,
	}
}

func (s Service) GetDepartmentByID(userID, departmentID string) (*entity.Department, error) {

	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		if err.Error() == "not found" {
			return nil, entity.NewErrorNotFound(err, presenter.ErrUserNotFound)
		}
		return nil, entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	department, err := s.departmentRepository.GetDepartmentByID(departmentID)
	if err != nil {
		if err.Error() == "not found" {
			return nil, entity.NewErrorNotFound(err, presenter.ErrAdvivceNotFound)
		}
		return nil, entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	err = department.ValidateUniversity(user.University.ID)
	if err != nil {
		return nil, entity.NewErrorConflict(err, presenter.ErrInsufficientPermissions)
	}

	return department, nil
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

	err := newDepartment.ValidateCost()
	if err != nil {
		return entity.NewErrorConflict(err, presenter.ErrInvDepCost)
	}
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

func (s Service) UpdateDepartment(userID, departmentID string, departmentPayload entity.UpdateDepartmentPayload) error {

	oldDepartment, err := s.GetDepartmentByID(userID, departmentID)
	if err != nil {
		return err
	}

	err = oldDepartment.ValidateRequestedDepartment(userID)
	if err != nil {
		return entity.NewErrorConflict(err, presenter.ErrInsufficientPermissions)
	}

	if oldDepartment.Cost != departmentPayload.Cost {
		err := departmentPayload.ValidateCost()
		if err != nil {
			return entity.NewErrorConflict(err, presenter.ErrInvDepCost)
		}
	}

	updatedDepartment := entity.DepartmentPayload{
		ID:           oldDepartment.ID,
		UserID:       oldDepartment.User.ID,
		UniversityID: oldDepartment.UniversityID,
		Description:  departmentPayload.Description,
		Street:       departmentPayload.Street,
		Neighborhood: departmentPayload.Neighborhood,
		Cost:         departmentPayload.Cost,
		Available:    departmentPayload.Available,
		Status:       oldDepartment.Status,
		CreationDate: oldDepartment.CreationDate,
	}

	err = s.departmentRepository.UpdateDepartment(updatedDepartment)
	if err != nil {
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}
	return nil
}

func (s Service) DeleteDepartment(userID, departmentID string) error {

	oldDepartment, err := s.GetDepartmentByID(userID, departmentID)
	if err != nil {
		return err
	}

	err = oldDepartment.ValidateRequestedDepartment(userID)
	if err != nil {
		return entity.NewErrorConflict(err, presenter.ErrInsufficientPermissions)
	}

	updatedDepartment := entity.DepartmentPayload{
		ID:           oldDepartment.ID,
		UserID:       oldDepartment.User.ID,
		UniversityID: oldDepartment.UniversityID,
		Description:  oldDepartment.Description,
		Street:       oldDepartment.Street,
		Neighborhood: oldDepartment.Neighborhood,
		Cost:         oldDepartment.Cost,
		Available:    oldDepartment.Available,
		Status:       entity.DeletedStatus,
		CreationDate: oldDepartment.CreationDate,
	}

	err = s.departmentRepository.UpdateDepartment(updatedDepartment)
	if err != nil {
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}
	return nil
}
