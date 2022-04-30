package users

import (
	"bookApi/utils/msErrors"
	"github.com/go-playground/validator/v10"
	"strings"
)

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (u *UserLogin) Validate() *msErrors.RestErrors {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		return msErrors.NewBadRequest("Error validating user", err)
	}
	return nil
}
