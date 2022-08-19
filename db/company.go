package db

import (
	"encoding/json"
	"fmt"
	"go-auth/db/entities"
	"go-auth/domains"
)

func (db AuthDatabase) Create(company *domains.Company) (*domains.Company, error) {
	result := db.DB.Model(&entities.Company{}).Create(&company)
	if result.Error != nil {
		return nil, result.Error
	}
	return company, nil
}

func (db AuthDatabase) Update(data *domains.Company) bool {
	err := db.DB.Save(&data).Error
	if err != nil {
		return false
	}
	return true
}

func (db AuthDatabase) FindBy(key string, value interface{}) (*domains.Company, error) {
	var company entities.Company
	var companyResp domains.Company
	if err := db.DB.Model(&entities.Company{}).Preload("User").Where(fmt.Sprintf("%s = ?", key), value).
		Take(&company).Error; err != nil {
		return nil, err
	}
	stringify, _ := json.Marshal(company)
	json.Unmarshal(stringify, &companyResp)
	return &companyResp, nil
}
