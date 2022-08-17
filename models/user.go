package models

import (
	"time"
)

type User struct {
	ID        uint   `json:"id" gorm:"primarykey"`
	Name      string `json:"name"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
	CreatedAt time.Time
}
