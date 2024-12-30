package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/levyvix/goapi/docs"
	"github.com/levyvix/goapi/handler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(router *gin.Engine) {

	basePath := "/api/v1"

	//init docs
	docs.SwaggerInfo.BasePath = basePath

	//initialize handler
	handler.InitializeHandler()

	v1 := router.Group(basePath)
	{
		v1.GET("/opening", handler.GetOpening)
		v1.POST("/opening", handler.PostOpening)
		v1.PUT("/opening", handler.PutOpening)
		v1.DELETE("/opening", handler.DeleteOpening)
		v1.GET("/openings", handler.ListOpenings)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
