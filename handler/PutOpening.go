package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/levyvix/goapi/schemas"
)

func PutOpening(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)
	err := request.Validate()
	if err != nil {
		logger.Errorf("error validating request: %v", err.Error())
		sendErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendErr(ctx, http.StatusBadRequest, "id (query param) is required")
		return
	}

	opening := schemas.Opening{}

	//find opening
	err = db.First(&opening, id).Error
	if err != nil {
		sendErr(ctx, http.StatusNotFound, fmt.Sprintf("opening %s not found", id))
		return
	}

	//update opening
	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Salary != 0 {
		opening.Salary = request.Salary
	}

	//update opening
	err = db.Save(&opening).Error
	if err != nil {
		sendErr(ctx, http.StatusInternalServerError, fmt.Sprintf("error updating opening %s", id))
		return
	}

	sendSuccess(ctx, "update", request)
}
