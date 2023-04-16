package controller

import (
	"net/http"

	"github.com/ervinismu/devstore/internal/app/schema"
	"github.com/ervinismu/devstore/internal/app/service"
	"github.com/ervinismu/devstore/internal/pkg/handler"
	"github.com/gin-gonic/gin"
)

type SessionController struct {
	service service.ISessionService
}

func NewSessionController(service service.ISessionService) *SessionController {
	return &SessionController{service: service}
}

func (ctrl *SessionController) SignIn(ctx *gin.Context) {
	req := &schema.SignInReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	resp,err := ctrl.service.SignIn(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success login", resp)
}
