package db

import (
	"fmt"
	"go-auth/db/entities"
	"go-auth/domains"
)

func (db AuthDatabase) CreateUser(user *domains.User) (*domains.User, error) {
	result := db.DB.Model(&entities.User{}).Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (db AuthDatabase) GetUserBy(key string, value interface{}) (*domains.User, error) {
	var user domains.User
	if err := db.DB.Model(&entities.User{}).Where(fmt.Sprintf("%s = ?", key), value).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
