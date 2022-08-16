package services

import (
	"go-auth/domains"
	"go-auth/models"
)

type UserService struct {
	userRepo domains.UserRepoUseCase
}

func NewUserService(ur domains.UserRepoUseCase) *UserService {
	return &UserService{userRepo: ur}
}

//func (u *UserService) createUser() {
//	//TODO implement me
//	panic("implement me")
//}

func (u *UserService) GetUserById(userId int) (*models.User, error) {
	user, err := u.userRepo.GetUser("id", userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) GetUserByEmail(userEmail string) (*models.User, error) {
	user, err := u.userRepo.GetUser("email", userEmail)
	if err != nil {
		return nil, err
	}
	return user, nil
}
