package services

import (
	"bookApi/domain/users"
	"bookApi/utils/cyrpto_utils"
	"bookApi/utils/msErrors"
	"errors"
	"time"
)

type userService struct {
}

type UserServiceInterface interface {
	CreateUser(user *users.User) (*users.User, *msErrors.RestErrors)
	GetUser(id int64) (*users.User, *msErrors.RestErrors)
	GetUserEmail(email string) (*users.User, *msErrors.RestErrors)
	UpdateUser(u *users.User) *msErrors.RestErrors
	PatchUser(u *users.User) *msErrors.RestErrors
	Delete(id int64) *msErrors.RestErrors
	Search(status string) (users.UserList, *msErrors.RestErrors)
	FindByCredentials(login *users.UserLogin) (*users.User, *msErrors.RestErrors)
}

func NewUserService() UserServiceInterface {
	return &userService{}
}

func (usrv *userService) FindByCredentials(login *users.UserLogin) (*users.User, *msErrors.RestErrors) {
	vErr := login.Validate()
	if vErr != nil {
		return nil, vErr
	}

	user, getErr := usrv.GetUserEmail(login.Email)
	if getErr != nil {
		return nil, msErrors.NewNotFoundRequestError("Invalid credentials", errors.New("try again"))
	}

	res := &users.User{
		Email:    login.Email,
		Password: login.Password,
	}
	hash := cyrpto_utils.CheckPasswordHash(res.Password, user.Password)
	if !hash {
		return nil, msErrors.NewNotFoundRequestError("Invalid credentials", errors.New("try again"))
	}

	return user, nil
}

func (usrv *userService) CreateUser(user *users.User) (*users.User, *msErrors.RestErrors) {
	validateError := user.Validate()
	if validateError != nil {
		return nil, validateError
	}
	user.DateCreated = time.Now().UTC()
	user.DateUpdated = time.Now().UTC()
	user.Status = users.StatusActive
	password, err2 := cyrpto_utils.HashPassword(user.Password)
	if err2 != nil {
		panic("can't save ")
	}
	user.Password = password
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

func (usrv *userService) GetUserEmail(email string) (*users.User, *msErrors.RestErrors) {
	res := &users.User{Email: email}
	err := res.GetUserByEmail()
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
	res.FirstName = u.FirstName
	res.LastName = u.LastName
	res.Email = u.Email
	res.DateUpdated = time.Now()
	updateErr := res.UpdateUser()
	if updateErr != nil {
		return updateErr
	}
	return nil
}

func (usrv *userService) PatchUser(u *users.User) *msErrors.RestErrors {
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
	mainUser.DateUpdated = time.Now()
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
