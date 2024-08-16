package handlers

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", serveHomePage)
	router.POST("/upload", handleFileUpload)
}
