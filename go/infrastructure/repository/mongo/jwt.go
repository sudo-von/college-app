package mongo

import "gopkg.in/mgo.v2"

type JWTRepository struct {
	Session      *mgo.Session
	DatabaseName string
}

func NewJWTRepository(repository *Repository) *JWTRepository {
	return &JWTRepository{
		Session:      repository.Session,
		DatabaseName: repository.DatabaseName,
	}
}
