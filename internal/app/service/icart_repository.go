package service

import "github.com/ervinismu/devstore/internal/app/model"

type CartRepository interface {
	GetByUserID(userID int) (model.Cart, error)
	Create(cart model.Cart) (model.Cart, error)
}
