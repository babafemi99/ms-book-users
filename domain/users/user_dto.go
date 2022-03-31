package users

import (
	"bookApi/utils/msErrors"
	"github.com/go-playground/validator/v10"
	"strings"
	"time"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64     `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email" validate:"required,email"`
	Password    string    `json:"-"`
	Status      string    `json:"status"`
	DateCreated time.Time `json:"-"`
	DateUpdated time.Time `json:"-"`
}

type UserList []*User

func (u *User) Validate() *msErrors.RestErrors {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		return msErrors.NewBadRequest("Error validating user", err)
	}
	return nil
}
