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

func PostOpening(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "posts openings",
	})

}

func DeleteOpening(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete openings",
	})

}

func PutOpening(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "put openings",
	})

}

func ListOpenings(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "list openings",
	})
}
