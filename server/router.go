package server

import (
	"github.com/gin-gonic/gin"

	"chat/controllers"
	"chat/service/chat"
)

func CreateRouter() *gin.Engine {
	hub := chat.NewHub()
	go hub.Run()

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Dynamic routes
	router.GET("/health", controllers.Health)
	router.GET("/ws", func(c *gin.Context) {
		controllers.ServeWS(hub, c)
	})

	// Resources
	router.Static("/assets", "./assets")
	router.StaticFile("/", "./assets/index.html")

	return router
}
