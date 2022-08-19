package services

import (
	"go-auth/domains"
	"go-auth/types"
)

type CompanyService struct {
	repo domains.CompanyRepoUseCase
}

func NewCompanyService(repo domains.CompanyRepoUseCase) *CompanyService {
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

func (c CompanyService) FindCompanyByUserId(userId uint) (*types.CompanyResp, error) {
	data, err := c.repo.FindBy("owner_user_id", userId)
	if err != nil {
		return nil, err
	}
	return c.makeCompanyResp(data), nil
}

func (c CompanyService) makeCompanyModel(req *types.CompanyReq) *domains.Company {
	var company = &domains.Company{}
	company.Title = req.Title
	company.OwnerUserId = req.OwnerUserId
	company.Phone = req.Phone
	company.Address = req.Address
	return company
}

func (c CompanyService) makeCompanyResp(data *domains.Company) *types.CompanyResp {
	var company = &types.CompanyResp{}
	company.ID = data.ID
	company.Title = data.Title
	company.OwnerUserId = data.OwnerUserId
	company.Phone = data.Phone
	company.Address = data.Address
	company.User.ID = data.User.ID
	company.User.Name = data.User.Name
	company.User.Email = data.User.Email
	return company
}
