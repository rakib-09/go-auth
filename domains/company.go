package domains

import (
	"go-auth/models"
	"go-auth/types"
)

type CompanyRepoUseCase interface {
	Create(company *models.Company) (*models.Company, error)
	Update(company *models.Company) (*models.Company, error)
	FindBy(key string, value interface{}) (*models.Company, error)
}

type CompanySvcUseCase interface {
	Create(req *types.CompanyReq) (*types.CompanyResp, error)
	Update(id int, req *types.CompanyReq) bool
	FindCompanyByUserId(userId int) (*types.CompanyResp, error)
}
