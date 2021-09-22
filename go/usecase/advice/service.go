package advice

import (
	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/usecase/university"
	"freelancer/college-app/go/usecase/user"
)

type Service struct {
	adviceRepository     AdviceRepository
	userRepository       user.UserRepository
	universityRepository university.UniversityRepository
}

func NewService(adviceRepository AdviceRepository, userRepository user.UserRepository, universityRepository university.UniversityRepository) *Service {
	return &Service{
		adviceRepository,
		userRepository,
		universityRepository,
	}
}

func (s *Service) GetAdvices(userID string, adviceFilters entity.AdviceFilters) ([]entity.Advice, *int, error) {

	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		if err.Error() == "not found" {
			return nil, nil, entity.NewErrorNotFound(err, presenter.ErrUserNotFound)
		}
		return nil, nil, entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	advices, total, err := s.adviceRepository.GetAdvices(user.University.ID, adviceFilters)
	if err != nil {
		return nil, nil, entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	return advices, total, nil
}

func (s *Service) GetAdviceByID(userID, adviceID string) (*entity.Advice, error) {

	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		if err.Error() == "not found" {
			return nil, entity.NewErrorNotFound(err, presenter.ErrUserNotFound)
		}
		return nil, entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	advice, err := s.adviceRepository.GetAdviceByID(adviceID)
	if err != nil {
		if err.Error() == "not found" {
			return nil, entity.NewErrorNotFound(err, presenter.ErrAdvivceNotFound)
		}
		return nil, entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	// Checks if user's university id is the same as the advice's.
	err = advice.ValidateUniversity(user.University.ID)
	if err != nil {
		return nil, entity.NewErrorConflict(err, presenter.ErrInsufficientPermissions)
	}

	return advice, nil
}

func (s *Service) CreateAdvice(newAdvice entity.AdvicePayload) error {

	// Checks if current date is less than the advice date.
	err := newAdvice.ValidateDate()
	if err != nil {
		return entity.NewErrorConflict(err, presenter.ErrInvAdviceDate)
	}

	// Gets university id from the user.
	user, err := s.userRepository.GetUserByID(newAdvice.UserID)
	if err != nil {
		if err.Error() == "not found" {
			return entity.NewErrorNotFound(err, presenter.ErrUserNotFound)
		}
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}
	newAdvice.UniversityID = user.University.ID

	// Gets given university.
	university, err := s.universityRepository.GetUniversityByID(newAdvice.UniversityID)
	if err != nil {
		if err.Error() == "not found" {
			return entity.NewErrorNotFound(err, presenter.ErrUniversityNotFound)
		}
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	// Checks if the given classroom is a valid classroom from that university
	err = university.ValidateClassroom(newAdvice.ClassroomID)
	if err != nil {
		return entity.NewErrorConflict(err, presenter.ErrInvClassroom)
	}

	// Stores new advice.
	err = s.adviceRepository.CreateAdvice(newAdvice)
	if err != nil {
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}
	return nil
}

func (s *Service) UpdateAdvice(userID, adviceID string, newAdvice entity.UpdateAdvicePayload) error {

	// Gets old advice.
	oldAdvice, err := s.adviceRepository.GetAdviceByID(adviceID)
	if err != nil {
		if err.Error() == "not found" {
			return entity.NewErrorNotFound(err, presenter.ErrAdvivceNotFound)
		}
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	// If the old date is different from the new date then checks if current date is less than the new advice date.
	if oldAdvice.AdviceDate != newAdvice.AdviceDate {
		err := newAdvice.ValidateDate()
		if err != nil {
			return entity.NewErrorConflict(err, presenter.ErrInvAdviceDate)
		}
	}

	// If the old classroom is different from the new classroom then checks if given classroom id is a valid classroom from that university.
	university, err := s.universityRepository.GetUniversityByID(oldAdvice.UniversityID)
	if err != nil {
		if err.Error() == "not found" {
			return entity.NewErrorNotFound(err, presenter.ErrUniversityNotFound)
		}
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	if oldAdvice.Classroom.ID != newAdvice.ClassroomID {
		err = university.ValidateClassroom(newAdvice.ClassroomID)
		if err != nil {
			return entity.NewErrorConflict(err, presenter.ErrInvClassroom)
		}
	}

	// Creates a new advice payload and then replaces the old one.
	updatedAdvice := entity.AdvicePayload{
		ID:                 oldAdvice.ID,
		UserID:             oldAdvice.User.ID,
		UniversityID:       university.ID,
		ClassroomID:        newAdvice.ClassroomID,
		Subject:            newAdvice.Subject,
		AdviceDate:         newAdvice.AdviceDate,
		StudentsWillAttend: oldAdvice.StudentsWillAttend,
		Status:             oldAdvice.Status,
		CreationDate:       oldAdvice.CreationDate,
	}

	// Updates new advice.
	err = s.adviceRepository.UpdateAdvice(updatedAdvice)
	if err != nil {
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}
	return nil
}

func (s *Service) DeleteAdvice(userID, adviceID string) error {

	// Gets old advice.
	oldAdvice, err := s.adviceRepository.GetAdviceByID(adviceID)
	if err != nil {
		if err.Error() == "not found" {
			return entity.NewErrorNotFound(err, presenter.ErrAdvivceNotFound)
		}
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	// Creates a new advice payload and then replaces the old one but now it will have the deleted status.
	updatedAdvice := entity.AdvicePayload{
		ID:                 oldAdvice.ID,
		UserID:             oldAdvice.User.ID,
		ClassroomID:        oldAdvice.Classroom.ID,
		UniversityID:       oldAdvice.UniversityID,
		Subject:            oldAdvice.Subject,
		AdviceDate:         oldAdvice.AdviceDate,
		StudentsWillAttend: oldAdvice.StudentsWillAttend,
		Status:             entity.DeletedStatus,
		CreationDate:       oldAdvice.CreationDate,
	}

	// Stores new advice.
	err = s.adviceRepository.UpdateAdvice(updatedAdvice)
	if err != nil {
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}
	return nil
}
