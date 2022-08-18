package repositories

import (
	"fmt"
	"go-auth/models"
	"gorm.io/gorm"
)

type CompanyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{db: db}
}

func (repo *CompanyRepository) Create(data *models.Company) (*models.Company, error) {
	result := repo.db.Model(&models.Company{}).Create(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

func (repo *CompanyRepository) Update(data *models.Company) (*models.Company, error) {
	return nil, nil
}

func (repo *CompanyRepository) findBy(key string, value interface{}) (*models.Company, error) {
	var company models.Company
	if err := repo.db.Model(&models.User{}).Where(fmt.Sprintf("%s = ?", key), value).
		First(&company).Error; err != nil {
		return nil, err
	}
	return &company, nil
}
