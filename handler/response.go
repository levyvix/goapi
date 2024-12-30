package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendErr(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, gin.H{"message": msg, "error_code": code})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation %s success", op),
		"data":    data,
	})
}
