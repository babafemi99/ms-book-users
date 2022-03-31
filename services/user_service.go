package services

import (
	"bookApi/domain/users"
	"bookApi/utils/cyrpto_utils"
	"bookApi/utils/msErrors"
	"fmt"
	"time"
)

type userService struct {
}

type userServiceInterface interface {
	CreateUser(user *users.User) (*users.User, *msErrors.RestErrors)
	GetUser(id int64) (*users.User, *msErrors.RestErrors)
	UpdateUser(u *users.User) *msErrors.RestErrors
	PatchUser(u *users.User) *msErrors.RestErrors
	Delete(id int64) *msErrors.RestErrors
	Search(status string) (users.UserList, *msErrors.RestErrors)
}

func NewUserService() userServiceInterface {
	return &userService{}
}

func (usrv *userService) CreateUser(user *users.User) (*users.User, *msErrors.RestErrors) {
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

func (usrv *userService) GetUser(id int64) (*users.User, *msErrors.RestErrors) {
	res := &users.User{Id: id}
	err := res.GetUser()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (usrv *userService) UpdateUser(u *users.User) *msErrors.RestErrors {
	res := &users.User{Id: u.Id}
	_, getErr := usrv.GetUser(u.Id)
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

func (usrv *userService) PatchUser(u *users.User) *msErrors.RestErrors {
	fmt.Println("first u:", u)
	mainUser, getErr := usrv.GetUser(u.Id)
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

func (usrv *userService) Delete(id int64) *msErrors.RestErrors {
	res := &users.User{
		Id: id,
	}
	err := res.DeleteUser()
	if err != nil {
		return err
	}
	return nil
}

func (usrv *userService) Search(status string) (users.UserList, *msErrors.RestErrors) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
