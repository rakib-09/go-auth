package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primarykey"`
	Title       string `json:"title"`
	OwnerUserId uint   `json:"ownerUserId"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	CreatedAt   string `json:"created_At"`
	Owner       User   `gorm:"references:OwnerUserId"`
}
