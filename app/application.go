package app

import (
	"bookApi/datasources/postgres/user_db"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router = gin.Default()
)

func StartApp() {
	user_db.Init()
	log.Println("run succefully")
	mapUrl()
	router.Run(":8080")
}
