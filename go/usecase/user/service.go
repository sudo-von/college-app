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

func (s *Service) CreateUser(newUser entity.UserPayload) error {

	// Check if the email is valid.
	err := checkmail.ValidateFormat(newUser.Email)
	if err != nil {
		return errors.New("invalid email")
	}
	// Check if email is already in use.
	_, err = s.userRepository.GetUserByEmail(newUser.Email)
	if err == nil {
		return errors.New("email already in use")
	}
	// Check if the registration number is valid.
	if len(strings.Replace(newUser.RegistrationNumber, " ", "", -1)) != 8 {
		return errors.New("invalid registration_number")
	}
	// Check if registration number is already in use.
	_, err = s.userRepository.GetUserByRegistrationNumber(newUser.RegistrationNumber)
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
