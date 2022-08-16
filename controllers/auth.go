package controllers

import (
	"github.com/labstack/echo/v4"
	"go-auth/domains"
	"go-auth/errors"
	"go-auth/types"
	"net/http"
)

type AuthController struct {
	AuthUseCase domains.AuthUseCase
}

func NewAuthController(
	AuthSvc domains.AuthUseCase,
) *AuthController {
	return &AuthController{
		AuthUseCase: AuthSvc,
	}
}

func (auth AuthController) Login(c echo.Context) error {
	var req *types.LoginRequest
	var err error
	var token = &types.LoginResp{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, errors.SomethingWentWrong)
	}

	if token, err = auth.AuthUseCase.Login(req.Email, req.Password); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, token)

}
