package domains

import "time"

type BrandRepoUseCase interface {
}

type BrandSvcUseCase interface {
}

type Brand struct {
	ID        uint   `json:"id" gorm:"primarykey"`
	Title     string `json:"title"`
	CompanyId uint   `json:"company_id"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	CreatedAt time.Time
	Company   Company
}
