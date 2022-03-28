package user_controller

import (
	"bookApi/domain/users"
	"bookApi/services"
	"bookApi/utils/msErrors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user users.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		restErr := msErrors.NewBadRequest("invalid json body", err)
		//restErr := msErrors.RestErrors{
		//	Message: "invalid json body - " + err.Error(),
		//	Status:  http.StatusBadRequest,
		//	Error:   "Unable to marshall JSON",
		//}
		c.JSON(restErr.Status, restErr)
		return
	}
	result, createErr := services.CreateUser(&user)
	if createErr != nil {
		c.JSON(createErr.Status, createErr)
	}
	c.JSON(http.StatusOK, result)
}
func GetUser(c *gin.Context) {
	id, idErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if idErr != nil {
		err := msErrors.NewBadRequest("invalid user id", idErr)
		c.JSON(err.Status, err)
	}
	user, errors := services.GetUser(id)
	if errors != nil {
		c.JSON(errors.Status, errors)
	}
	c.JSON(http.StatusOK, user)
}
func FindUser(c *gin.Context) {

}
func SearchUser(c *gin.Context) {

}
