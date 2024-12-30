package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.GET("/openings", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "get openings",
			})
		})
		v1.POST("/openings", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "post openings",
			})
		})
		v1.PUT("/openings", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "put openings",
			})
		})
		v1.DELETE("/openings", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "delete openings",
			})
		})
		v1.PATCH("/openings", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "patch openings",
			})
		})
	}
}
