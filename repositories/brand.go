package repositories

import (
	"go-auth/db"
	"go-auth/domains"
)

type BrandRepository struct {
	ds *db.AuthDatabase
}

func NewBrandRepository(ds *db.AuthDatabase) *BrandRepository {
	return &BrandRepository{ds: ds}
}

func (repo *BrandRepository) CreateBrand(brand *domains.Brand) (*domains.Brand, error) {
	return repo.ds.CreateBrand(brand)
}
func (repo *BrandRepository) FindBrandsByCompanyId(companyId uint) ([]*domains.Brand, error) {
	return repo.ds.FindBrandsByCompanyId(companyId)
}
func (repo *BrandRepository) FindBrand(key string, value interface{}) (*domains.Brand, error) {
	return repo.ds.FindBrand(key, value)
}
