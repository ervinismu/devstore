package controller

import (
	"net/http"

	"github.com/ervinismu/devstore/internal/app/schema"
	"github.com/ervinismu/devstore/internal/pkg/handler"
	"github.com/gin-gonic/gin"
)

type SessionService interface {
	Login(req *schema.LoginReq) (schema.LoginResp, error)
}

type SessionController struct {
	service SessionService
}

func NewSessionController(service SessionService) *SessionController {
	return &SessionController{
		service: service,
	}
}

func (ctrl *SessionController) Login(ctx *gin.Context) {
	req := &schema.LoginReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	resp, err := ctrl.service.Login(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success login", resp)
}

// refresh
func (ctrl *SessionController) Refresh(ctx *gin.Context) {

}

// logout
