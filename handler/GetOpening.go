package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/levyvix/goapi/schemas"
)

func GetOpening(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendErr(ctx, http.StatusBadRequest, "id (query param) is required")
		return
	}

	opening := schemas.Opening{}

	//find opening
	err := db.First(&opening, id).Error
	if err != nil {
		sendErr(ctx, http.StatusNotFound, fmt.Sprintf("opening %s not found", id))
		return
	}

	sendSuccess(ctx, "get", opening)

}
