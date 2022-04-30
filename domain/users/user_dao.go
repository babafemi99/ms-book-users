package users

import (
	"bookApi/datasources/postgres/user_db"
	"bookApi/logger"
	"bookApi/utils/msErrors"
	"context"
	"errors"
	"fmt"
)

var pql = user_db.Init()

const (
	selectStatementById     = "SELECT first_name, last_name, email, status FROM user_db WHERE id =$1"
	selectStatementByEmail  = "SELECT id, first_name, last_name, email, status, password FROM user_db WHERE email =$1;"
	selectStatementByStatus = "SELECT id, first_name, last_name, email, status FROM user_db WHERE status =$1"
	insertStatement         = "INSERT INTO user_db (id, first_name, last_name, email, password, data_created, date_updated, status) VALUES($1, $2, $3, $4, $5, $6, $7, $8)"
	updateStatement         = "UPDATE user_db SET first_name = $1, last_name = $2, email =$3, date_updated = $4 WHERE id = $5;"
	deleteStatement         = "DELETE FROM user_db WHERE id =$1"
)

func (u *User) GetUser() *msErrors.RestErrors {
	getErr := pql.QueryRow(context.Background(), selectStatementById, u.Id).Scan(&u.FirstName, &u.LastName, &u.Email, &u.Status)
	if getErr != nil {
		logger.Error("Error Getting user", getErr)
		return msErrors.NewInternalServerError("Unable to fetch user", getErr)
	}
	return nil
}

func (u *User) GetUserByEmail() *msErrors.RestErrors {
	getErr := pql.QueryRow(context.Background(), selectStatementByEmail, u.Email).Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.Status, &u.Password)
	if getErr != nil {
		logger.Error("Error Getting user", getErr)
		return msErrors.NewInternalServerError("Unable to fetch user", getErr)
	}
	return nil
}

func (u *User) Save() *msErrors.RestErrors {
	_, insertErr := pql.Exec(context.Background(), insertStatement, u.Id, u.FirstName, u.LastName, u.Email, u.Password, u.DateCreated, u.DateUpdated, u.Status)
	if insertErr != nil {
		logger.Error("Error saving user", insertErr)
		return msErrors.NewInternalServerError("Error inserting", insertErr)
	}
	return nil
}
func (u *User) UpdateUser() *msErrors.RestErrors {
	_, UpdateErr := pql.Exec(context.Background(), updateStatement, u.FirstName, u.LastName, u.Email, u.DateUpdated, u.Id)
	if UpdateErr != nil {
		logger.Error("Error updating user", UpdateErr)
		return msErrors.NewInternalServerError("Error inserting", UpdateErr)
	}
	return nil
}
func (u *User) DeleteUser() *msErrors.RestErrors {
	_, getErr := pql.Exec(context.Background(), deleteStatement, u.Id)
	if getErr != nil {
		logger.Error("Error deleting user", getErr)
		return msErrors.NewInternalServerError("Unable to delete user", getErr)
	}
	return nil
}

func (u *User) FindByStatus(status string) (UserList, *msErrors.RestErrors) {
	query, getErr := pql.Query(context.Background(), selectStatementByStatus, status)
	if getErr != nil {
		return nil, msErrors.NewNotFoundRequestError("Unable to fetch users by status", getErr)
	}
	defer query.Close()
	results := make([]*User, 0)
	for query.Next() {
		var user User
		queryErr := query.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Status)
		if queryErr != nil {
			logger.Error("Error finding by status", queryErr)
			return nil, msErrors.NewInternalServerError("unable to scan", queryErr)
		}
		results = append(results, &user)
	}
	if len(results) == 0 {
		return nil, msErrors.NewNotFoundRequestError(fmt.Sprintf("no user matching %s found", status), errors.New(""))
	}
	return results, nil
}
