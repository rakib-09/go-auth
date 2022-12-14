package services

import (
	"github.com/rakib-09/go-auth/domains"
	"github.com/rakib-09/go-auth/types"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo domains.UserRepoUseCase
}

func NewUserService(ur domains.UserRepoUseCase) *UserService {
	return &UserService{userRepo: ur}
}

func (u *UserService) CreateUser(req *types.UserReq) error {
	data, _ := u.makeUserData(req)
	err := u.userRepo.CreateUser(data)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) GetUserById(userId uint, showPassword bool) (*types.UserResp, error) {
	user, err := u.userRepo.GetUserBy("id", userId)
	if err != nil {
		return nil, err
	}
	return u.makeUserResp(user, showPassword), nil
}

func (u *UserService) GetUserByEmail(userEmail string, password bool) (*types.UserResp, error) {
	user, err := u.userRepo.GetUserBy("email", userEmail)
	if err != nil {
		return nil, err
	}
	return u.makeUserResp(user, password), nil
}

func (u *UserService) hashPassword(plainPass string) string {
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(plainPass), 10)
	return string(hashedPass)
}

func (u *UserService) makeUserData(req *types.UserReq) (*domains.User, error) {
	var user = &domains.User{}
	user.Name = req.Name
	user.Password = u.hashPassword(req.Password)
	user.Email = req.Email

	return user, nil
}

func (u *UserService) makeUserResp(user *domains.User, showPassword bool) *types.UserResp {
	var userDetails = &types.UserResp{}
	if showPassword {
		userDetails.Password = user.Password
	}
	userDetails.ID = user.ID
	userDetails.Email = user.Email
	userDetails.Name = user.Name
	userDetails.CreatedAt = user.CreatedAt
	return userDetails
}
