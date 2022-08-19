package services

import (
	"go-auth/domains"
	"go-auth/repositories"
	"go-auth/types"
)

type BrandService struct {
	repo *repositories.BrandRepository
}

func NewBrandService(repo *repositories.BrandRepository) *BrandService {
	return &BrandService{repo: repo}
}

func (c BrandService) Create(req *types.BrandReq) (*types.BrandResp, error) {
	data := c.makeBrandModel(req)
	BrandModel, err := c.repo.CreateBrand(data)
	if err != nil {
		return nil, err
	}
	return c.makeBrandResp(BrandModel), nil
}

func (c BrandService) FindBrandsByCompanyId(companyId uint) (*types.BrandResp, error) {
	data, err := c.repo.FindBrandsByCompanyId(companyId)
	if err != nil {
		return nil, err
	}
	return c.makeBrandResp(data), nil
}

func (c BrandService) makeBrandModel(req *types.BrandReq) *domains.Brand {
	var Brand = &domains.Brand{}
	Brand.Title = req.Title
	Brand.Phone = req.Phone
	Brand.Address = req.Address
	return Brand
}

func (c BrandService) makeBrandResp(data *domains.Brand) *types.BrandResp {
	var Brand = &types.BrandResp{}
	Brand.ID = data.ID
	Brand.Title = data.Title
	Brand.Phone = data.Phone
	Brand.Address = data.Address
	Brand.Company.ID = data.Company.ID
	Brand.Company.Title = data.Company.Title
	Brand.Company.Phone = data.Company.Phone
	Brand.Company.Address = data.Company.Address
	return Brand
}
