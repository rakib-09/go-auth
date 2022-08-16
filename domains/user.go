package domains

import (
	"go-auth/models"
	"go-auth/types"
)

type UserRepoUseCase interface {
	UpsertUser(user *models.User) (*types.UserResp, error)
	GetUser(key string, value interface{}) (*types.UserResp, error)
}

type UserSvcUseCase interface {
	createUser()
	GetUserById(userId int) (*types.UserResp, error)
	GetUserByEmail(userEmail string) (*types.UserResp, error)
}
