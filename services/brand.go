package services

import (
	"go-auth/domains"
	"go-auth/types"
	"go-auth/utils"
)

type BrandService struct {
	repo domains.BrandRepoUseCase
}

func NewBrandService(repo domains.BrandRepoUseCase) *BrandService {
	return &BrandService{repo: repo}
}

func (c BrandService) Create(req *types.BrandReq) (*types.BrandResp, error) {
	var brand domains.Brand
	utils.MapStruct(&req, &brand)
	BrandModel, err := c.repo.CreateBrand(&brand)
	if err != nil {
		return nil, err
	}
	return c.makeBrandResp(BrandModel), nil
}

func (c BrandService) FindBrandsByCompanyId(companyId uint) (*types.BrandResp, error) {
	//data, err := c.repo.FindBrandsByCompanyId(companyId)
	//if err != nil {
	//	return nil, err
	//}
	//return c.makeBrandResp(data), nil
	return nil, nil
}

func (c BrandService) makeBrandResp(data *domains.Brand) *types.BrandResp {
	var Brand = &types.BrandResp{}
	Brand.ID = data.ID
	Brand.Title = data.Title
	Brand.Phone = data.Phone
	Brand.Address = data.Address
	return Brand
}
