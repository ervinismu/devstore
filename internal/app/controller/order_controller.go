package controller

import (
	"net/http"

	"github.com/ervinismu/devstore/internal/pkg/handler"
	"github.com/gin-gonic/gin"
)

type OrderService interface{
	Checkout() error
}

type OrderController struct {
	orderService OrderService
}

func NewOrderController(orderService OrderService) *OrderController {
	return &OrderController{orderService: orderService}
}

func (ctrl *OrderController) Checkout(ctx *gin.Context) {
	err := ctrl.orderService.Checkout()
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success checkout", nil)
}
