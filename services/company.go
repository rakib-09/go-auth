package services

import (
	"go-auth/domains"
	"go-auth/types"
	"go-auth/utils"
)

type CompanyService struct {
	repo domains.CompanyRepoUseCase
}

func NewCompanyService(repo domains.CompanyRepoUseCase) *CompanyService {
	return &CompanyService{repo: repo}
}

func (c CompanyService) Create(req *types.CompanyReq) error {
	data := c.makeCompanyModel(req)
	err := c.repo.Create(data)
	if err != nil {
		return err
	}
	return nil
}

func (c CompanyService) Update(id int, req *types.CompanyReq) error {
	var company domains.Company
	utils.MapStruct(&req, &company)
	company.ID = uint(id)
	result := c.repo.Update(&company)
	if result != nil {
		return result
	}
	return nil
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
