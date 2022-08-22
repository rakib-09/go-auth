package db

import (
	"fmt"
	"github.com/rakib-09/go-auth/db/entities"
	"github.com/rakib-09/go-auth/domains"
	"github.com/rakib-09/go-auth/utils"
)

func (db AuthDatabase) CreateUser(user *domains.User) error {
	var entity entities.User
	err := utils.MapStruct(&user, &entity)
	if err != nil {
		return err
	}
	result := db.DB.Model(&entities.User{}).Create(&entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db AuthDatabase) GetUserBy(key string, value interface{}) (*domains.User, error) {
	var user domains.User
	if err := db.DB.Model(&entities.User{}).Where(fmt.Sprintf("%s = ?", key), value).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
