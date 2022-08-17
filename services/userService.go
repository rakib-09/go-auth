package services

import (
	"go-auth/domains"
	"go-auth/models"
	"go-auth/types"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo domains.UserRepoUseCase
}

func NewUserService(ur domains.UserRepoUseCase) *UserService {
	return &UserService{userRepo: ur}
}

func (u *UserService) CreateUser(req *types.UserReq) (*types.UserResp, error) {
	var newUser = &types.UserResp{}
	data, _ := u.makeUserData(req)
	user, err := u.userRepo.CreateUser(data)
	if err != nil {
		return nil, err
	}
	newUser.ID = user.ID
	newUser.Email = user.Email
	newUser.Name = user.Name
	newUser.CreatedAt = user.CreatedAt
	return newUser, err
}

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

func (u *UserService) makeUserData(req *types.UserReq) (*models.User, error) {
	var user = &models.User{}
	user.Name = req.Name
	user.Password = u.hashPassword(req.Password)
	user.Email = req.Email

	return user, nil
}

func (u *UserService) hashPassword(plainPass string) string {
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(plainPass), 10)
	return string(hashedPass)
}
