package jwt

import (
	"errors"
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/usecase/user"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	jwtRepository  JWTRepository
	userRepository user.UserRepository
}

func NewService(jwtRepository JWTRepository, userRepository user.UserRepository) *Service {
	return &Service{
		jwtRepository,
		userRepository,
	}
}

// AuthenticateUser returns a user in case of successful match.
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
