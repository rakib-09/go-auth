package domains

import (
	"github.com/rakib-09/go-auth/types"
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
	ID          uint    `json:"id,omitempty"`
	Title       string  `json:"title,omitempty"`
	OwnerUserId int     `json:"owner_user_id,omitempty"`
	Phone       string  `json:"phone,omitempty"`
	Address     string  `json:"address,omitempty"`
	User        User    `json:"user,omitempty"`
	Brand       []Brand `json:"Brand,omitempty"`
}
