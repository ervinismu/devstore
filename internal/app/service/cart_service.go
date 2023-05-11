package service

import (
	"errors"
	"strconv"

	"github.com/ervinismu/devstore/internal/app/model"
	"github.com/ervinismu/devstore/internal/app/schema"
	"github.com/ervinismu/devstore/internal/pkg/reason"
)

type CartService struct {
	productRepo  ProductRepository
	cartRepo     CartRepository
	cartItemRepo CartItemRepository
}

func NewCartService(productRepo ProductRepository, cartRepo CartRepository, cartItemRepo CartItemRepository) *CartService {
	return &CartService{
		productRepo:  productRepo,
		cartRepo:     cartRepo,
		cartItemRepo: cartItemRepo,
	}
}

func (svc *CartService) AddToCart(req *schema.AddToCartReq) error {

	// find product
	existingProduct, err := svc.productRepo.GetByID(strconv.Itoa(req.ProductID))
	if err != nil {
		return errors.New(reason.ProductCannotGetDetail)
	}

	// check stock
	if existingProduct.TotalStock < req.Quantity {
		return errors.New(reason.InsufficientStock)
	}

	// find or init cart
	existingCart, err := svc.cartRepo.GetByUserID(req.UserID)
	if existingCart.ID <= 0  && err != nil {
		initCartData := model.Cart{UserID: req.UserID, TotalPrice: 0}
		existingCart, _ = svc.cartRepo.Create(initCartData)
	}

	// get cart item by product id and cart id
	cartItem, err := svc.cartItemRepo.GetByCartIDAndProductID(existingCart.ID, existingCart.ID)
	if cartItem.ID <= 0 && err != nil {
		// create cart item
		newCartItem := model.CartItem{}
		newCartItem.BasePrice = existingProduct.Price
		newCartItem.TotalPrice = existingProduct.Price * req.Quantity
		newCartItem.Quantity = req.Quantity
		newCartItem.CartID = existingCart.ID
		newCartItem.ProductID = existingProduct.ID

		strconv.Itoa(existingProduct.Price)

		err := svc.cartItemRepo.Create(newCartItem)
		if err != nil {
			return errors.New(reason.FailedAddCart)
		}

	}

	return nil
}
