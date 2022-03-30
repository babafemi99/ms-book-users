package app

import (
	"bookApi/controllers/ping_controller"
	"bookApi/controllers/user_controller"
)

func mapUrl() {
	router.GET("ping", ping_controller.Ping)
	router.GET("users/:user_id", user_controller.GetUser)
	router.GET("internal/users/search", user_controller.Search)
	router.POST("users", user_controller.CreateUser)
	router.PUT("users/:user_id", user_controller.EditUser)
	router.PATCH("users/:user_id", user_controller.PatchUser)
	router.DELETE("users/:user_id", user_controller.DeleteUser)
}
