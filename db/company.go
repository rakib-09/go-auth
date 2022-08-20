package db

import (
	"encoding/json"
	"fmt"
	_const "go-auth/const"
	"go-auth/db/entities"
	"go-auth/domains"
	"go-auth/utils"
)

func (db AuthDatabase) Create(company *domains.Company) error {
	var entity entities.Company
	err := utils.MapStruct(&company, &entity)
	if err != nil {
		return nil
	}
	result := db.DB.Model(&entities.Company{}).Create(&entity)
	if result.Error != nil {
		return _const.SomethingWentWrong
	}
	return nil
}

func (db AuthDatabase) Update(data *domains.Company) error {
	err := db.DB.Save(&data).Error
	if err != nil {
		return _const.SomethingWentWrong
	}
	return nil
}

func (db AuthDatabase) FindBy(key string, value interface{}) (*domains.Company, error) {
	var company entities.Company
	var companyResp domains.Company
	if err := db.DB.Model(&entities.Company{}).Joins("User").Where(fmt.Sprintf("%s = ?", key), value).
		First(&company).Error; err != nil {
		return nil, _const.SomethingWentWrong
	}
	stringify, _ := json.Marshal(company)
	json.Unmarshal(stringify, &companyResp)
	return &companyResp, nil
}
