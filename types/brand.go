package types

import v "github.com/go-ozzo/ozzo-validation/v4"

type BrandReq struct {
	Title     string `json:"title,omitempty"`
	CompanyId uint   `json:"company_id,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Address   string `json:"address,omitempty"`
}

type BrandResp struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	CompanyId uint   `json:"company_id"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	//Company   CompanyResp `json:"company,omitempty"`
}

func (r *BrandReq) Validate() *ValidationError {
	err := v.ValidateStruct(r,
		v.Field(&r.Title, v.Required),
		v.Field(&r.CompanyId, v.Required),
		v.Field(&r.Phone, v.Required, v.Length(10, 20)),
		v.Field(&r.Address, v.Required),
	)
	if err != nil {
		return &ValidationError{
			Error: err,
		}
	}
	return nil
}
