package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/levyvix/goapi/schemas"
)

//	@BasePath		/api/v1
//	@Summary		Create Opening
//	@Description	Create a single Opening by ID
//	@Tags			Openings
//	@Accept			json
//	@Produce		json
//	@Param			opening	body		handler.CreateOpeningRequest	true	"Opening"
//	@Success		200		{object}	handler.GetOpeningResponse
//	@Failure		400		{object}	handler.ErrorResponse
//	@Failure		500		{object}	handler.ErrorResponse
//	@Router			/opening [post]
func PostOpening(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)
	err := request.Validate()
	if err != nil {
		logger.Errorf("error validating request: %v", err.Error())
		sendErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	logger.Infof("received request: %+v", request)

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	err = db.Create(&opening).Error

	if err != nil {
		logger.Errorf("error creating row in database: %v", err.Error())
		sendErr(ctx, http.StatusInternalServerError, "Error creating opening on database")
		return
	}
	sendSuccess(ctx, "create", opening)

}
