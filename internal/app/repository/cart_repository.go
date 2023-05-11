package repository

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/ervinismu/devstore/internal/app/model"
	"github.com/jmoiron/sqlx"
)

type CartRepository struct {
	DB *sqlx.DB
}

func NewCartRepository(db *sqlx.DB) *CartRepository {
	return &CartRepository{DB: db}
}

// get detail cart
func (cr *CartRepository) GetByUserID(userID int) (model.Cart, error) {
	var (
		sqlStatement = `
			SELECT id, total_price, user_id
			FROM carts
			WHERE user_id = $1
		`
		cart model.Cart
	)
	err := cr.DB.QueryRowx(sqlStatement, userID).StructScan(&cart)
	if err != nil {
		log.Error(fmt.Errorf("error CartRepository - GetByID : %w", err))
		return cart, err
	}

	return cart, nil
}

// get detail cart
func (cr *CartRepository) Create(cart model.Cart) (model.Cart, error) {
	var (
		lastInsertID = 0
		sqlStatement = `
			INSERT INTO carts (user_id, total_price)
			VALUES ($1, $2) RETURNING id
		`
	)
	err := cr.DB.QueryRowx(sqlStatement, cart.UserID, cart.TotalPrice).Scan(&lastInsertID)
	if err != nil {
		log.Error(fmt.Errorf("error CartRepository - Create : %w", err))
		return cart, err
	}

	cart.ID = lastInsertID

	return cart, nil
}
