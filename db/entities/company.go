package entities

import (
	"gorm.io/gorm"
	"time"
)

type Company struct {
	gorm.Model
	Title       string `json:"title"`
	OwnerUserId int    `json:"owner_user_id"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	CreatedAt   time.Time
	User        User `gorm:"foreignKey:OwnerUserId"`
	Brand       []Brand
}
