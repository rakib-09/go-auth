package entities

import (
	"gorm.io/gorm"
	"time"
)

type Company struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primarykey"`
	Title       string `json:"title"`
	OwnerUserId uint   `json:"ownerUserId"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	CreatedAt   time.Time
	User        User `gorm:"foreignKey:OwnerUserId"`
}
