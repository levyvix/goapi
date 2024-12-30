package router

import (
	"github.com/gin-gonic/gin"
	"github.com/levyvix/goapi/handler"
)

func InitializeRoutes(router *gin.Engine) {

	//initialize handler

	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.GetOpening)
		v1.POST("/opening", handler.PostOpening)
		v1.PUT("/opening", handler.PutOpening)
		v1.DELETE("/opening", handler.DeleteOpening)
		v1.GET("/openings", handler.ListOpenings)
	}
}
