package repositories

import (
	"fmt"
	"go-auth/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	result := repo.db.Model(&models.User{}).Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) GetUser(key string, value interface{}) (*models.User, error) {
	var user models.User
	if err := repo.db.Model(&models.User{}).Where(fmt.Sprintf("%s = ?", key), value).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
