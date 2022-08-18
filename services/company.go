package services

import (
	"go-auth/models"
	"go-auth/repositories"
	"go-auth/types"
)

type CompanyService struct {
	repo *repositories.CompanyRepository
}

func NewCompanyService(repo *repositories.CompanyRepository) *CompanyService {
	return &CompanyService{repo: repo}
}

func (c CompanyService) Create(req *types.CompanyReq) (*types.CompanyResp, error) {
	data := c.makeCompanyModel(req)
	companyModel, err := c.repo.Create(data)
	if err != nil {
		return nil, err
	}
	return c.makeCompanyResp(companyModel), nil
}

func (c CompanyService) Update(id int, req *types.CompanyReq) bool {
	company, err := c.repo.FindBy("id", id)
	if err != nil {
		return false
	}
	company.Title = req.Title
	company.Phone = req.Phone
	company.Address = req.Address
	result := c.repo.Update(company)
	if !result {
		return false
	}
	return true
}

func (c CompanyService) FindCompanyByUserId(userId int) (*types.CompanyResp, error) {
	data, err := c.repo.FindBy("ownerUserId", userId)
	if err != nil {
		return nil, err
	}
	return c.makeCompanyResp(data), nil
}

func (c CompanyService) makeCompanyModel(req *types.CompanyReq) *models.Company {
	var company = &models.Company{}
	company.Title = req.Title
	company.OwnerUserId = req.OwnerUserId
	company.Phone = req.Phone
	company.Address = req.Address
	return company
}

func (c CompanyService) makeCompanyResp(data *models.Company) *types.CompanyResp {
	var company = &types.CompanyResp{}
	company.ID = data.ID
	company.Title = data.Title
	company.OwnerUserId = data.OwnerUserId
	company.Phone = data.Phone
	company.Address = data.Address
	return company
}
