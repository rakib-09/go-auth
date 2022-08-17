package controllers

import (
	"github.com/labstack/echo/v4"
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
	return c.NoContent(http.StatusOK)
}

func (uc *UserController) CreateUser(c echo.Context) error {
	var req *types.UserReq
	var err error
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	res, err := uc.UserSvc.CreateUser(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, res)
}
