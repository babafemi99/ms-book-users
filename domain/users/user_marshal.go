package users

import (
	"encoding/json"
	"time"
)

type publicUser struct {
	Id     int64  `json:"user_id"`
	Email  string `json:"email" validate:"required,email"`
	Status string `json:"status"`
}
type privateUser struct {
	Id          int64     `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email" validate:"required,email"`
	Status      string    `json:"status"`
	DateCreated time.Time `json:"-"`
	DateUpdated time.Time `json:"-"`
}

func (u *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return publicUser{
			Id:     u.Id,
			Email:  u.Email,
			Status: u.Status,
		}
	} else {
		user, _ := json.Marshal(u)
		var privateUser privateUser
		if err := json.Unmarshal(user, &privateUser); err != nil {
			return nil
		}
		return privateUser
	}
}
func (u UserList) Marshall(isPublic bool) interface{} {
	result := make([]interface{}, len(u))
	for index, user := range u {
		result[index] = user.Marshall(isPublic)
	}
	return result
}
