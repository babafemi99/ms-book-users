package users

import (
	"bookApi/utils/msErrors"
	"github.com/go-playground/validator/v10"
	"strings"
	"time"
)

type User struct {
	Id          int64     `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email" validate:"required"`
	DateCreated time.Time `json:"dateCreated"`
	DateUpdated time.Time `json:"dateUpdated"`
}

func (u *User) Validate() *msErrors.RestErrors {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	//if u.Email == "" {
	//	return msErrors.NewBadRequest("Email cannot be empty", errors.New("provide a valid email address"))
	//}
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		return msErrors.NewBadRequest("Error validating user", err)
	}
	return nil
}
