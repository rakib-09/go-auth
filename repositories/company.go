package repositories

import (
	"fmt"
	"go-auth/models"
	"go-auth/types"
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

func (repo *CompanyRepository) Update(id int, data *types.CompanyReq) (bool, error) {
	err := repo.db.Model(&models.Company{}).Where("id", id).Updates(data).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *CompanyRepository) FindBy(key string, value interface{}) (*models.Company, error) {
	var company models.Company
	if err := repo.db.Model(&models.Company{}).Where(fmt.Sprintf("%s = ?", key), value).
		First(&company).Error; err != nil {
		return nil, err
	}
	return &company, nil
}
