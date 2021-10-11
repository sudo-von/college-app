package user

import (
	"errors"
	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/entity"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	userRepository       UserRepository
	universityRepository UniversityRepository
}

func NewService(userRepository UserRepository, universityRepository UniversityRepository) *Service {
	return &Service{
		userRepository:       userRepository,
		universityRepository: universityRepository,
	}
}

func (s Service) GetTinyUserByID(userID, requestedUserID string) (*entity.TinyUser, error) {

	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		if err.Error() == "not found" {
			return nil, entity.NewErrorNotFound(err, presenter.ErrUserNotFound)
		}
		return nil, entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	err = user.ValidateRequestedUser(requestedUserID)
	if err != nil {
		return nil, entity.NewErrorUnauthorized(err, presenter.ErrInsufficientPermissions)
	}

	requestedUser, err := s.userRepository.GetTinyUserByID(requestedUserID)
	if err != nil {
		if err.Error() == "not found" {
			return nil, entity.NewErrorNotFound(err, presenter.ErrUserNotFound)
		}
		return nil, entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}
	return requestedUser, nil

}

func (s Service) CreateUser(newUser entity.UserPayload) error {

	// Check if email has a correct format .
	err := checkmail.ValidateFormat(newUser.Email)
	if err != nil {
		return entity.NewErrorConflict(errors.New("invalid email"), presenter.ErrInvUserEmail)
	}
	// Checks if registration number has a correct format.
	err = newUser.ValidateRegistrationNumberDigits()
	if err != nil {
		return entity.NewErrorConflict(err, presenter.ErrInvUserRegistrationNumberDigits)
	}
	err = newUser.ValidateRegistrationNumberLength()
	if err != nil {
		return entity.NewErrorConflict(err, presenter.ErrInvUserRegistrationNumberLength)
	}
	// Checks if email is already in use.
	_, err = s.userRepository.GetTinyUserByEmail(newUser.Email)
	if err != nil && err.Error() != "not found" {
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	} else if err == nil {
		return entity.NewErrorConflict(errors.New("email already in use"), presenter.ErrUserEmailAlreadyRegistered)
	}
	// Checks if registration number is already in use.
	_, err = s.userRepository.GetTinyUserByRegistrationNumber(newUser.RegistrationNumber)
	if err != nil && err.Error() != "not found" {
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	} else if err == nil {
		return entity.NewErrorConflict(errors.New("registration number already in use"), presenter.ErrUserRegistrationNumberAlreadyRegistered)
	}
	// Checks if university exists.
	_, err = s.universityRepository.GetUniversityByID(newUser.UniversityID)
	if err != nil {
		if err.Error() == "not found" {
			return entity.NewErrorNotFound(err, presenter.ErrUniversityNotFound)
		}
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}
	// Encrypt password.
	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}
	newUser.Password = string(hash)
	// Stores newUser.
	err = s.userRepository.CreateUser(newUser)
	if err != nil {
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}
	return nil
}

func (s Service) UpdateTinyUser(userID, requestedUserID string, newUser entity.UpdateUserPayload) error {

	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	err = user.ValidateRequestedUser(requestedUserID)
	if err != nil {
		return entity.NewErrorUnauthorized(err, presenter.ErrInsufficientPermissions)
	}

	requestedUser, err := s.userRepository.GetUserByID(requestedUserID)
	if err != nil {
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	if requestedUser.Email != newUser.Email {
		err := checkmail.ValidateFormat(newUser.Email)
		if err != nil {
			return entity.NewErrorConflict(errors.New("invalid email"), presenter.ErrInvUserEmail)
		}
		_, err = s.userRepository.GetTinyUserByEmail(newUser.Email)
		if err != nil && err.Error() != "not found" {
			return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
		} else if err == nil {
			return entity.NewErrorConflict(errors.New("email already in use"), presenter.ErrUserEmailAlreadyRegistered)
		}
	}

	if requestedUser.RegistrationNumber != newUser.RegistrationNumber {
		err = newUser.ValidateRegistrationNumberDigits()
		if err != nil {
			return entity.NewErrorConflict(err, presenter.ErrInvUserRegistrationNumberDigits)
		}
		err = newUser.ValidateRegistrationNumberLength()
		if err != nil {
			return entity.NewErrorConflict(err, presenter.ErrInvUserRegistrationNumberLength)
		}
		_, err = s.userRepository.GetTinyUserByRegistrationNumber(newUser.RegistrationNumber)
		if err != nil && err.Error() != "not found" {
			return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
		} else if err == nil {
			return entity.NewErrorConflict(errors.New("registration number already in use"), presenter.ErrUserRegistrationNumberAlreadyRegistered)
		}
	}

	// Creates a new user payload and then replaces the old one.
	updatedUser := entity.UserPayload{
		ID:                 requestedUser.ID,
		Name:               newUser.Name,
		BirthDate:          newUser.BirthDate,
		Email:              newUser.Email,
		RegistrationNumber: newUser.RegistrationNumber,
		Password:           requestedUser.Password,
		UniversityID:       requestedUser.University.ID,
		Status:             requestedUser.Status,
		CreationDate:       requestedUser.CreationDate,
	}

	err = s.userRepository.UpdateUser(updatedUser)
	if err != nil {
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}
	return nil
}

func (s Service) AuthenticateUser(email, password string) (*entity.User, error) {

	user, err := s.userRepository.GetUserByEmail(email)
	if err != nil {
		if err.Error() == "not found" {
			return nil, entity.NewErrorNotFound(err, presenter.ErrUserNotFound)
		}
		return nil, entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, entity.NewErrorUnauthorized(errors.New("invalid credentials"), presenter.ErrInvCredentials)
	}
	return user, nil
}
