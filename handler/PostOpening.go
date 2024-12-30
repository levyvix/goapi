package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostOpening(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "posts openings",
	})

}
