package controller

import (
	"net/http"
	"strconv"

	"github.com/ervinismu/devstore/internal/app/schema"
	"github.com/ervinismu/devstore/internal/pkg/handler"
	"github.com/gin-gonic/gin"
)

type CartService interface {
	AddToCart(req *schema.AddToCartReq) error
}

type CartController struct {
	service CartService
}

func NewCartController(service CartService) *CartController {
	return &CartController{service: service}
}

func (ctrl *CartController) AddToCart(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.GetString("user_id"))
	req := &schema.AddToCartReq{}
	req.UserID = userID
	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := ctrl.service.AddToCart(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success add to cart", nil)
}
