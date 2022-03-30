package services

import (
	"bookApi/domain/users"
	"bookApi/utils/cyrpto_utils"
	"bookApi/utils/msErrors"
	"fmt"
	"time"
)

func CreateUser(user *users.User) (*users.User, *msErrors.RestErrors) {
	validateError := user.Validate()
	if validateError != nil {
		return nil, validateError
	}
	user.DateCreated = time.Now().UTC()
	user.DateUpdated = time.Now().UTC()
	user.Password = cyrpto_utils.GetMd5(user.Password)
	user.Status = users.StatusActive
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

func UpdateUser(u *users.User) *msErrors.RestErrors {
	res := &users.User{Id: u.Id}
	_, getErr := GetUser(u.Id)
	if getErr != nil {
		return getErr
	}
	fmt.Println(u)
	res.FirstName = u.FirstName
	res.LastName = u.LastName
	res.Email = u.Email
	res.DateUpdated = time.Now()
	fmt.Println(res)
	updateErr := res.UpdateUser()
	if updateErr != nil {
		return updateErr
	}
	return nil
}

func PatchUser(u *users.User) *msErrors.RestErrors {
	fmt.Println("first u:", u)
	mainUser, getErr := GetUser(u.Id)
	if getErr != nil {
		return getErr
	}
	if u.FirstName != "" {
		mainUser.FirstName = u.FirstName
	}
	if u.LastName != "" {
		mainUser.LastName = u.LastName
	}
	if u.Email != "" {
		mainUser.Email = u.Email
	}
	fmt.Println("second u:", u)
	mainUser.DateUpdated = time.Now()
	fmt.Println("third u:", mainUser)
	updateErr := mainUser.UpdateUser()
	if updateErr != nil {
		return updateErr
	}

	return nil
}

func Delete(id int64) *msErrors.RestErrors {
	res := &users.User{
		Id: id,
	}
	err := res.DeleteUser()
	if err != nil {
		return err
	}
	return nil
}
func Search(status string) ([]*users.User, *msErrors.RestErrors) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
