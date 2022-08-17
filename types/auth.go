package types

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JwtToken struct {
	UserID        uint   `json:"uid"`
	AccessToken   string `json:"act"`
	RefreshToken  string `json:"rft"`
	AccessExpiry  int64  `json:"axp"`
	RefreshExpiry int64  `json:"rxp"`
}

type LoginResp struct {
	AccessToken string `json:"access_token"`
}

func (r *LoginRequest) Validate() *ValidationError {
	err := v.ValidateStruct(r,
		v.Field(&r.Email, v.Required, is.Email),
		v.Field(&r.Password, v.Required, v.Length(6, 50)),
	)
	if err != nil {
		return &ValidationError{
			Error: err,
		}
	}
	return nil
}
