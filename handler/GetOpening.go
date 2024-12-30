package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOpening(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get openings",
	})

}
