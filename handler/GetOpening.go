package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/levyvix/goapi/schemas"
)

//	@BasePath		/api/v1
//	@Summary		Get Opening
//	@Description	Get a single Opening by ID
//	@Tags			Openings
//	@Accept			json
//	@Produce		json
//	@Param			id	query		string	true	"Opening ID"
//	@Success		200	{object}	handler.GetOpeningResponse
//	@Failure		400	{object}	handler.ErrorResponse
//	@Failure		404	{object}	handler.ErrorResponse
//	@Router			/opening [get]
func GetOpening(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendErr(ctx, http.StatusBadRequest, "id (query param) is required")
		return
	}

	opening := schemas.Opening{}

	// Find opening
	err := db.First(&opening, id).Error
	if err != nil {
		sendErr(ctx, http.StatusNotFound, fmt.Sprintf("opening %s not found", id))
		return
	}

	sendSuccess(ctx, "get", opening)
}
