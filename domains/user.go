package domains

import (
	"go-auth/models"
	"go-auth/types"
)

type UserRepoUseCase interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUser(key string, value interface{}) (*models.User, error)
}

type UserSvcUseCase interface {
	CreateUser(req *types.UserReq) (*types.UserResp, error)
	GetUserById(userId int) (*models.User, error)
	GetUserByEmail(userEmail string) (*models.User, error)
}
