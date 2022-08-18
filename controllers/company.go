package controllers

import "go-auth/services"

type CompanyController struct {
	svc services.CompanyService
}

func NewCompanyController(svc services.CompanyService) *CompanyController {
	return &CompanyController{svc: svc}
}

func (c CompanyController) Create() {

}
