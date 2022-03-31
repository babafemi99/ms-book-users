package app

import (
	"bookApi/datasources/postgres/user_db"
	logger "bookApi/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	user_db.Init()
	logs := logger.GetLogger()
	logs.Info("About to start application")
	mapUrl()
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
