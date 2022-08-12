package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	//authService domains.AuthUseCase
}

func NewUserController() *UserController {
	return &UserController{
		//authService: authSvc,
	}
}

func (uc *UserController) GetUser(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (uc *UserController) CreateUser(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
