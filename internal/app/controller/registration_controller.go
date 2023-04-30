package controller

import (
	"net/http"

	"github.com/ervinismu/devstore/internal/app/schema"
	"github.com/ervinismu/devstore/internal/app/service"
	"github.com/ervinismu/devstore/internal/pkg/handler"
	"github.com/gin-gonic/gin"
)

type RegistrationController struct {
	service service.IRegistrationService
}

func NewRegistrationController(service service.IRegistrationService) *RegistrationController {
	return &RegistrationController{service: service}
}

func (ctrl *RegistrationController) Register(ctx *gin.Context) {
	req := &schema.RegisterReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	err := ctrl.service.Register(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success register", nil)
}
