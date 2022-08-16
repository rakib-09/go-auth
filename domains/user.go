package domains

import (
	"go-auth/models"
)

type UserRepoUseCase interface {
	UpsertUser(user *models.User) (*models.User, error)
	GetUser(key string, value interface{}) (*models.User, error)
}

type UserSvcUseCase interface {
	//createUser()
	GetUserById(userId int) (*models.User, error)
	GetUserByEmail(userEmail string) (*models.User, error)
}
