package services

import (
	"go-auth/domains"
	"go-auth/types"
)

type UserService struct {
	userRepo domains.UserRepoUseCase
}

func NewUserService(ur domains.UserRepoUseCase) *UserService {
	return &UserService{userRepo: ur}
}

func (u *UserService) createUser() {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) GetUserById(userId int) (*types.UserResp, error) {
	user, err := u.userRepo.GetUser("id", userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) GetUserByEmail(userEmail string) (*types.UserResp, error) {
	user, err := u.userRepo.GetUser("email", userEmail)
	if err != nil {
		return nil, err
	}
	return user, nil
}
