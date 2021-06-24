package server

import (
	"github.com/gin-gonic/gin"

	"chat/controllers"
)

func CreateRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("templates/*")

	router.GET("/health", controllers.Health)
	router.GET("/", controllers.Index)

	return router
}
