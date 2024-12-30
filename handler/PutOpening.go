package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PutOpening(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "put openings",
	})

}
