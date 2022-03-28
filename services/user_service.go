package services

import (
	"bookApi/domain/users"
	"bookApi/utils/msErrors"
)

func CreateUser(user *users.User) (*users.User, *msErrors.RestErrors) {
	validateError := user.Validate()
	if validateError != nil {
		return nil, validateError
	}
	err := user.Save()
	if err != nil {
		return nil, err
	}
	return user, nil
}
func GetUser(id int64) (*users.User, *msErrors.RestErrors) {
	res := &users.User{Id: id}
	err := res.GetUser()
	if err != nil {
		return nil, err
	}
	return res, nil
}
