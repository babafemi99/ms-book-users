package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApp() {
	mapUrl()
	router.Run(":8080")
}
