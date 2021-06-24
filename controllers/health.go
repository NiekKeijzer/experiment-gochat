package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func Health(c *gin.Context) {
	c.String(http.StatusOK, "Alive!")
}
