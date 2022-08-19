package domains

import (
	"go-auth/types"
	"time"
)

type CompanyRepoUseCase interface {
	Create(company *Company) (*Company, error)
	Update(data *Company) bool
	FindBy(key string, value interface{}) (*Company, error)
}

type CompanySvcUseCase interface {
	Create(req *types.CompanyReq) (*types.CompanyResp, error)
	Update(id int, req *types.CompanyReq) bool
	FindCompanyByUserId(userId uint) (*types.CompanyResp, error)
}

type Company struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	OwnerUserId uint   `json:"OwnerUserId"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	CreatedAt   time.Time
	User        User
}
