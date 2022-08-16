package domains

import "go-auth/types"

type AuthUseCase interface {
	Login(email string, password string) (*types.LoginResp, error)
}

type JwtUseCase interface {
	CreateToken(userId uint) (*types.JwtToken, error)
}
