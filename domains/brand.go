package domains

import (
	"go-auth/types"
	"time"
)

type BrandRepoUseCase interface {
	CreateBrand(brand *Brand) (*Brand, error)
	FindBrandsByCompanyId(companyId uint) ([]*Brand, error)
	FindBrand(key string, value interface{}) (*Brand, error)
}

type BrandSvcUseCase interface {
	Create(req *types.BrandReq) (*types.BrandResp, error)
	FindBrandsByCompanyId(companyId uint) (*types.BrandResp, error)
}

type Brand struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	CompanyId uint   `json:"company_id"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	CreatedAt time.Time
	Company   Company
}
