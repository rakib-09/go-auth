package types

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
)

type CompanyReq struct {
	Title       string `json:"title,omitempty"`
	OwnerUserId int    `json:"ownerUserId,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Address     string `json:"address,omitempty"`
}

type CompanyResp struct {
	ID          uint     `json:"id"`
	Title       string   `json:"title"`
	OwnerUserId int      `json:"ownerUserId"`
	Phone       string   `json:"phone"`
	Address     string   `json:"address"`
	User        UserResp `json:"user,omitempty"`
}

func (r *CompanyReq) Validate() *ValidationError {
	err := v.ValidateStruct(r,
		v.Field(&r.Title, v.Required),
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
