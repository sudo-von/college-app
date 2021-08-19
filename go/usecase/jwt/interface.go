package jwt

import "freelancer/college-app/go/entity"

type JWTRepository interface {
}

type UseCase interface {
	AuthenticateUser(user, password string) (*entity.User, error)
}
