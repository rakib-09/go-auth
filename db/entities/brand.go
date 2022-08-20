package entities

import (
	"gorm.io/gorm"
	"time"
)

type Brand struct {
	gorm.Model
	ID        uint   `json:"id" gorm:"primarykey"`
	Title     string `json:"title"`
	CompanyId uint   `json:"company_id"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	CreatedAt time.Time
	Company   Company `json:"company,omitempty"gorm:"foreignKey:CompanyId"`
}
