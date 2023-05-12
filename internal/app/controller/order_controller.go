package controller

import (
	"net/http"
	"strconv"

	"github.com/ervinismu/devstore/internal/app/schema"
	"github.com/ervinismu/devstore/internal/pkg/handler"
	"github.com/gin-gonic/gin"
)

type OrderService interface {
	Checkout(req schema.OrderReq) error
}

type OrderController struct {
	orderService OrderService
}

func NewOrderController(orderService OrderService) *OrderController {
	return &OrderController{orderService: orderService}
}

func (ctrl *OrderController) Checkout(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.GetString("user_id"))
	req := schema.OrderReq{}
	req.UserID = userID

	err := ctrl.orderService.Checkout(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success checkout", nil)
}
