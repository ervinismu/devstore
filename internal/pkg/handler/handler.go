package handler

import (
	"net/http"

	"github.com/ervinismu/devstore/internal/pkg/reason"
	"github.com/ervinismu/devstore/internal/pkg/validator"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func ResponseError(ctx *gin.Context, statusCode int, message string) {
	resp := ResponseBody{
		Status:  "error",
		Message: message,
	}
	ctx.JSON(statusCode, resp)
}

func ResponseSuccess(ctx *gin.Context, statusCode int, message string, data interface{}) {
	resp := ResponseBody{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	ctx.JSON(statusCode, resp)
}

// BindAndCheck bind request and check
func BindAndCheck(ctx *gin.Context, data interface{}) bool {

	if err := ctx.ShouldBind(data); err != nil {
		log.Errorf("http_handle BindAndCheck fail, %s", err.Error())

		ResponseError(ctx, http.StatusUnprocessableEntity, reason.RequestFormatError)
		return true
	}

	isError := validator.Check(data)
	if isError {
		ResponseError(ctx, http.StatusUnprocessableEntity, reason.RequestFormatError)
		return true
	}

	return false
}
