package controllers

import (
	"github.com/labstack/echo/v4"
	"go-auth/services"
)

type CompanyController struct {
	svc services.CompanyService
}

func NewCompanyController(svc services.CompanyService) *CompanyController {
	return &CompanyController{svc: svc}
}

func (cc CompanyController) Create(c echo.Context) {

}
