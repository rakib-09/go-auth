package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/rakib-09/go-auth/domains"
	"github.com/rakib-09/go-auth/types"
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
		return c.JSON(http.StatusBadRequest, types.CommonResponse(err.Error()))
	}

	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	if token, err = auth.AuthUseCase.Login(req.Email, req.Password); err != nil {
		return c.JSON(http.StatusBadRequest, types.CommonResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, token)

}
