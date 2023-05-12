package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/ervinismu/devstore/internal/app/schema"
	"github.com/ervinismu/devstore/internal/pkg/reason"
	log "github.com/sirupsen/logrus"
)

type OrderRepository interface{}

type PaymentCharger interface {
	Charge(paymentPayload schema.OrderChargePayload) error
}

type OrderService struct {
	orderRepo OrderRepository
	payment   PaymentCharger
}

func NewOrderService(orderRepo OrderRepository, payment PaymentCharger) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
		payment:   payment,
	}
}

func (svc *OrderService) Checkout(req schema.OrderReq) error {
	paymentPayload := schema.OrderChargePayload{
		Amount:        10000,
		OrderID:       svc.orderID(),
		CustomerName:  "ervin",
		CustomerEmail: "social.ervin@gmail.com",
		VANumber:      "12345678901",
	}

	err := svc.payment.Charge(paymentPayload)
	if err != nil {
		log.Error(err)
		return errors.New(reason.FailedCheckout)
	}

	return nil
}

func (svc *OrderService) orderID() string {
	t := time.Now()
	formatted := fmt.Sprintf("ORDER-%d%d%d-%d%d%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	return formatted
}
