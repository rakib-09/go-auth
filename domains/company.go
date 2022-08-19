package domains

import (
	"go-auth/types"
	"time"
)

type CompanyRepoUseCase interface {
	Create(company *Company) error
	Update(data *Company) error
	FindBy(key string, value interface{}) (*Company, error)
}

type CompanySvcUseCase interface {
	Create(req *types.CompanyReq) error
	Update(id int, req *types.CompanyReq) error
	FindCompanyByUserId(userId uint) (*types.CompanyResp, error)
}

type Company struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title"`
	OwnerUserId int    `json:"owner_user_id"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	CreatedAt   time.Time
	User        User `json:"user,omitempty"`
	Brand       []Brand
}
