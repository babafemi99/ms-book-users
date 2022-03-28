package users

import (
	"bookApi/utils/msErrors"
	"errors"
	"fmt"
)

var userDB = make(map[int64]*User)

func (u *User) GetUser() *msErrors.RestErrors {
	result := userDB[u.Id]
	if result == nil {
		return msErrors.NewNotFoundRequestError(fmt.Sprintf("user %d not found", u.Id), errors.New("not found"))

	}
	u.Email = result.Email
	u.Id = result.Id
	u.FirstName = result.FirstName
	u.LastName = result.LastName
	u.DateCreated = result.DateCreated
	u.DateUpdated = result.DateUpdated
	return nil
}

func (u *User) Save() *msErrors.RestErrors {
	current := userDB[u.Id]
	if current != nil {
		if current.Email == u.Email {
			return msErrors.NewBadRequest(fmt.Sprint("email already exists"), errors.New("bad request"))
		}
		return msErrors.NewBadRequest(fmt.Sprintf("user %d already exists", u.Id), errors.New("bad request"))
	}
	userDB[u.Id] = u
	return nil
}
