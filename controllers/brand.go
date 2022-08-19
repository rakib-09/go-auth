package controllers

import (
	"github.com/labstack/echo/v4"
	_const "go-auth/const"
	"go-auth/domains"
	"go-auth/types"
	"net/http"
)

type BrandController struct {
	svc domains.BrandSvcUseCase
}

func NewBrandController(svc domains.BrandSvcUseCase) *BrandController {
	return &BrandController{
		svc: svc,
	}
}

func (bc BrandController) Create(c echo.Context) error {
	var req *types.BrandReq
	var err error
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, types.CommonResponse(err.Error()))
	}

	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	_, err = bc.svc.Create(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.CommonResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, types.CommonResponse(_const.SuccessfullyCreated))
}
