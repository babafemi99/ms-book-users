package user_controller

import (
	"bookApi/domain/users"
	"bookApi/services"
	"bookApi/utils/msErrors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

var (
	userService = services.NewUserService()
)

func CreateUser(c *gin.Context) {
	var user users.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		restErr := msErrors.NewBadRequest("invalid json body", err)
		c.JSON(restErr.Status, restErr)
		return
	}
	fmt.Println(user)
	result, createErr := userService.CreateUser(&user)
	if createErr != nil {
		c.JSON(createErr.Status, createErr)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}
func GetUser(c *gin.Context) {
	id, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	user, errors := userService.GetUser(id)
	if errors != nil {
		c.JSON(errors.Status, errors)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}
func FindUser(c *gin.Context) {

}
func Search(c *gin.Context) {
	status := c.Query("status")
	fmt.Println(status)
	StatusUser, err := userService.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, StatusUser.Marshall(c.GetHeader("X-Public") == "true"))
}
func EditUser(c *gin.Context) {
	var user users.User
	id, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		restErr := msErrors.NewBadRequest("invalid json body", err)
		c.JSON(restErr.Status, restErr)
		return
	}
	user.Id = id
	fmt.Println(&user)
	updateErr := userService.UpdateUser(&user)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"result": "Successful"})
}

func PatchUser(c *gin.Context) {
	log.Printf("Inside %d", 1)
	var user users.User
	id, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		restErr := msErrors.NewBadRequest("invalid json body", err)
		c.JSON(restErr.Status, restErr)
		return
	}
	user.Id = id
	patchErr := userService.PatchUser(&user)
	if patchErr != nil {
		c.JSON(patchErr.Status, patchErr)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"result": "Successful"})
}
func DeleteUser(c *gin.Context) {
	id, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	errors := userService.Delete(id)
	if errors != nil {
		c.JSON(errors.Status, errors)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"result": "Successful"})
}

func getUserId(idParam string) (int64, *msErrors.RestErrors) {
	id, idErr := strconv.ParseInt(idParam, 10, 64)
	if idErr != nil {
		return 0, msErrors.NewBadRequest("invalid user id", idErr)
	}
	return id, nil
}
