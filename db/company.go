package db

import (
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
	var company entities.Company
	err := utils.MapStruct(&data, &company)
	err = db.DB.Model(&entities.Company{}).Where("id", data.ID).Updates(&company).Error
	if err != nil {
		return _const.SomethingWentWrong
	}
	return nil
}

func (db AuthDatabase) FindBy(key string, value interface{}) (*domains.Company, error) {
	var company entities.Company
	var companyResp domains.Company
	if err := db.DB.Model(&entities.Company{}).Joins("User").Preload("Brand").Where(fmt.Sprintf("companies.%s = ?", key), value).
		First(&company).Error; err != nil {
		return nil, _const.SomethingWentWrong
	}
	utils.MapStruct(&company, &companyResp)
	return &companyResp, nil
}
