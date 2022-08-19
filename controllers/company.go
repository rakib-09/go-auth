package controllers

import (
	"github.com/labstack/echo/v4"
	"go-auth/domains"
	"go-auth/types"
	"go-auth/utils"
	"net/http"
	"strconv"
)

type CompanyController struct {
	svc domains.CompanySvcUseCase
}

func NewCompanyController(svc domains.CompanySvcUseCase) *CompanyController {
	return &CompanyController{svc: svc}
}

func (cc CompanyController) Create(c echo.Context) error {
	var req *types.CompanyReq
	var err error
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, types.CommonResponse(err.Error()))
	}

	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	req.OwnerUserId = utils.GetUserIdFromJwt(c)
	res, err := cc.svc.Create(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.CommonResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, res)
}

func (cc CompanyController) Show(c echo.Context) error {
	userId := utils.GetUserIdFromJwt(c)
	res, err := cc.svc.FindCompanyByUserId(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.CommonResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, res)
}

func (cc CompanyController) Update(c echo.Context) error {
	var req *types.CompanyReq

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, types.CommonResponse(err.Error()))
	}

	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	req.OwnerUserId = utils.GetUserIdFromJwt(c)
	id, _ := strconv.Atoi(c.Param("id"))
	res := cc.svc.Update(id, req)
	if !res {
		return c.JSON(http.StatusBadRequest, types.CommonResponse("Update Failed"))
	}
	return c.JSON(http.StatusOK, types.CommonResponse("Successfully Updated!"))
}
