package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/ervinismu/devstore/internal/app/schema"
	"github.com/ervinismu/devstore/internal/pkg/reason"
	log "github.com/sirupsen/logrus"
)

type PaymentCharger interface {
	Charge(chargePayload schema.OrderChargePayload) error
}

type OrderService struct {
	payment PaymentCharger
}

func NewOrderService(payment PaymentCharger) *OrderService {
	return &OrderService{payment: payment}
}

// only acccept bca va
func (svc *OrderService) Checkout() error {
	paymentPayload := schema.OrderChargePayload{
		Amount:        13000,
		OrderID:       svc.orderID(),
		CustomerName:  "Andi",
		CustomerEmail: "andi@studidevsecops.com",
		VANumber:      "1234567",
	}

	err := svc.payment.Charge(paymentPayload)
	if err != nil {
		log.Error(err)
		return errors.New(reason.FailedCheckout)
	}

	return nil
}

// ORDER-YYYYMMDD-HHMMSS
func (svc *OrderService) orderID() string {
	t := time.Now()
	strOrderID := fmt.Sprintf("ORDER-%d%d%d-%d%d%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	return strOrderID
}
