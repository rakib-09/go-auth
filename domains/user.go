package domains

import (
	"github.com/rakib-09/go-auth/types"
	"time"
)

type UserRepoUseCase interface {
	CreateUser(user *User) error
	GetUserBy(key string, value interface{}) (*User, error)
}

type UserSvcUseCase interface {
	CreateUser(req *types.UserReq) error
	GetUserById(userId uint, showPassword bool) (*types.UserResp, error)
	GetUserByEmail(userEmail string, password bool) (*types.UserResp, error)
}

type User struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
	CreatedAt time.Time
}
