package repositories

import (
	"go-auth/db"
	"go-auth/domains"
)

type UserRepository struct {
	ds *db.AuthDatabase
}

func NewUserRepository(ds *db.AuthDatabase) *UserRepository {
	return &UserRepository{ds: ds}
}

func (repo *UserRepository) CreateUser(user *domains.User) error {
	return repo.ds.CreateUser(user)
}

func (repo *UserRepository) GetUserBy(key string, value interface{}) (*domains.User, error) {
	return repo.ds.GetUserBy(key, value)
}
