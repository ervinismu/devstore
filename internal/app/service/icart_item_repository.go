package service

import "github.com/ervinismu/devstore/internal/app/model"

type CartItemRepository interface {
	Create(cartItem model.CartItem) error
	GetByCartIDAndProductID(cartID int, productID int) (model.CartItem, error)
}
