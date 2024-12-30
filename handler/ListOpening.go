package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/levyvix/goapi/schemas"
)

func ListOpenings(ctx *gin.Context) {
	// return all openings

	opening := []schemas.Opening{}

	err := db.Find(&opening).Error
	if err != nil {
		sendErr(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list", opening)
}
