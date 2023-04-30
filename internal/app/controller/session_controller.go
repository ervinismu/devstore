package controller

import (
	"net/http"
	"strconv"

	"github.com/ervinismu/devstore/internal/app/schema"
	"github.com/ervinismu/devstore/internal/pkg/handler"
	"github.com/ervinismu/devstore/internal/pkg/reason"
	"github.com/gin-gonic/gin"
)

type SessionService interface {
	Login(req *schema.LoginReq) (schema.LoginResp, error)
	Logout(UserID int) error
	Refresh(req *schema.RefreshTokenReq) (schema.RefreshTokenResp, error)
}

type RefreshTokenVerifier interface {
	VerifyRefreshToken(tokenString string) (string, error)
}

type SessionController struct {
	service SessionService
	tokenMaker RefreshTokenVerifier
}

func NewSessionController(service SessionService, tokenMaker RefreshTokenVerifier) *SessionController {
	return &SessionController{
		service: service,
		tokenMaker: tokenMaker,
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
	refreshToken := ctx.GetHeader("refresh_token")
	if refreshToken == "" {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "cannot refresh token")
	}

	sub, err := ctrl.tokenMaker.VerifyRefreshToken(refreshToken)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnauthorized, reason.FailedRefreshToken)
		return
	}

	intSub, _ := strconv.Atoi(sub)
	req := &schema.RefreshTokenReq{}
	req.RefreshToken = refreshToken
	req.UserID = intSub

	resp, err := ctrl.service.Refresh(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.FailedRefreshToken)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success refresh", resp)
}

// logout
func (ctrl *SessionController) Logout(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.GetString("user_id"))
	err := ctrl.service.Logout(userID)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success logout", nil)
}
