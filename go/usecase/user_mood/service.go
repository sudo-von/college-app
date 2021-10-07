package user_mood

import (
	"errors"
	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/entity"
	"time"
)

type Service struct {
	userRepository     UserRepository
	userMoodRepository UserMoodRepository
}

func NewService(userRepository UserRepository, userMoodRepository UserMoodRepository) *Service {
	return &Service{
		userRepository,
		userMoodRepository,
	}
}

func (s Service) GetUserMoodByUserID(userID, requestedUserID string, userMoodFilters entity.UserMoodFilters) (*entity.UserMood, error) {

	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		return nil, entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	err = user.ValidateRequestedUser(requestedUserID)
	if err != nil {
		return nil, entity.NewErrorUnauthorized(err, presenter.ErrInsufficientPermissions)
	}

	userMood, err := s.userMoodRepository.GetUserMoodByUserID(requestedUserID, userMoodFilters)
	if err != nil {
		if err.Error() == "not found" {
			return nil, entity.NewErrorNotFound(err, presenter.ErrUserMoodNotFound)
		}
		return nil, entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	return userMood, nil
}

func (s Service) CreateUserMood(userID, requestedUserID string, newUserMood entity.UserMoodPayload) error {

	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	err = user.ValidateRequestedUser(requestedUserID)
	if err != nil {
		return entity.NewErrorUnauthorized(err, presenter.ErrInsufficientPermissions)
	}

	// Checks if given mood is in the valid range.
	err = newUserMood.ValidateMood()
	if err != nil {
		return entity.NewErrorConflict(err, presenter.ErrInvMood)
	}
	// Checks if the user mood has already been registered today.
	currentDate := time.Now().In(time.Local)
	userMoodFilters := entity.UserMoodFilters{
		CreationDate: &currentDate,
	}
	_, err = s.GetUserMoodByUserID(userID, requestedUserID, userMoodFilters)
	if err != nil && err.Error() != "not found" {
		return err
	} else if err == nil {
		return entity.NewErrorConflict(errors.New("user mood has already been registered today"), presenter.ErrUserMoodAlreadyRegistered)
	}
	// Stores new user mood.
	err = s.userMoodRepository.CreateUserMood(newUserMood)
	if err != nil {
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}
	return nil
}
