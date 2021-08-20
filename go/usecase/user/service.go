package user

import (
	"errors"
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/usecase/university"
	"strings"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	userRepository       UserRepository
	universityRepository university.UniversityRepository
}

func NewService(userRepository UserRepository, universityRepository university.UniversityRepository) *Service {
	return &Service{
		userRepository:       userRepository,
		universityRepository: universityRepository,
	}
}

func (s *Service) GetTinyUserByID(userID string) (*entity.TinyUser, error) {

	user, err := s.userRepository.GetTinyUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (s *Service) CreateUser(newUser entity.UserPayload) error {

	// Check if the email is valid.
	err := checkmail.ValidateFormat(newUser.Email)
	if err != nil {
		return errors.New("invalid email")
	}
	// Check if the registration number is valid.
	if len(strings.Replace(newUser.RegistrationNumber, " ", "", -1)) != 8 {
		return errors.New("invalid registration_number")
	}
	// Check if email is already in use.
	_, err = s.userRepository.GetTinyUserByEmail(newUser.Email)
	if err == nil {
		return errors.New("email already in use")
	}
	// Check if registration number is already in use.
	_, err = s.userRepository.GetTinyUserByRegistrationNumber(newUser.RegistrationNumber)
	if err == nil {
		return errors.New("registration number already in use")
	}
	// Check if university exists.
	_, err = s.universityRepository.GetUniversityByID(newUser.UniversityID)
	if err != nil {
		return err
	}
	// Encrypt password.
	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser.Password = string(hash)
	// Stores newUser.
	err = s.userRepository.CreateUser(newUser)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateTinyUser(userID string, newUser entity.UpdateUserPayload) error {

	// Checks if the new email is valid.
	err := checkmail.ValidateFormat(newUser.Email)
	if err != nil {
		return errors.New("invalid email")
	}
	// Check if the new registration number is valid.
	if len(strings.Replace(newUser.RegistrationNumber, " ", "", -1)) != 8 {
		return errors.New("invalid registration_number")
	}
	// Gets old user.
	oldUser, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		return err
	}
	// If both emails are different, checks if the new email is already in use.
	if oldUser.Email != newUser.Email {
		_, err = s.userRepository.GetTinyUserByEmail(newUser.Email)
		if err == nil {
			return errors.New("email already in use")
		}
	}
	// If both registration numbers are different, checks if the new registration number is already in use.
	if oldUser.RegistrationNumber != newUser.RegistrationNumber {
		_, err = s.userRepository.GetTinyUserByRegistrationNumber(newUser.RegistrationNumber)
		if err == nil {
			return errors.New("registration number already in use")
		}
	}
	// Creates a new user payload and then replaces the old one.
	updatedUser := entity.UserPayload{
		ID:                 oldUser.ID,
		Name:               newUser.Name,
		BirthDate:          newUser.BirthDate,
		Email:              newUser.Email,
		RegistrationNumber: newUser.RegistrationNumber,
		Password:           oldUser.Password,
		UniversityID:       oldUser.University.ID,
		Status:             oldUser.Status,
		CreationDate:       oldUser.CreationDate,
	}
	err = s.userRepository.UpdateUser(updatedUser)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) AuthenticateUser(email, password string) (*entity.User, error) {

	user, err := s.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	// Compares passwords.
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
