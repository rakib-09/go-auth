package services

import "go-auth/repositories"

type CompanyService struct {
	repo repositories.CompanyRepository
}

func NewCompanyService(repo repositories.CompanyRepository) *CompanyService {
	return &CompanyService{repo: repo}
}
