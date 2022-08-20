package db

import (
	"encoding/json"
	"fmt"
	"go-auth/db/entities"
	"go-auth/domains"
	"go-auth/utils"
)

func (db AuthDatabase) CreateBrand(brand *domains.Brand) (*domains.Brand, error) {
	var entity entities.Brand
	utils.MapStruct(&brand, &entity)
	result := db.DB.Model(&entities.Brand{}).Create(&entity)
	if result.Error != nil {
		return nil, result.Error
	}
	return brand, nil
}

func (db AuthDatabase) FindBrandsByCompanyId(companyId uint) ([]*domains.Brand, error) {
	var brands []entities.Brand
	var brandResp []domains.Brand
	if err := db.DB.Model(&entities.Brand{}).Joins("Company").Where("company_id", companyId).Take(&brands).Error; err != nil {
		return nil, err
	}
	stringify, _ := json.Marshal(brands)
	json.Unmarshal(stringify, &brandResp)
	return nil, nil
}

func (db AuthDatabase) FindBrand(key string, value interface{}) (*domains.Brand, error) {
	var brand entities.Brand
	var brandResp domains.Brand
	if err := db.DB.Model(&entities.Brand{}).Joins("Company").Where(fmt.Sprintf("%s = ?", key), value).
		First(&brand).Error; err != nil {
		return nil, err
	}
	stringify, _ := json.Marshal(brand)
	json.Unmarshal(stringify, &brandResp)
	return &brandResp, nil
}
