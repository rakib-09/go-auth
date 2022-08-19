package repositories

import (
	"go-auth/db"
	"go-auth/domains"
)

type CompanyRepository struct {
	ds *db.AuthDatabase
}

func NewCompanyRepository(ds *db.AuthDatabase) *CompanyRepository {
	return &CompanyRepository{ds: ds}
}

func (repo *CompanyRepository) Create(company *domains.Company) error {
	return repo.ds.Create(company)
}

func (repo *CompanyRepository) Update(data *domains.Company) error {
	return repo.ds.Update(data)
}

func (repo *CompanyRepository) FindBy(key string, value interface{}) (*domains.Company, error) {
	return repo.ds.FindBy(key, value)
}
