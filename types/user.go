package types

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"time"
)

type UserResp struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
	CreatedAt time.Time
}

type UserReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *UserReq) Validate() *ValidationError {
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
