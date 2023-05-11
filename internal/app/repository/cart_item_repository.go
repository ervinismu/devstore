package repository

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/ervinismu/devstore/internal/app/model"
	"github.com/jmoiron/sqlx"
)

type CartItemRepository struct {
	DB *sqlx.DB
}

func NewCartItemRepository(db *sqlx.DB) *CartItemRepository {
	return &CartItemRepository{DB: db}
}

// get detail cartItem
func (cr *CartItemRepository) GetByCartIDAndProductID(cartID int, productID int) (model.CartItem, error) {
	var (
		sqlStatement = `
			SELECT id, quantity, base_price, total_price, cart_id, product_id
			FROM cart_items
			WHERE cart_id = $1 AND product_id = $2 LIMIT 1
		`
		cartItem model.CartItem
	)
	err := cr.DB.QueryRowx(sqlStatement, cartID, productID).StructScan(&cartItem)
	if err != nil {
		log.Error(fmt.Errorf("error CartItemRepository - GetByID : %w", err))
		return cartItem, err
	}

	return cartItem, nil
}

// get detail cartItem
func (cr *CartItemRepository) Create(cartItem model.CartItem) error {
	var (
		sqlStatement = `
			INSERT INTO cart_items (quantity, base_price, total_price, cart_id, product_id)
			VALUES ($1, $2, $3, $4, $5)
		`
	)

	_, err := cr.DB.Exec(
		sqlStatement,
		cartItem.Quantity,
		cartItem.BasePrice,
		cartItem.TotalPrice,
		cartItem.CartID,
		cartItem.ProductID)
	if err != nil {
		log.Error(fmt.Errorf("error CartItemRepository - Create : %w", err))
		return err
	}

	return nil
}
