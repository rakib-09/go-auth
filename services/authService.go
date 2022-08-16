package services

import (
	"go-auth/domains"
	"go-auth/errors"
	"go-auth/types"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	jwtSvc  domains.JwtUseCase
	userSvc domains.UserSvcUseCase
}

func NewAuthService(
	userSvc domains.UserSvcUseCase,
	jwtSvc domains.JwtUseCase,
) *AuthService {
	return &AuthService{
		userSvc: userSvc,
		jwtSvc:  jwtSvc,
	}
}

func (svc *AuthService) Login(email string, password string) (*types.LoginResp, error) {
	token := &types.JwtToken{}
	ac := &types.LoginResp{}
	var err error
	identity, err := svc.authenticate(email, password)
	if err != nil {
		return nil, err
	}
	token, err = svc.jwtSvc.CreateToken(identity.ID)
	if err != nil {
		return nil, err
	}
	ac.AccessToken = token.AccessToken
	return ac, nil
}

func (svc *AuthService) authenticate(email string, password string) (*types.UserResp, error) {
	user, err := svc.userSvc.GetUserByEmail(email)
	if err != nil {
		return nil, errors.InvalidCreds
	}
	loginPass := []byte(password)
	hashedPass := []byte(user.Password)

	if err = bcrypt.CompareHashAndPassword(hashedPass, loginPass); err != nil {
		return nil, errors.InvalidCreds
	}

	return user, nil
}