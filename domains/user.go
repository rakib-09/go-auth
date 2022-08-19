package domains

import (
	"go-auth/types"
	"time"
)

type UserRepoUseCase interface {
	CreateUser(user *User) (*User, error)
	GetUserBy(key string, value interface{}) (*User, error)
}

type UserSvcUseCase interface {
	CreateUser(req *types.UserReq) (*types.UserResp, error)
	GetUserById(userId uint, password bool) (*types.UserResp, error)
	GetUserByEmail(userEmail string, password bool) (*types.UserResp, error)
}

type User struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
	CreatedAt time.Time
}
