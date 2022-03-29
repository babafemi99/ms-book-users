package users

import (
	"bookApi/datasources/postgres/user_db"
	"bookApi/utils/msErrors"
	"context"
	"time"
)

var pql = user_db.Init()

func (u *User) GetUser() *msErrors.RestErrors {
	getErr := pql.QueryRow(context.Background(), "SELECT first_name, last_name, email FROM user_db WHERE id =$1", u.Id).Scan(&u.FirstName, &u.LastName, &u.Email)
	if getErr != nil {
		return msErrors.NewInternalServerError("Unable to fetch user", getErr)
	}
	return nil
}

func (u *User) Save() *msErrors.RestErrors {
	u.DateCreated = time.Now().UTC()
	u.DateUpdated = time.Now().UTC()
	_, insertErr := pql.Exec(context.Background(), "INSERT INTO user_db (id, first_name, last_name, email, data_created, date_updated) VALUES($1, $2, $3, $4, $5, $6)", u.Id, u.FirstName, u.LastName, u.Email, u.DateCreated, u.DateUpdated)
	if insertErr != nil {
		return msErrors.NewInternalServerError("Error inserting", insertErr)
	}
	return nil
}
