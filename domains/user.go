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
	GetUserById(userId int, password bool) (*types.UserResp, error)
	GetUserByEmail(userEmail string, password bool) (*types.UserResp, error)
}
