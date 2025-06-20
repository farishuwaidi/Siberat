package router

import (
	"Siberat/middleware"

	"github.com/gin-gonic/gin"
)

func TestRoute(api *gin.RouterGroup) {

	api.Use(middleware.JWTMiddleWare())
	api.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
		"message": "Middleware pass",
		})
	})
}