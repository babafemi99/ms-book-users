package app

import (
	"bookApi/controllers/ping_controller"
	"bookApi/controllers/user_controller"
)

func mapUrl() {
	router.GET("ping", ping_controller.Ping)
	router.GET("users/:user_id", user_controller.GetUser)
	//router.GET("users/search", controllers.SearchUser)
	router.POST("users", user_controller.CreateUser)
}
