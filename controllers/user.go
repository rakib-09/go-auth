package controllers

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go-auth/config"
	"go-auth/domains"
	"go-auth/types"
	"net/http"
)

type UserController struct {
	UserSvc domains.UserSvcUseCase
}

func NewUserController(UserSvc domains.UserSvcUseCase) *UserController {
	return &UserController{
		UserSvc: UserSvc,
	}
}

func (uc *UserController) GetUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.JwtCustomClaims)
	userId := claims.UID
	userInfo, err := uc.UserSvc.GetUserById(int(userId), false)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, userInfo)
}

func (uc *UserController) CreateUser(c echo.Context) error {
	var req *types.UserReq
	var err error
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorMessage(err.Error()))
	}

	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	res, err := uc.UserSvc.CreateUser(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorMessage(err.Error()))
	}
	return c.JSON(http.StatusOK, res)
}
