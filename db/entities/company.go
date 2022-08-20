package entities

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	ID          uint    `gorm:"primarykey"`
	Title       string  `json:"title"`
	OwnerUserId int     `json:"owner_user_id"`
	Phone       string  `json:"phone"`
	Address     string  `json:"address"`
	User        User    `json:"user,omitempty" gorm:"foreignKey:OwnerUserId;references:ID"`
	Brand       []Brand `json:"brand,omitempty"`
}
