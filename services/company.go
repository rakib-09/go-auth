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
	var company domains.Company
	utils.MapStruct(&req, &company)
	err := c.repo.Create(&company)
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
	var companyDetails types.CompanyResp
	data, err := c.repo.FindBy("owner_user_id", userId)
	if err != nil {
		return nil, err
	}
	utils.MapStruct(&data, &companyDetails)
	return &companyDetails, nil
}
