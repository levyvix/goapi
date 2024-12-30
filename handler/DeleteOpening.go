package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/levyvix/goapi/schemas"
)

func DeleteOpening(ctx *gin.Context) {

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

	//delete opening
	err = db.Delete(&opening).Error
	if err != nil {
		sendErr(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening %s", id))
		return
	}

	sendSuccess(ctx, "delete", opening)
}
